package pipeline

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/auth"
)

func UseAccountValidator() Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			if _, ok := ctx.Value(CTXKEY_ACCOUNT).(*auth.Account); !ok {
				return ctx, errors.New("no account is associated with request")
			}

			return next(ctx)
		}
	}
}
