package redowl

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestWorkerPool(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	prefix := fmt.Sprintf("test_pool_%d", time.Now().UnixNano())

	// 创建 worker 池：10 个 worker 处理无限多个队列
	pool := NewWorkerPool(
		client,
		prefix,
		func(ctx context.Context, queueName string, msg *Message) error {
			fmt.Printf("[%s] Processing: %s\n", queueName, string(msg.Body))
			time.Sleep(10 * time.Millisecond)
			return nil
		},
		WithWorkerCount(10),
		WithIdleTimeout(2*time.Minute),
		WithPollInterval(500*time.Millisecond),
	)

	if err := pool.Start(ctx); err != nil {
		t.Fatalf("start pool: %v", err)
	}
	defer pool.Stop()

	// 模拟 100 个队列发送消息
	for i := 0; i < 100; i++ {
		queueName := fmt.Sprintf("queue_%d", i)
		q, err := New(client, queueName,
			WithPrefix(prefix),
			WithVisibilityTimeout(30*time.Second),
		)
		if err != nil {
			t.Fatalf("create queue: %v", err)
		}

		for j := 0; j < 5; j++ {
			_, err := q.Send(ctx, []byte(fmt.Sprintf("msg_%d", j)), nil)
			if err != nil {
				t.Fatalf("send: %v", err)
			}
		}
	}

	// 等待处理完成
	time.Sleep(10 * time.Second)

	// 查看统计
	stats := pool.Stats()
	fmt.Printf("Processed messages by queue:\n")
	for name, count := range stats {
		fmt.Printf("  %s: %d\n", name, count)
	}
}

func ExampleWorkerPool() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	ctx := context.Background()

	// 10 个 worker 处理所有队列
	pool := NewWorkerPool(
		client,
		"myapp",
		func(ctx context.Context, queueName string, msg *Message) error {
			fmt.Printf("Queue %s: %s\n", queueName, string(msg.Body))
			return nil
		},
		WithWorkerCount(10),
		WithIdleTimeout(5*time.Minute),
	)

	pool.Start(ctx)
	defer pool.Stop()

	// 动态创建队列并发送消息
	for i := 0; i < 1000; i++ {
		q, _ := New(client, fmt.Sprintf("game_%d", i),
			WithPrefix("myapp"),
		)
		q.Send(ctx, []byte("player_joined"), nil)
	}

	// 10 个 worker 会自动处理所有 1000 个队列的消息
}
