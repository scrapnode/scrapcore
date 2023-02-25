package auth_test

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/mocks"
	"github.com/scrapnode/scrapcore/utils"
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestChainSign(t *testing.T) {
	creds := &auth.SignCreds{
		Username: gofakeit.UUID(),
		Password: gofakeit.Password(true, true, true, false, false, 64),
	}

	ctx := context.Background()
	ctx = xlogger.WithContext(ctx, xlogger.New(xlogger.LEVEL_TEST))

	cache := &mocks.Cache{}
	ctx = xcache.WithContext(ctx, cache)

	authenticator, err := auth.NewChain(ctx, []*auth.Chain{
		{Auth: auth.NewAccessKey(creds.Username, creds.Password)},
	})
	assert.Nil(t, err)

	cacheKey, _ := utils.MD5(creds)
	cache.On("Get", context.Background(), cacheKey).Return(nil, errors.New("entry was not found")).Once()
	cache.On("Set", context.Background(), cacheKey, mock.AnythingOfType("[]uint8")).Return(nil).Once()

	tokens, err := authenticator.Sign(context.Background(), creds)
	assert.Nil(t, err)
	assert.NotNil(t, tokens)
}
