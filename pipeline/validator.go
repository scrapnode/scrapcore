package pipeline

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
)

func UseValidator() Pipeline {
	return func(next Pipe) Pipe {
		validate := validator.New()

		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(CTXKEY_REQ)
			err := validate.Struct(req)
			if err != nil {
				if details, ok := err.(validator.ValidationErrors); !ok {
					return context.WithValue(ctx, CTXKEY_ERR, details), errors.New("validation is failed")
				}

				return ctx, err
			}

			return next(ctx)
		}
	}
}
