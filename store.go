package ratelimit

import "sync"

type store struct {
	shards []shard
}

type shard struct {
	mu      sync.Mutex
	buckets map[string]*bucket
}

func newStore(n int) *store {
	s := &store{
		shards: make([]shard, n),
	}
	for i := range s.shards {
		s.shards[i].buckets = make(map[string]*bucket)
	}
	return s
}

func (s *store) get(key string, cfg Config, now time.Time) *bucket {
	sh := &s.shards[hash(key)%uint32(len(s.shards))]

	sh.mu.Lock()
	defer sh.mu.Unlock()

	b, ok := sh.buckets[key]
	if !ok {
		b = &bucket{
			capacity:   cfg.Capacity,
			tokens:     cfg.Capacity,
			refillRate: cfg.RefillRate,
			lastRefill: now,
		}
		sh.buckets[key] = b
	}
	return b
}

