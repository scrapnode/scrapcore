package xcache_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigCacheFunctions(t *testing.T) {
	ctx := xlogger.WithContext(context.Background(), xlogger.New(xlogger.LEVEL_TEST))
	cfg := &xcache.Configs{Dsn: "bigcache://localhost", SecondsToLive: 60}

	cache, err := xcache.NewBigCache(ctx, cfg)
	assert.Nil(t, err)
	assert.Nil(t, cache.Connect(ctx))
	defer func() {
		_ = cache.Disconnect(ctx)
	}()

	key := gofakeit.Username()
	value := gofakeit.Bool()

	t.Run("Set", func(t *testing.T) {
		bytes, err := xcache.Encode(value)
		assert.Nil(t, err)
		assert.Nil(t, cache.Set(ctx, key, bytes))
	})
	t.Run("Get", func(t *testing.T) {
		bytes, err := cache.Get(ctx, key)
		assert.Nil(t, err)
		dataptr, err := xcache.Decode[bool](bytes)
		assert.Nil(t, err)
		assert.Equal(t, value, *dataptr)
	})
	t.Run("Exists/True", func(t *testing.T) {
		assert.True(t, cache.Exists(ctx, key))
	})
	t.Run("Del", func(t *testing.T) {
		assert.Nil(t, cache.Del(ctx, key))
	})
	t.Run("Exists/False", func(t *testing.T) {
		assert.False(t, cache.Exists(ctx, key))
	})

	t.Run("Decr", func(t *testing.T) {
		intkey := gofakeit.Username()
		intvalue := int64(0)

		t.Run("NotFound(0->-1)", func(t *testing.T) {
			count, err := cache.Decr(ctx, intkey)
			assert.Nil(t, err)
			assert.Equal(t, intvalue-1, count)
			intvalue = count
		})

		t.Run("Decr/2-1", func(t *testing.T) {
			count, err := cache.Decr(ctx, intkey)
			assert.Nil(t, err)
			assert.Equal(t, intvalue-1, count)
			intvalue = count
		})
	})

	t.Run("Incr", func(t *testing.T) {
		intkey := gofakeit.Username()
		intvalue := int64(0)

		t.Run("NotFound(0->1)", func(t *testing.T) {
			count, err := cache.Incr(ctx, intkey)
			assert.Nil(t, err)
			assert.Equal(t, intvalue+1, count)
			intvalue = count
		})
		t.Run("1->2", func(t *testing.T) {
			count, err := cache.Incr(ctx, intkey)
			assert.Nil(t, err)
			assert.Equal(t, intvalue+1, count)
			intvalue = count
		})
	})
}
