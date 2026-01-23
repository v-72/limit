package ratelimit

type Limiter interface {
    Allow(ctx context.Context, key string) (Result, error)
}

type Result struct {
    Allowed     bool
    Remaining   int64
    RetryAfter  time.Duration
}

type Config struct {
    Capacity   int64
    RefillRate int64 // tokens per second
    Shards     int
    Now        func() time.Time
}

func New(cfg Config) Limiter

