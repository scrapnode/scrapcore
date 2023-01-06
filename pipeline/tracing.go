package pipeline

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
)

type TracingConfigs struct {
	TraceName string
	SpanName  string
}

func UseTracing(pipeline Pipeline, cfg *TracingConfigs) Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ctx, span := otel.Tracer(cfg.TraceName).Start(ctx, cfg.SpanName)

			// use fake next function to make sure we don't trace subchain pipeline execute time
			ctx, err := pipeline(func(ctx context.Context) (context.Context, error) {
				return ctx, nil
			})(ctx)
			if err == nil {
				span.SetStatus(codes.Ok, fmt.Sprintf("%s.%s: ok", cfg.TraceName, cfg.SpanName))
			} else {
				span.SetStatus(codes.Error, err.Error())
			}

			// don't use defer span.End()
			// becuase it will measure subchain pipeline execute time too
			span.End()
			return next(ctx)
		}
	}
}

func UseTracingPropagator(key string) Pipeline {
	return func(next Pipe) Pipe {
		propagator := otel.GetTextMapPropagator()
		return func(ctx context.Context) (context.Context, error) {
			value := utils.StructValueByKey(ctx.Value(CTXKEY_REQ), key)
			if value == nil {
				return next(ctx)
			}

			metadata, ok := value.(map[string]string)
			if !ok {
				return next(ctx)
			}

			carier := propagation.MapCarrier(metadata)
			ctx = propagator.Extract(ctx, carier)
			return next(ctx)
		}
	}
}
