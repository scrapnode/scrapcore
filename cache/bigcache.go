package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"sync"
	"time"
)

func NewBigCache(ctx context.Context, cfg *Configs) (Cache, error) {
	opts := bigcache.DefaultConfig(time.Duration(cfg.SecondsToLive) * time.Second)
	client, err := bigcache.New(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &BigCache{cfg: cfg, client: client}, nil
}

type BigCache struct {
	cfg    *Configs
	client *bigcache.BigCache
	mu     sync.Mutex
}

func (cache *BigCache) Client() interface{} {
	return cache.client
}

func (cache *BigCache) Get(ctx context.Context, key string) ([]byte, error) {
	return cache.client.Get(key)
}

func (cache *BigCache) Set(ctx context.Context, key string, value []byte) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	return cache.client.Set(key, value)
}

func (cache *BigCache) Del(ctx context.Context, key string) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	return cache.client.Delete(key)
}

func (cache *BigCache) Exists(ctx context.Context, key string) bool {
	_, res, err := cache.client.GetWithInfo(key)
	if err != nil {
		return false
	}
	if res.EntryStatus != 0 {
		return false
	}
	return true
}

func (cache *BigCache) Incr(ctx context.Context, key string) (int64, error) {
	return cache.incr(ctx, key, 1)
}

func (cache *BigCache) Decr(ctx context.Context, key string) (int64, error) {
	return cache.incr(ctx, key, -1)
}

func (cache *BigCache) incr(ctx context.Context, key string, change int64) (int64, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	val, err := cache.client.Get(key)
	if err != nil {
		return 0, err
	}

	value, err := Decode[int64](val)
	if err != nil {
		return 0, err
	}

	newValue := *value + change

	bytes, err := Encode(newValue)
	if err != nil {
		return 0, err
	}
	if err := cache.client.Set(key, bytes); err != nil {
		return 0, err
	}

	return newValue, nil
}
