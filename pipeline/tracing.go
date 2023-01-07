package pipeline

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type TracingConfigs struct {
	TraceName string
	SpanName  string
}

func UseTracing(pipeline Pipeline, cfg *TracingConfigs) Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ctx, span := otel.Tracer(cfg.TraceName).Start(ctx, cfg.SpanName)

			// use fake next function to make sure we don't trace sub-chain pipeline execute time
			ctx, err := pipeline(func(ctx context.Context) (context.Context, error) {
				return ctx, nil
			})(ctx)

			// return error as soon as we got it
			if err != nil {
				span.SetStatus(codes.Error, err.Error())
				// don't use defer span.End()
				// because it will measure sub-chain pipeline execute time too
				span.End()
				return ctx, err
			}

			span.SetStatus(codes.Ok, fmt.Sprintf("%s.%s: ok", cfg.TraceName, cfg.SpanName))
			return next(ctx)
		}
	}
}
