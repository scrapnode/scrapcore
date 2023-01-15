package pipeline

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/xmonitor"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"time"
)

func UseTracing(pipeline Pipeline, monitor xmonitor.Monitor, ns, name string) Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ctx, span := monitor.Trace(ctx, ns, name)

			start := time.Now()
			// use fake next function to make sure we don't trace sub-chain pipeline execute time
			ctx, err := pipeline(func(ctx context.Context) (context.Context, error) {
				return ctx, nil
			})(ctx)
			span.SetAttributes(attributes.FromContext(ctx))

			// return error as soon as we got it
			if err != nil {
				span.KO(err.Error())
				// don't use defer span.End()
				// because it will measure sub-chain pipeline execute time too
				span.End()
				return ctx, err
			}

			span.OK(fmt.Sprintf("%s/%s: %dms", ns, name, time.Since(start).Milliseconds()))
			span.End()
			return next(ctx)
		}
	}
}

func UseMetrics(monitor xmonitor.Monitor, ns, name string) Pipeline {
	return func(next Pipe) Pipe {
		return func(ctx context.Context) (context.Context, error) {
			start := time.Now()

			ctx, err := next(ctx)

			duration := time.Since(start)
			monitor.Record(ctx, ns, name, duration.Milliseconds())

			return ctx, err
		}
	}
}
