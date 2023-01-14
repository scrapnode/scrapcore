package noop

import "context"

type Propergator struct{}

func (propergator *Propergator) Extract(ctx context.Context) map[string]string {
	return map[string]string{}
}
func (propergator *Propergator) Inject(ctx context.Context, carrier map[string]string) context.Context {
	return ctx
}
