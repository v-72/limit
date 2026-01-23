package ratelimit

import (
	"context"
	"testing"
)

func BenchmarkLimiter_Allow(b *testing.B) {
	l := New(Config{
		Capacity:   100,
		RefillRate: 10,
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = l.Allow(context.Background(), "user")
		}
	})
}

