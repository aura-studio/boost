# redowl

使用Redis List 和Redis PubSub模拟一个SQS功能

1. 支持消息确认
2. 支持死信队列
3. 支持触发器
4. 支持redis Cluster Client 和Redis Client，使用redis.cmdable 对象传入作为构造函数参数，使用redisv9库
5. 支持消息可见性超时

## Redis Key 结构

默认 `Prefix` 为 `redowl`，队列名为 `name` 时：

- Ready 队列（List）：`{Prefix}:{name}:ready`
	- 存储：消息 ID（string）
	- 方向：生产 `RPUSH`，消费 `LPOP/BLPOP`
- DLQ 队列（List）：`{Prefix}:{name}:dlq`
	- 存储：消息 ID（string）
	- 方向：进入 DLQ 时 `RPUSH`；DLQ 消费 `LPOP/BLPOP`；redrive 使用 `RPOPLPUSH` 移回 ready
- 消息内容（Hash）：`{Prefix}:{name}:msg:{id}`
	- 字段：
		- `body`：base64 编码后的消息体
		- `attrs`：JSON（map[string]string）
		- `rc`：接收次数（int）
		- `created_at_ms`：创建时间（Unix 毫秒）
- Receipt 映射（Hash）：`{Prefix}:{name}:receipt`
	- 映射：`receiptHandle -> messageID`
	- 用途：`Ack(receiptHandle)` 时定位消息 ID
- In-flight 可见性（ZSet）：`{Prefix}:{name}:inflight`
	- member：`receiptHandle`
	- score：`visibleAtUnixMs`（可见性超时到期时间的 Unix 毫秒）
	- 用途：`RequeueExpiredOnce` 扫描超时 receipt 并重新投递
- 事件 Channel（PubSub）：`{Prefix}:{name}:events`
	- payload：JSON `redowl.Event`
	- 事件类型：`sent / received / requeued / to_dlq / dlq_received / dlq_redriven / acked`
	- 说明：除了 per-queue channel 外，redowl 还会同时向 namespace channel 发布相同事件，便于消费者“先订阅再动态发现队列”
- 事件 Channel（PubSub, namespace）：`{Prefix}:events`
	- payload：JSON `redowl.Event`（包含 `Queue` 字段）
	- 用途：消费者无需预先知道队列名/数量，通过订阅该 channel 动态发现 `Queue` 并启动对应消费逻辑

## 最小用例（含事件订阅）

下面示例包含：创建队列、订阅事件、发送、接收、Ack。

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aura-studio/boost/redowl"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	q, _ := redowl.New(
		rdb,
		"orders",
		redowl.WithPrefix("redowl"),
		redowl.WithVisibilityTimeout(5*time.Second),
		redowl.WithMaxReceiveCount(3),
	)

	fmt.Println("ready:", "redowl:orders:ready")
	fmt.Println("dlq:", "redowl:orders:dlq")
	fmt.Println("events:", "redowl:orders:events")

	unsub, _ := q.Subscribe(ctx, func(e redowl.Event) {
		fmt.Println("event:", e.Type, e.MessageID)
	})
	defer func() { _ = unsub() }()

	id, _ := q.Send(ctx, []byte("hello"), map[string]string{"trace_id": "t-1"})
	_ = id

	msg, _ := q.Receive(ctx)
	if msg == nil {
		return
	}
	fmt.Println("recv:", msg.ID, string(msg.Body), msg.ReceiveCount)

	_ = q.Ack(ctx, msg.ReceiptHandle)
}
```

## 跨进程示例（生产者 / 消费者）

下面演示“生产者”和“消费者”分别在两个独立进程中运行的典型用法：

### producer.go

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aura-studio/boost/redowl"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	q, _ := redowl.New(rdb, "orders", redowl.WithVisibilityTimeout(30*time.Second))

	for i := 0; i < 10; i++ {
		_, _ = q.Send(ctx, []byte("hello"), map[string]string{"i": fmt.Sprint(i)})
	}
}
```

