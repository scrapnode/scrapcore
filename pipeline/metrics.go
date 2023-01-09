package pipeline

import (
	"context"
	"go.opentelemetry.io/otel/metric/global"
	"time"
)

type MetricsConfigs struct {
	DryRun              bool
	InstrumentationName string
	MetricName          string
}

func UseMetrics(cfg *MetricsConfigs) Pipeline {
	return func(next Pipe) Pipe {
		historam, err := global.Meter(cfg.InstrumentationName).SyncInt64().Histogram(cfg.MetricName)
		return func(ctx context.Context) (context.Context, error) {
			if cfg.DryRun {
				return next(ctx)
			}
			// drop this metrics because of init error
			if err != nil {
				return next(ctx)
			}

			start := time.Now()
			ctx, err := next(ctx)
			duration := time.Since(start)
			historam.Record(ctx, duration.Milliseconds())

			return ctx, err
		}
	}
}
