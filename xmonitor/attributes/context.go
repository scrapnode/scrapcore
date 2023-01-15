package attributes

import "context"

type ctxkey string

const CTXKEY ctxkey = "xmonitor.attrbutes"

func WithContext(ctx context.Context, attributes Attributes) context.Context {
	attrs := FromContext(ctx)
	for key, value := range attributes {
		attrs[key] = value
	}
	return context.WithValue(ctx, CTXKEY, attrs)
}

func FromContext(ctx context.Context) Attributes {
	attributes, ok := ctx.Value(CTXKEY).(Attributes)

	if !ok {
		return Attributes{}
	}

	return attributes
}
