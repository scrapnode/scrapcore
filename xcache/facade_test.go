package xcache_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSet_Ok(t *testing.T) {
	ctx, cleanup := setupFacade()
	defer cleanup()

	key := gofakeit.Username()
	value := EncodeDecodeStruct{
		Username: gofakeit.Username(),
		Age:      gofakeit.IntRange(10, 30),
		Active:   gofakeit.Bool(),
	}
	err := xcache.Set(ctx, key, value)
	assert.Nil(t, err)

	data, err := xcache.Get[EncodeDecodeStruct](ctx, key)
	assert.Nil(t, err)

	assert.Equal(t, value, *data)
}

func TestGet_Err(t *testing.T) {
	ctx, cleanup := setupFacade()
	defer cleanup()

	key := gofakeit.Username()
	value, err := xcache.Get[interface{}](ctx, key)
	assert.NotNil(t, err)
	assert.Nil(t, value)
}

func setupFacade() (context.Context, func()) {
	ctx := xlogger.WithContext(context.Background(), xlogger.New(xlogger.LEVEL_TEST))

	cfg := &xcache.Configs{Dsn: "bigcache://localhost", SecondsToLive: 60}
	cache, err := xcache.NewBigCache(ctx, cfg)
	if err != nil {
		panic(err)
	}
	if err := cache.Connect(ctx); err != nil {
		panic(err)
	}
	return xcache.WithContext(ctx, cache), func() { _ = cache.Disconnect(ctx) }
}
