package ratelimit

import (
	"sync"
	"time"
)

type bucket struct {
	mu         sync.Mutex
	capacity   int64
	tokens     int64
	refillRate int64
	lastRefill time.Time
}

func (b *bucket) refill(now time.Time) {
	elapsed := now.Sub(b.lastRefill).Seconds()
	if elapsed <= 0 {
		return
	}

	refill := int64(elapsed * float64(b.refillRate))
	if refill > 0 {
		b.tokens = min(b.capacity, b.tokens+refill)
		b.lastRefill = now
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

