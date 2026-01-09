package redowl

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestWorkerPool_DynamicRelease(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	prefix := fmt.Sprintf("test_dynamic_%d", time.Now().UnixNano())

	var processed int64
	var activeWorkers int64

	pool := NewWorkerPool(
		client,
		prefix,
		func(ctx context.Context, queueName string, msg *Message) error {
			atomic.AddInt64(&activeWorkers, 1)
			defer atomic.AddInt64(&activeWorkers, -1)

			atomic.AddInt64(&processed, 1)
			time.Sleep(50 * time.Millisecond)
			return nil
		},
		WithWorkerCount(5),
		WithPollInterval(200*time.Millisecond),
	)

	require.NoError(t, pool.Start(ctx))
	defer pool.Stop()

	// 阶段 1: 发送 20 个队列，每个 5 条消息
	t.Log("Phase 1: Sending to 20 queues")
	for i := 0; i < 20; i++ {
		q, err := New(client, fmt.Sprintf("wave1_q%d", i),
			WithPrefix(prefix),
			WithVisibilityTimeout(5*time.Second),
		)
		require.NoError(t, err)

		for j := 0; j < 5; j++ {
			_, err := q.Send(ctx, []byte(fmt.Sprintf("msg_%d", j)), nil)
			require.NoError(t, err)
		}
	}

	// 等待第一波处理完成
	time.Sleep(5 * time.Second)
	phase1Processed := atomic.LoadInt64(&processed)
	t.Logf("Phase 1 processed: %d messages", phase1Processed)
	require.Equal(t, int64(100), phase1Processed)

	// 记录 goroutine 数量
	goroutinesBefore := runtime.NumGoroutine()
	t.Logf("Goroutines after phase 1: %d", goroutinesBefore)

	// 阶段 2: 等待 worker 释放（连续 3 次空闲）
	t.Log("Phase 2: Waiting for workers to release...")
	time.Sleep(3 * time.Second)

	goroutinesAfterIdle := runtime.NumGoroutine()
	t.Logf("Goroutines after idle: %d", goroutinesAfterIdle)

	// 阶段 3: 发送新的 20 个队列（不同的队列名）
	t.Log("Phase 3: Sending to new 20 queues")
	for i := 0; i < 20; i++ {
		q, err := New(client, fmt.Sprintf("wave2_q%d", i),
			WithPrefix(prefix),
			WithVisibilityTimeout(5*time.Second),
		)
		require.NoError(t, err)

		for j := 0; j < 5; j++ {
			_, err := q.Send(ctx, []byte(fmt.Sprintf("msg_%d", j)), nil)
			require.NoError(t, err)
		}
	}

	// 等待第二波处理完成
	time.Sleep(5 * time.Second)
	phase2Processed := atomic.LoadInt64(&processed)
	t.Logf("Phase 2 processed: %d messages", phase2Processed)
	require.Equal(t, int64(200), phase2Processed)

	goroutinesAfterPhase2 := runtime.NumGoroutine()
	t.Logf("Goroutines after phase 2: %d", goroutinesAfterPhase2)

	// 验证：worker 能够复用，goroutine 数量不会无限增长
	// 允许一定的波动，但不应该翻倍
	require.Less(t, goroutinesAfterPhase2, goroutinesBefore*2,
		"Goroutines should not double after processing new queues")
}

func TestWorkerPool_RollingQueues(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	prefix := fmt.Sprintf("test_rolling_%d", time.Now().UnixNano())

	var processed int64

	pool := NewWorkerPool(
		client,
		prefix,
		func(ctx context.Context, queueName string, msg *Message) error {
			atomic.AddInt64(&processed, 1)
			time.Sleep(10 * time.Millisecond)
			return nil
		},
		WithWorkerCount(10),
		WithPollInterval(100*time.Millisecond),
	)

	require.NoError(t, pool.Start(ctx))
	defer pool.Stop()

	// 模拟滚动队列：每秒创建 10 个新队列，每个队列 3 条消息
	// 持续 5 秒，总共 50 个队列，150 条消息
	t.Log("Simulating rolling queues...")
	for wave := 0; wave < 5; wave++ {
		for i := 0; i < 10; i++ {
			qName := fmt.Sprintf("rolling_w%d_q%d", wave, i)
			q, err := New(client, qName,
				WithPrefix(prefix),
				WithVisibilityTimeout(5*time.Second),
			)
			require.NoError(t, err)

			for j := 0; j < 3; j++ {
				_, err := q.Send(ctx, []byte(fmt.Sprintf("msg_%d", j)), nil)
				require.NoError(t, err)
			}
		}
		time.Sleep(1 * time.Second)
	}

	// 等待所有消息处理完成
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		if atomic.LoadInt64(&processed) >= 150 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	finalProcessed := atomic.LoadInt64(&processed)
	t.Logf("Final processed: %d messages", finalProcessed)
	require.Equal(t, int64(150), finalProcessed)

	stats := pool.Stats()
	t.Logf("Tracked queues: %d", len(stats))

	// 验证：10 个 worker 能够处理 50 个滚动队列
	require.GreaterOrEqual(t, len(stats), 40, "Should track most queues")
}
