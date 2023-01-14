package xmonitor

import (
	"context"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scrapcore/xmonitor/configs"
	"github.com/scrapnode/scrapcore/xmonitor/noop"
	"github.com/scrapnode/scrapcore/xmonitor/propergator"
)

func New(ctx context.Context, cfg *configs.Configs) (Monitor, error) {
	return noop.New(ctx, cfg)
}

type Monitor interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error

	Propergator() propergator.Propergator
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
