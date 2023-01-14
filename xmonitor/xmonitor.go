package xmonitor

import (
	"context"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
)

func New(ctx context.Context, cfg *Configs) (Monitor, error) {
	return NewNoop(ctx, cfg)
}

type Monitor interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	Propergator() Propergator
	Trace(ctx context.Context, ns, name string) (context.Context, Span)
	Record(ctx context.Context, ns, name string, incr int64)
	Count(ctx context.Context, ns, name string, incr int64)
}

type Span interface {
	SetAttributes(attributes attributes.Attributes)
	OK(desc string)
	KO(desc string)
	End()
}

type Propergator interface {
	Extract(ctx context.Context) map[string]string
	Inject(ctx context.Context, carrier map[string]string) context.Context
}
