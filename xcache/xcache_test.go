package xcache_test

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/redis/go-redis/v9"
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew_Err(t *testing.T) {
	ctx := xlogger.WithContext(context.Background(), xlogger.New(xlogger.LEVEL_TEST))

	c, err := xcache.New(ctx, &xcache.Configs{
		Dsn:           "/\n/",
		SecondsToLive: 0,
	})
	assert.NotNil(t, err)
	assert.Nil(t, c)
}

func TestNew_ReturnMemoryCache(t *testing.T) {
	ctx := xlogger.WithContext(context.Background(), xlogger.New(xlogger.LEVEL_TEST))

	c, err := xcache.New(ctx, &xcache.Configs{
		Dsn: "bigcache://localhost",
	})
	assert.Nil(t, err)

	_, ok := c.Client().(*bigcache.BigCache)
	assert.True(t, ok)
}

func TestNew_ReturnRedis(t *testing.T) {
	ctx := xlogger.WithContext(context.Background(), xlogger.New(xlogger.LEVEL_TEST))

	c, err := xcache.New(ctx, &xcache.Configs{
		Dsn: "redis://localhost:6379/0",
	})
	assert.Nil(t, err)

	_, ok := c.Client().(*redis.Client)
	assert.True(t, ok)
}
