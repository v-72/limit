type Config struct {
    Capacity   int64
    RefillRate int64
    Shards     int
    Now        func() time.Time
}

func (c *Config) normalize() error {
    if c.Capacity <= 0 {
        return errors.New("capacity must be > 0")
    }
    if c.RefillRate <= 0 {
        return errors.New("refill rate must be > 0")
    }
    if c.Shards <= 0 {
        c.Shards = 64
    }
    if c.Now == nil {
        c.Now = time.Now
    }
    return nil
}

