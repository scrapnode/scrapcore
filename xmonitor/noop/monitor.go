package noop

import (
	"context"
	"github.com/scrapnode/scrapcore/xmonitor"
	"github.com/scrapnode/scrapcore/xmonitor/propergator"
)

type Monitor struct {
	propergator *Propergator
}

func (monitor *Monitor) Propergator() propergator.Propergator {
	return monitor.propergator
}
func (monitor *Monitor) Connect(ctx context.Context) error    { return nil }
func (monitor *Monitor) Disconnect(ctx context.Context) error { return nil }
func (monitor *Monitor) Trace(ctx context.Context, ns, name string) (context.Context, xmonitor.Span) {
	return ctx, &Span{}
}
func (monitor *Monitor) Record(ctx context.Context, ns, name string, incr int64) {}
func (monitor *Monitor) Count(ctx context.Context, ns, name string, incr int64)  {}
