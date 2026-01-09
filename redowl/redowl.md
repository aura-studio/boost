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

