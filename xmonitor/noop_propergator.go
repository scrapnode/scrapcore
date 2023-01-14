package xmonitor

import "context"

type NoopPropergator struct{}

func (propergator *NoopPropergator) Extract(ctx context.Context) map[string]string {
	return map[string]string{}
}
func (propergator *NoopPropergator) Inject(ctx context.Context, carrier map[string]string) context.Context {
	return ctx
}
