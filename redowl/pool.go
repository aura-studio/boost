package redowl

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type WorkerPool struct {
	client       redis.UniversalClient
	prefix       string
	workerCount  int
	handler      func(context.Context, string, *Message) error
	idleTimeout  time.Duration
	pollInterval time.Duration

	mu      sync.RWMutex
	queues  map[string]*queueState
	workCh  chan string
	stopCh  chan struct{}
	wg      sync.WaitGroup
	pubsub  *redis.PubSub
}

type queueState struct {
	name       string
	lastActive time.Time
	msgCount   int64
}

type PoolOption func(*WorkerPool)

func WithWorkerCount(n int) PoolOption {
	return func(p *WorkerPool) { p.workerCount = n }
}

func WithIdleTimeout(d time.Duration) PoolOption {
	return func(p *WorkerPool) { p.idleTimeout = d }
}

func WithPollInterval(d time.Duration) PoolOption {
	return func(p *WorkerPool) { p.pollInterval = d }
}

func NewWorkerPool(
	client redis.UniversalClient,
	prefix string,
	handler func(context.Context, string, *Message) error,
	opts ...PoolOption,
) *WorkerPool {
	p := &WorkerPool{
		client:       client,
		prefix:       prefix,
		workerCount:  10,
		handler:      handler,
		idleTimeout:  5 * time.Minute,
		pollInterval: 500 * time.Millisecond,
		queues:       make(map[string]*queueState),
		workCh:       make(chan string, 1000),
		stopCh:       make(chan struct{}),
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func (p *WorkerPool) Start(ctx context.Context) error {
	p.pubsub = p.client.Subscribe(ctx, p.prefix+":events")
	if _, err := p.pubsub.Receive(ctx); err != nil {
		return err
	}

	// 启动事件监听器
	p.wg.Add(1)
	go p.eventListener(ctx)

	// 启动 worker 池
	for i := 0; i < p.workerCount; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}

	// 启动空闲队列清理器
	p.wg.Add(1)
	go p.idleCleaner(ctx)

	return nil
}

func (p *WorkerPool) Stop() {
	close(p.stopCh)
	if p.pubsub != nil {
		_ = p.pubsub.Close()
	}
	p.wg.Wait()
}

func (p *WorkerPool) eventListener(ctx context.Context) {
	defer p.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case <-p.stopCh:
			return
		case msg, ok := <-p.pubsub.Channel():
			if !ok {
				return
			}

			var ev Event
			if err := json.Unmarshal([]byte(msg.Payload), &ev); err != nil {
				continue
			}

			if ev.Type == EventSent && ev.Queue != "" {
				p.addQueue(ev.Queue)
			}
		}
	}
}

func (p *WorkerPool) addQueue(name string) {
	p.mu.Lock()
	if _, exists := p.queues[name]; !exists {
		p.queues[name] = &queueState{
			name:       name,
			lastActive: time.Now(),
		}
	}
	p.mu.Unlock()

	// 非阻塞地通知 worker
	select {
	case p.workCh <- name:
	default:
		// channel 满了，说明已经有足够的任务在等待
	}
}

func (p *WorkerPool) worker(ctx context.Context, id int) {
	defer p.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case <-p.stopCh:
			return
		case queueName := <-p.workCh:
			p.processQueue(ctx, queueName)
		}
	}
}

func (p *WorkerPool) processQueue(ctx context.Context, queueName string) {
	q, err := New(p.client, queueName,
		WithPrefix(p.prefix),
		WithVisibilityTimeout(30*time.Second),
	)
	if err != nil {
		return
	}
	defer q.StopReaper()

	idleCount := 0
	const maxIdle = 3 // 连续 3 次空闲则释放

	for {
		select {
		case <-ctx.Done():
			return
		case <-p.stopCh:
			return
		default:
		}

		msg, err := q.ReceiveWithWait(ctx, p.pollInterval)
		if err != nil {
			return
		}

		if msg == nil {
			idleCount++
			if idleCount >= maxIdle {
				// 连续空闲，释放 worker
				return
			}
			continue
		}

		// 有消息，重置空闲计数
		idleCount = 0
		p.updateQueueActivity(queueName)

		if err := p.handler(ctx, queueName, msg); err == nil {
			_ = q.Ack(ctx, msg.ReceiptHandle)
		}
	}
}

func (p *WorkerPool) updateQueueActivity(name string) {
	p.mu.Lock()
	if qs, ok := p.queues[name]; ok {
		qs.lastActive = time.Now()
		qs.msgCount++
	}
	p.mu.Unlock()
}

func (p *WorkerPool) idleCleaner(ctx context.Context) {
	defer p.wg.Done()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-p.stopCh:
			return
		case <-ticker.C:
			p.cleanIdleQueues()
		}
	}
}

func (p *WorkerPool) cleanIdleQueues() {
	now := time.Now()
	p.mu.Lock()
	for name, qs := range p.queues {
		if now.Sub(qs.lastActive) > p.idleTimeout {
			delete(p.queues, name)
		}
	}
	p.mu.Unlock()
}

func (p *WorkerPool) Stats() map[string]int64 {
	p.mu.RLock()
	defer p.mu.RUnlock()

	stats := make(map[string]int64)
	for name, qs := range p.queues {
		stats[name] = qs.msgCount
	}
	return stats
}
