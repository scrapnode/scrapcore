package xconfig

import (
	"context"
	"github.com/spf13/viper"
)

type ctxkey string

const CTXKEY ctxkey = "xconfig.provider"

func WithContext(ctx context.Context, provider *viper.Viper) context.Context {
	ctx = context.WithValue(ctx, CTXKEY, provider)
	return ctx
}

func FromContext(ctx context.Context) *viper.Viper {
	return ctx.Value(CTXKEY).(*viper.Viper)
}
