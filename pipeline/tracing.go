package pipeline

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

func UseTracing(name, spanname string, pipeline Pipeline) Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ctx, span := otel.Tracer(name).Start(ctx, spanname)

			// use fake next function to make sure we don't trace subchain pipeline execute time
			ctx, err := pipeline(func(ctx context.Context) (context.Context, error) {
				return ctx, nil
			})(ctx)
			if err != nil {
				span.SetStatus(codes.Error, err.Error())
			}

			// don't use defer span.End()
			// becuase it will measure subchain pipeline execute time too
			span.End()
			return next(ctx)
		}
	}
}