### consumer.go

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aura-studio/boost/redowl"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	q, _ := redowl.New(rdb, "orders", redowl.WithVisibilityTimeout(30*time.Second))

	for {
		msg, err := q.ReceiveWithWait(ctx, 5*time.Second)
		if err != nil {
			panic(err)
		}
		if msg == nil {
			continue
		}

		fmt.Println("got:", msg.ID, string(msg.Body), msg.ReceiveCount)

		// 业务处理成功后 Ack
		_ = q.Ack(ctx, msg.ReceiptHandle)
	}
}
```

### 多队列并发消费示例（Game*）：队列内顺序、队列间并发、每条消息独立 goroutine

下面演示：

- A 进程：分别向多个队列（`Game1..GameN`）按序生产
- B 进程：同时监听所有 `Game*` 队列
	- 每个队列内严格顺序消费（上一条处理完成后再取下一条）
	- 不同队列之间并发消费
	- 每条消息处理都在独立 goroutine 中执行

#### producer_games.go

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aura-studio/boost/redowl"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	const (
		gameCount   = 14
		perGameMsgs = 5
	)

	for i := 0; i < gameCount; i++ {
		qName := fmt.Sprintf("Game%d", i+1)
		q, _ := redowl.New(rdb, qName, redowl.WithVisibilityTimeout(30*time.Second))

		prefix := byte('A' + i)
		for j := 1; j <= perGameMsgs; j++ {
			payload := fmt.Sprintf("%c%d", prefix, j)
			_, _ = q.Send(ctx, []byte(payload), map[string]string{"game": fmt.Sprint(i + 1)})
		}
	}
}
```

#### consumer_games.go

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aura-studio/boost/redowl"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	const (
		gameCount   = 14
		perGameMsgs = 5
	)

	// 每个队列一个 worker：保证队列内顺序；不同 worker 之间天然并发。
	var wg sync.WaitGroup
	for i := 0; i < gameCount; i++ {
		qName := fmt.Sprintf("Game%d", i+1)
		q, _ := redowl.New(rdb, qName, redowl.WithVisibilityTimeout(30*time.Second))

		wg.Add(1)
		go func(q *redowl.Queue) {
			defer wg.Done()
			for k := 0; k < perGameMsgs; k++ {
				msg, err := q.ReceiveWithWait(ctx, 5*time.Second)
				if err != nil {
					panic(err)
				}
				if msg == nil {
					k--
					continue
				}

				done := make(chan struct{})
				go func(m *redowl.Message) {
					// 业务处理逻辑...
					time.Sleep(30 * time.Millisecond)
					_ = q.Ack(ctx, m.ReceiptHandle)
					close(done)
				}(msg)

				<-done // 等待本队列的本条消息处理完成，再取下一条
			}
		}(q)
	}

	wg.Wait()
}
```

### 关于“消费者崩溃”和可见性超时

- 如果消费者在处理过程中崩溃/未 Ack，消息会在 `VisibilityTimeout` 到期后变为可再次投递。
- 你可以：
  - 让消费者周期性调用 `RequeueExpiredOnce`（或开启 `WithReaperInterval`），以便及时回收超时 in-flight 消息；
  - 或者依赖下一次 `Receive/ReceiveWithWait` 的 best-effort 回收（内部会尝试回收一批超时消息）。

## 快速开始

### 基本用法：发送 / 接收 / Ack

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aura-studio/boost/redowl"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	q, _ := redowl.New(rdb, "orders", redowl.WithVisibilityTimeout(30*time.Second))

	_, _ = q.Send(ctx, []byte("hello"), map[string]string{"trace_id": "t-1"})

	msg, _ := q.Receive(ctx)
	if msg == nil {
		return
	}

	fmt.Println(string(msg.Body), msg.Attributes["trace_id"], msg.ReceiveCount)

	// 成功处理后确认删除
	_ = q.Ack(ctx, msg.ReceiptHandle)
}
```

