package cache_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSet_Ok(t *testing.T) {
	ctx := setupFacade()

	key := gofakeit.Username()
	value := EncodeDecodeStruct{
		Username: gofakeit.Username(),
		Age:      gofakeit.IntRange(10, 30),
		Active:   gofakeit.Bool(),
	}
	err := cache.Set(ctx, key, value)
	assert.Nil(t, err)

	data, err := cache.Get[EncodeDecodeStruct](ctx, key)
	assert.Nil(t, err)

	assert.Equal(t, value, *data)
}

func TestGet_Err(t *testing.T) {
	ctx := setupFacade()

	key := gofakeit.Username()
	value, err := cache.Get[interface{}](ctx, key)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func setupFacade() context.Context {
	cfg := &cache.Configs{Dsn: "bigcache", SecondsToLive: 60}
	c, _ := cache.NewBigCache(context.Background(), cfg)
	return cache.WithContext(context.Background(), c)
}
