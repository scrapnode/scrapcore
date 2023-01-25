package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func NewRedis(ctx context.Context, cfg *Configs) (Cache, error) {
	uri, err := url.Parse(cfg.Dsn)
	if err != nil {
		return nil, err
	}

	db, err := strconv.Atoi(strings.Replace(uri.Path, "/", "", -1))
	if err != nil {
		return nil, err
	}

	opts := &redis.Options{
		Addr:     uri.Host,
		DB:       db,
		Username: uri.User.Username(),
	}
	if password, ok := uri.User.Password(); ok {
		opts.Password = password
	}

	return &Redis{cfg: cfg, client: redis.NewClient(opts)}, nil
}

type Redis struct {
	cfg    *Configs
	client *redis.Client
}

func (cache *Redis) Client() interface{} {
	return cache.client
}

func (cache *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	return cache.client.Get(ctx, key).Bytes()
}

func (cache *Redis) Set(ctx context.Context, key string, value []byte) error {
	stl := time.Duration(cache.cfg.SecondsToLive) * time.Second
	if ttl := ctx.Value(OPT_TTL); ttl != nil {
		stl = ttl.(time.Duration)
	}
	return cache.client.Set(ctx, key, value, stl).Err()
}

func (cache *Redis) Del(ctx context.Context, key string) error {
	return cache.client.Del(ctx, key).Err()
}

func (cache *Redis) Exists(ctx context.Context, key string) bool {
	val, err := cache.client.Exists(ctx, key).Result()
	return err == nil && val == 1
}

func (cache *Redis) Incr(ctx context.Context, key string) (int64, error) {
	return cache.client.Incr(ctx, key).Result()
}

func (cache *Redis) Decr(ctx context.Context, key string) (int64, error) {
	return cache.client.Incr(ctx, key).Result()
}
