package ratelimit

import (
	"context"
	"time"
)

type Limiter interface {
	Allow(ctx context.Context, key string) (Result, error)
}

type Result struct {
	Allowed    bool
	Remaining  int64
	RetryAfter time.Duration
}

type limiter struct {
	cfg   Config
	store *store
}

func New(cfg Config) Limiter {
	if err := cfg.normalize(); err != nil {
		panic(err)
	}

	return &limiter{
		cfg:   cfg,
		store: newStore(cfg.Shards),
	}
}

func (l *limiter) Allow(ctx context.Context, key string) (Result, error) {
	select {
	case <-ctx.Done():
		return Result{}, ctx.Err()
	default:
	}

	now := l.cfg.Now()
	b := l.store.get(key, l.cfg, now)

	b.mu.Lock()
	defer b.mu.Unlock()

	b.refill(now)

	if b.tokens <= 0 {
		retry := time.Second / time.Duration(l.cfg.RefillRate)
		return Result{
			Allowed:    false,
			Remaining:  0,
			RetryAfter: retry,
		}, nil
	}

	b.tokens--

	return Result{
		Allowed:   true,
		Remaining: b.tokens,
	}, nil
}

