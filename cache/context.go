package cache

import "context"

type ctxkey string

const (
	OPT_TTL ctxkey = "cache.options.ttl"
	CACHE   ctxkey = "cache"
)

func WithContext(ctx context.Context, cache Cache) context.Context {
	return context.WithValue(ctx, CACHE, cache)
}

func FromContext(ctx context.Context) Cache {
	return ctx.Value(CACHE).(Cache)
}
