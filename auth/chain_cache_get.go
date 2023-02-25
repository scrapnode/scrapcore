package auth

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/utils"
	"github.com/scrapnode/scrapcore/xcache"
)

type ChainCacheGet struct {
	cache xcache.Cache
	next  Auth
}

func (chain *ChainCacheGet) SetNext(auth Auth) {
	if auth != nil {
		chain.next = auth
	}
}

func (chain *ChainCacheGet) Connect(ctx context.Context) error {
	if chain.next == nil {
		return nil
	}
	return chain.next.Connect(ctx)
}

func (chain *ChainCacheGet) Disconnect(ctx context.Context) error {
	if chain.next != nil {
		return nil
	}
	return chain.next.Connect(ctx)
}

func (chain *ChainCacheGet) Sign(ctx context.Context, creds *SignCreds) (*TokenPair, error) {
	if key, err := utils.MD5(creds); err == nil {
		bytes, err := chain.cache.Get(ctx, key)
		if err == nil {
			if pairs, err := xcache.Decode[TokenPair](bytes); err == nil {
				return pairs, nil
			}
		}
	}

	if chain.next == nil {
		return nil, errors.New("auth: incorrect username or password")
	}
	return chain.next.Sign(ctx, creds)
}

func (chain *ChainCacheGet) Verify(ctx context.Context, token string) (*Account, error) {
	if key, err := utils.MD5(token); err == nil {
		if bytes, err := chain.cache.Get(ctx, key); err == nil {
			if account, err := xcache.Decode[Account](bytes); err == nil {
				return account, nil
			}
		}
	}

	if chain.next == nil {
		return nil, errors.New("auth: invalid token")
	}
	return chain.next.Verify(ctx, token)
}

func (chain *ChainCacheGet) Refresh(ctx context.Context, tokens *TokenPair) (*TokenPair, error) {
	if key, err := utils.MD5(tokens); err == nil {
		if bytes, err := chain.cache.Get(ctx, key); err == nil {
			if pairs, err := xcache.Decode[TokenPair](bytes); err == nil {
				return pairs, nil
			}
		}
	}

	if chain.next == nil {
		return nil, errors.New("auth: invalid token")
	}
	return chain.next.Refresh(ctx, tokens)
}
