package xcache

import "context"

type ctxkey string

const (
	OPT_TTL ctxkey = "xcache.options.ttl"
	CACHE   ctxkey = "xcache"
)

func WithContext(ctx context.Context, cache Cache) context.Context {
	return context.WithValue(ctx, CACHE, cache)
}

func FromContext(ctx context.Context) Cache {
	return ctx.Value(CACHE).(Cache)
}
