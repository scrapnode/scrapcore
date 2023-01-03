package pipeline

import (
	"context"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func UseValidator(logger *zap.SugaredLogger) Pipeline {
	return func(next Pipe) Pipe {
		validate := validator.New()

		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(CTXKEY_REQ)
			err := validate.Struct(req)
			if err != nil {
				if errors, ok := err.(validator.ValidationErrors); !ok {
					return context.WithValue(ctx, CTXKEY_ERR, errors), nil
				}

				return ctx, err
			}

			return next(ctx)
		}
	}
}
