package monitor

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func NewTracer(ctx context.Context, cfg *Configs) Monitor {
	return &Tracer{cfg: cfg}
}

type Tracer struct {
	cfg      *Configs
	exporter *otlptrace.Exporter
}

func (tracer *Tracer) Connect(ctx context.Context) error {
	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNamespaceKey.String(tracer.cfg.Namespace),
			semconv.ServiceNameKey.String(tracer.cfg.Name),
		),
	)
	if err != nil {
		return err
	}

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(tracer.cfg.Tracer.Endpoint),
	)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return err
	}
	tracer.exporter = exporter

	bsp := trace.NewBatchSpanProcessor(exporter)
	provider := trace.NewTracerProvider(
		trace.WithSampler(trace.TraceIDRatioBased(tracer.cfg.Tracer.Ratio)),
		trace.WithResource(res),
		trace.WithSpanProcessor(bsp),
	)

	// set global propagator to trace context (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(provider)
	return nil
}

func (tracer *Tracer) Disconnect(ctx context.Context) error {
	if tracer.exporter == nil {
		return nil
	}

	return tracer.exporter.Shutdown(ctx)
}
