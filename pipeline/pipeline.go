package pipeline

import (
	"context"
)

type ctxkey string

const (
	CTXKEY_REQ ctxkey = "pipeline.request"
	CTXKEY_RES ctxkey = "pipeline.response"
	CTXKEY_ERR ctxkey = "pipeline.errors"
	CTXKEY_ACC ctxkey = "pipeline.account"
	CTXKEY_WS  ctxkey = "pipeline.workspace"
)

type Pipe func(ctx context.Context) (context.Context, error)
type Pipeline func(Pipe) Pipe

func New(pipelines []Pipeline) Pipe {
	pipe := func(ctx context.Context) (context.Context, error) {
		return ctx, nil
	}

	// we want to  get result in order of pipeline definition,
	// so we need to prepare it in reverse order
	for i := len(pipelines) - 1; i >= 0; i-- {
		pipe = pipelines[i](pipe)
	}

	return pipe
}

type BatchResult struct {
	Key   string
	Error string
}
