package xcache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/scrapnode/scrapcore/xlogger"
	"go.uber.org/zap"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

func NewRedis(ctx context.Context, cfg *Configs) (Cache, error) {
	logger := xlogger.FromContext(ctx).With("pkg", "xcache.redis")
	return &Redis{cfg: cfg, logger: logger}, nil
}

type Redis struct {
	cfg    *Configs
	logger *zap.SugaredLogger
	client *redis.Client
	mu     sync.Mutex
}

func (cache *Redis) Client() interface{} {
	return cache.client
}

func (cache *Redis) Connect(ctx context.Context) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	uri, err := url.Parse(cache.cfg.Dsn)
	if err != nil {
		return err
	}

	opts := &redis.Options{
		Addr:     uri.Host,
		DB:       0,
		Username: uri.User.Username(),
	}

	dbstr := strings.Replace(uri.Path, "/", "", -1)
	if dbstr != "" {
		db, err := strconv.Atoi(dbstr)
		if err == nil {
			opts.DB = db
		} else {
			cache.logger.Errorw("could not parse db number, use 0 as default", "error", err.Error())
		}
	}
	if password, ok := uri.User.Password(); ok {
		opts.Password = password
	}
	cache.client = redis.NewClient(opts)

	cache.logger.Info("connected")
	return nil
}

func (cache *Redis) Disconnect(ctx context.Context) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if cache.client == nil {
		return nil
	}

	err := cache.client.Close()
	cache.logger.Info("disconnected")
	return err
}

func (cache *Redis) Set(ctx context.Context, key string, value []byte) error {
	return cache.client.Set(ctx, key, value, cache.STL(ctx)).Err()
}

func (cache *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	return cache.client.Get(ctx, key).Bytes()
}

func (cache *Redis) Del(ctx context.Context, key string) error {
	return cache.client.Del(ctx, key).Err()
}

func (cache *Redis) Exists(ctx context.Context, key string) bool {
	val, err := cache.client.Exists(ctx, key).Result()
	return err == nil && val == 1
}

func (cache *Redis) Incr(ctx context.Context, key string) (int64, error) {
	count, err := cache.client.Incr(ctx, key).Result()
	// reset expire time when we increase counter
	if err == nil {
		if err := cache.client.Expire(ctx, key, cache.STL(ctx)).Err(); err != nil {
			cache.logger.Errorw("could not reset expire time after increased", "key", key)
		}
	}

	return count, err
}

func (cache *Redis) Decr(ctx context.Context, key string) (int64, error) {
	count, err := cache.client.Decr(ctx, key).Result()
	// reset expire time when we decrease counter
	if err == nil {
		if err := cache.client.Expire(ctx, key, cache.STL(ctx)).Err(); err != nil {
			cache.logger.Errorw("could not reset expire time after decreased", "key", key)
		}
	}

	return count, err
}

func (cache *Redis) STL(ctx context.Context) time.Duration {
	if ttl := ctx.Value(OPT_TTL); ttl != nil {
		return ttl.(time.Duration)
	}

	return time.Duration(cache.cfg.SecondsToLive) * time.Second
}
