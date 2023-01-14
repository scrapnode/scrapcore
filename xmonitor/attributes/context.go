package attributes

import "context"

type ctxkey string

const CTXKEY ctxkey = "xmonitor.attrbutes"

func WithContext(ctx context.Context, attributes Attributes) context.Context {
	return context.WithValue(ctx, CTXKEY, attributes)
}

func FromContext(ctx context.Context) Attributes {
	attributes, ok := ctx.Value(CTXKEY).(Attributes)

	if !ok {
		return Attributes{}
	}

	return attributes
}
