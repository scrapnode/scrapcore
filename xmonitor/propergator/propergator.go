package propergator

import "context"

type Propergator interface {
	Extract(ctx context.Context) map[string]string
	Inject(ctx context.Context, carrier map[string]string) context.Context
}
