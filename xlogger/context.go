package xlogger

import (
	"context"
	"errors"
)

type ctxkey string

const CTXKEY ctxkey = "xlogger"

func WithContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, CTXKEY, logger)
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(CTXKEY).(*Logger)
	if !ok {
		panic(errors.New("no logger was configured"))
	}

	return logger
}