### 可见性超时

- `Receive/ReceiveWithWait` 会把消息标记为 in-flight，并在 `VisibilityTimeout` 到期后允许重新投递。
- 你可以手动调用 `RequeueExpiredOnce` 回收超时消息；或配置 `WithReaperInterval` 开启后台回收器。

### 触发器（PubSub）

- 如果你传入的 `redis.Cmdable` 本身是 `*redis.Client/*redis.ClusterClient`，会自动启用同一个客户端作为触发器的 PubSub 客户端。
- 也可以显式传入 `WithTriggerClient(client)`。

```go
unsub, err := q.Subscribe(ctx, func(e redowl.Event) {
	// e.Type: sent/received/requeued/to_dlq/dlq_received/dlq_redriven/acked
})
defer func() { _ = unsub() }()
```

### Redis Cluster Client 用法

`redowl.New` 的第一个参数是 `redis.Cmdable`，因此可以直接传入 `*redis.ClusterClient`：

```go
cluster := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"}})
q, _ := redowl.New(cluster, "orders")
```

说明：本仓库的单元测试使用 `miniredis`，它不支持完整 Redis Cluster 协议；因此 Cluster 场景的测试以 build tag 的“外部集成测试”形式提供（需要真实 Redis Cluster）。

#### Cluster case（对应 multi-Game 并发/顺序用例）的集成测试

仓库中提供了一个 ClusterClient 版本的集成测试（需要真实 Redis Cluster）：

- 测试文件：`redowl/integration_test.go`
- build tag：`rediscluster`
- 环境变量：`REDIS_CLUSTER_ADDRS`（逗号分隔的节点地址）

该用例的消费者侧会先订阅 namespace channel：`{Prefix}:events`，从事件里的 `Queue` 字段动态发现队列并启动 per-queue worker（保证队列内顺序、队列间并发）。

运行方式示例（Windows PowerShell）：

```powershell
$env:REDIS_CLUSTER_ADDRS = "127.0.0.1:7000,127.0.0.1:7001,127.0.0.1:7002"
go test -tags=rediscluster ./redowl -run ClusterClient
```

运行方式示例（bash）：

```bash
REDIS_CLUSTER_ADDRS=127.0.0.1:7000,127.0.0.1:7001,127.0.0.1:7002 \
  go test -tags=rediscluster ./redowl -run ClusterClient
```

该测试复用了与本文档“多队列并发消费示例（Game*）”一致的语义验证：

- A 进程（producer）向 `Game1..Game14` 分别按序生产 `A1..A5` 到 `N1..N5`
- B 进程（consumer）同时监听所有 `Game*` 队列
	- 每个队列内严格顺序消费（上一条处理完成后再取下一条）
	- 不同队列之间并发消费
	- 每条消息处理在独立 goroutine 中执行

为避免在共享 Cluster 环境中相互影响，测试会为每次运行生成唯一的 `Prefix` 来隔离 keys。

### 死信队列（DLQ）

- 通过 `WithMaxReceiveCount(n)` 开启：当消息被投递次数 `ReceiveCount` 超过 n，会进入 DLQ。
- DLQ 消费：`ReceiveDLQ/ReceiveDLQWithWait`。
- DLQ 返回的消息同样带 `ReceiptHandle`，可直接用 `Ack` 删除消息。

```go
dlqMsg, _ := q.ReceiveDLQ(ctx)
if dlqMsg != nil {
	// 记录/报警/手动处理...
	_ = q.Ack(ctx, dlqMsg.ReceiptHandle)
}
```

### DLQ Redrive（重新投递）

将 DLQ 中的消息重新放回 ready：

```go
// 把最多 100 条 DLQ 消息放回 ready
moved, err := q.RedriveDLQ(ctx, 100)
_ = moved
_ = err
```

