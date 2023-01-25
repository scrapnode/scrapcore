package cache

import "context"

func Get[T any](ctx context.Context, key string) (*T, error) {
	cache := FromContext(ctx)

	data, err := cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	return Decode[T](data)
}

func Set(ctx context.Context, key string, value any) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}

	return FromContext(ctx).Set(ctx, key, data)
}
