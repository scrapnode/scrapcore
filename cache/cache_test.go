package cache_test

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/redis/go-redis/v9"
	"github.com/scrapnode/scrapcore/cache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew_Err(t *testing.T) {
	c, err := cache.New(context.Background(), &cache.Configs{
		Dsn:           "/\n/",
		SecondsToLive: 0,
	})
	assert.NotNil(t, err)
	assert.Nil(t, c)
}

func TestNew_ReturnMemoryCache(t *testing.T) {
	c, err := cache.New(context.Background(), &cache.Configs{})
	assert.Nil(t, err)

	_, ok := c.Client().(*bigcache.BigCache)
	assert.True(t, ok)
}

func TestNew_ReturnRedis(t *testing.T) {
	c, err := cache.New(context.Background(), &cache.Configs{
		Dsn: "redis://localhost:6379/0",
	})
	assert.Nil(t, err)

	_, ok := c.Client().(*redis.Client)
	assert.True(t, ok)
}
