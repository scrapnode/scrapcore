package xmonitor

import "context"

type ctxkey string

const CTXKEY ctxkey = "xmonitor"

func WithContext(ctx context.Context, monitor Monitor) context.Context {
	return context.WithValue(ctx, CTXKEY, monitor)
}

func FromContext(ctx context.Context) Monitor {
	return ctx.Value(CTXKEY).(Monitor)
}
