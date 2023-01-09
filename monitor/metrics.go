package monitor

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"time"
)

func NewMetrics(ctx context.Context, cfg *Configs) Monitor {
	return &Metrics{cfg: cfg}
}

type Metrics struct {
	cfg      *Configs
	exporter metric.Exporter
}

func (metrics *Metrics) Connect(ctx context.Context) error {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNamespaceKey.String(metrics.cfg.Namespace),
			semconv.ServiceNameKey.String(metrics.cfg.Name),
		),
	)
	if err != nil {
		return err
	}

	exporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(metrics.cfg.Metrics.Endpoint))
	if err != nil {
		return err
	}

	metrics.exporter = exporter
	provider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(
			metric.NewPeriodicReader(
				exporter,
				metric.WithInterval(2*time.Second),
			),
		),
	)
	global.SetMeterProvider(provider)

	return nil
}

func (metrics *Metrics) Disconnect(ctx context.Context) error {
	if metrics.exporter == nil {
		return nil
	}

	return metrics.exporter.Shutdown(ctx)
}
