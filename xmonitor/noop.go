package xmonitor

import (
	"context"
)

func NewNoop(ctx context.Context, cfg *Configs) (*Noop, error) {
	return &Noop{propergator: &NoopPropergator{}}, nil
}

type Noop struct {
	propergator *NoopPropergator
}

func (monitor *Noop) Propergator() Propergator {
	return monitor.propergator
}
func (monitor *Noop) Connect(ctx context.Context) error    { return nil }
func (monitor *Noop) Disconnect(ctx context.Context) error { return nil }
func (monitor *Noop) Trace(ctx context.Context, ns, name string) (context.Context, Span) {
	return ctx, &NoopSpan{}
}
func (monitor *Noop) Record(ctx context.Context, ns, name string, incr int64) {}
func (monitor *Noop) Count(ctx context.Context, ns, name string, incr int64)  {}
