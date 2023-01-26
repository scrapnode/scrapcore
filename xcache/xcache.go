package xcache

import (
	"context"
	"net/url"
	"strings"
)

type Cache interface {
	Client() interface{}
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	Del(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) bool
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
}

type Configs struct {
	Dsn           string `json:"dsn"`
	SecondsToLive int64  `json:"seconds_to_live"`
}

func New(ctx context.Context, cfg *Configs) (Cache, error) {
	uri, err := url.Parse(cfg.Dsn)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(uri.Scheme, "redis") {
		return NewRedis(ctx, cfg)
	}
	return NewBigCache(ctx, cfg)
}
