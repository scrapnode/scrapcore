package auth_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/utils"
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChainSign(t *testing.T) {
	creds := &auth.SignCreds{
		Username: gofakeit.UUID(),
		Password: gofakeit.Password(true, true, true, false, false, 64),
	}

	ctx := context.Background()
	ctx = xlogger.WithContext(ctx, xlogger.New(xlogger.LEVEL_TEST))

	cache, err := xcache.NewBigCache(ctx, &xcache.Configs{Dsn: "bigcache", SecondsToLive: 1800})
	assert.Nil(t, err)
	assert.Nil(t, cache.Connect(ctx))
	defer func() {
		_ = cache.Disconnect(ctx)
	}()
	ctx = xcache.WithContext(ctx, cache)

	authenticator, err := auth.NewChain(ctx, []*auth.Chain{
		{Auth: auth.NewAccessKey(creds.Username, creds.Password)},
	})
	assert.Nil(t, err)

	// sign
	tokens, err := authenticator.Sign(context.Background(), creds)
	assert.Nil(t, err)
	assert.NotNil(t, tokens)

	key, err := utils.MD5(creds)
	assert.Nil(t, err)
	bytes, err := cache.Get(ctx, key)
	assert.Nil(t, err)
	pairs, err := xcache.Decode[auth.TokenPair](bytes)
	assert.Nil(t, err)
	assert.Equal(t, tokens.AccessToken, (*auth.TokenPair)(pairs).AccessToken)
	assert.Equal(t, tokens.RefreshToken, (*auth.TokenPair)(pairs).RefreshToken)
}
