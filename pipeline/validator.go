package pipeline

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/scrapnode/scrapcore/auth"
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

func UseWorkspaceValidator() Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			id, ok := ctx.Value(CTXKEY_WS).(string)
			if !ok {
				return ctx, errors.New("no workspace is associated with request")
			}

			account, ok := ctx.Value(CTXKEY_ACC).(*auth.Account)
			if !ok {
				return ctx, errors.New("no account is associated with request")
			}

			if !account.OwnWorkspace(id) {
				return ctx, errors.New(fmt.Sprintf("workspace %s is not belong to account %s", account.Id, id))
			}

			return next(ctx)
		}
	}
}
