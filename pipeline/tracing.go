package pipeline

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type TracingConfigs struct {
	DryRun    bool
	TraceName string
	SpanName  string
}

const CTXKEY_TRACE_ATTRS ctxkey = "pipeline.traces.attributes"

func WithTraceAttributes(ctx context.Context, kv ...any) context.Context {
	if len(kv) < 2 {
		return ctx
	}
	if len(kv)%2 != 0 {
		kv = kv[:len(kv)-2]
	}

	attributes := map[string]interface{}{}
	for i := 0; i < len(kv)-1; i += 2 {
		key, ok := kv[i].(string)
		if !ok {
			continue
		}
		value := kv[i+1]
		attributes[key] = value
	}

	return context.WithValue(ctx, CTXKEY_TRACE_ATTRS, attributes)
}

func UseTracing(pipeline Pipeline, cfg *TracingConfigs) Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			if cfg.DryRun {
				return next(ctx)
			}

			ctx, span := otel.Tracer(cfg.TraceName).Start(ctx, cfg.SpanName)

			// use fake next function to make sure we don't trace sub-chain pipeline execute time
			ctx, err := pipeline(func(ctx context.Context) (context.Context, error) {
				return ctx, nil
			})(ctx)

			if traces, ok := ctx.Value(CTXKEY_TRACE_ATTRS).(map[string]interface{}); ok {
				for key, v := range traces {
					if value, ok := v.(string); ok {
						span.SetAttributes(attribute.String(key, value))
					}
					if value, ok := v.(int64); ok {
						span.SetAttributes(attribute.Int64(key, value))
					}
					if value, ok := v.(float64); ok {
						span.SetAttributes(attribute.Float64(key, value))
					}
					if value, ok := v.(bool); ok {
						span.SetAttributes(attribute.Bool(key, value))
					}
				}
			}

			// return error as soon as we got it
			if err != nil {
				span.SetStatus(codes.Error, err.Error())
				// don't use defer span.End()
				// because it will measure sub-chain pipeline execute time too
				span.End()
				return ctx, err
			}

			span.SetStatus(codes.Ok, fmt.Sprintf("%s.%s: ok", cfg.TraceName, cfg.SpanName))
			span.End()
			return next(ctx)
		}
	}
}
