package ratelimit

import (
	"context"
	"testing"
	"time"
)

func TestLimiter_Basic(t *testing.T) {
	now := time.Now()
	cfg := Config{
		Capacity:   2,
		RefillRate: 1,
		Now: func() time.Time {
			return now
		},
	}

	l := New(cfg)

	res, _ := l.Allow(context.Background(), "k")
	if !res.Allowed {
		t.Fatal("expected allowed")
	}

	res, _ = l.Allow(context.Background(), "k")
	if !res.Allowed {
		t.Fatal("expected allowed")
	}

	res, _ = l.Allow(context.Background(), "k")
	if res.Allowed {
		t.Fatal("expected rate limited")
	}

	now = now.Add(time.Second)

	res, _ = l.Allow(context.Background(), "k")
	if !res.Allowed {
		t.Fatal("expected refill")
	}
}

