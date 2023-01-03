package pipeline

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

func UseRecovery(logger *zap.SugaredLogger) Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (rctx context.Context, err error) {
			defer func() {
				if r := recover(); r != nil {
					logger.Error(r)
					rctx = ctx
					err = errors.New("pipeline: oops, something went wrong")
				}
			}()

			rctx, err = next(ctx)
			return
		}
	}
}
