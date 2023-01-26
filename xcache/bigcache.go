package xcache

import (
	"context"
	"errors"
	"github.com/allegro/bigcache/v3"
	"github.com/scrapnode/scrapcore/xlogger"
	"go.uber.org/zap"
	"sync"
	"time"
)

func NewBigCache(ctx context.Context, cfg *Configs) (Cache, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "xcache.bigcache")
	return &BigCache{cfg: cfg, logger: logger}, nil
}

type BigCache struct {
	cfg    *Configs
	logger *zap.SugaredLogger
	client *bigcache.BigCache
	mu     sync.Mutex
}

func (cache *BigCache) Client() interface{} {
	return cache.client
}

func (cache *BigCache) Connect(ctx context.Context) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	eviction := time.Duration(cache.cfg.SecondsToLive) * time.Second
	opts := bigcache.DefaultConfig(eviction)
	client, err := bigcache.New(ctx, opts)
	if err != nil {
		return err
	}

	cache.client = client
	cache.logger.Info("connected")
	return nil
}

func (cache *BigCache) Disconnect(ctx context.Context) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	err := cache.client.Close()
	cache.logger.Info("disconnected")
	return err
}

func (cache *BigCache) Set(ctx context.Context, key string, value []byte) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	return cache.client.Set(key, value)
}

func (cache *BigCache) Get(ctx context.Context, key string) ([]byte, error) {
	return cache.client.Get(key)
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

	var value int64

	val, err := cache.client.Get(key)
	// should return error if it is not entry not found error
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return 0, err
	}
	if err == nil {
		valueptr, err := Decode[int64](val)
		if err != nil {
			return 0, err
		}
		value = *valueptr
	}

	newValue := value + change

	bytes, err := Encode(newValue)
	if err != nil {
		return 0, err
	}
	if err := cache.client.Set(key, bytes); err != nil {
		return 0, err
	}

	return newValue, nil
}
