package auth

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/utils"
	"github.com/scrapnode/scrapcore/xcache"
)

type ChainCacheSet struct {
	cache xcache.Cache
	next  Auth
}

func (chain *ChainCacheSet) SetNext(auth Auth) {
	if auth != nil {
		chain.next = auth
	}
}

func (chain *ChainCacheSet) Connect(ctx context.Context) error {
	if chain.next == nil {
		return nil
	}
	return chain.next.Connect(ctx)
}

func (chain *ChainCacheSet) Disconnect(ctx context.Context) error {
	if chain.next != nil {
		return nil
	}
	return chain.next.Connect(ctx)
}

func (chain *ChainCacheSet) Sign(ctx context.Context, creds *SignCreds) (*TokenPair, error) {
	if chain.next == nil {
		return nil, errors.New("auth: incorrect username or password")
	}

	pairs, err := chain.next.Sign(ctx, creds)
	if err != nil {
		return nil, err
	}

	// @TODO: add some log here
	bytes, err := xcache.Encode(pairs)
	if err != nil {
		return pairs, nil
	}

	// @TODO: add some log here
	if key, err := utils.MD5(creds); err == nil {
		if err := chain.cache.Set(ctx, key, bytes); err == nil {
			return pairs, nil
		}
	}

	return pairs, nil
}

func (chain *ChainCacheSet) Verify(ctx context.Context, token string) (*Account, error) {
	if chain.next == nil {
		return nil, errors.New("auth: invalid token")
	}

	account, err := chain.next.Verify(ctx, token)
	if err != nil {
		return nil, err
	}

	// @TODO: add some log here
	bytes, err := xcache.Encode(account)
	if err != nil {
		return account, nil
	}

	// @TODO: add some log here
	if key, err := utils.MD5(token); err == nil {
		if err := chain.cache.Set(ctx, key, bytes); err == nil {
			return account, nil
		}
	}

	return account, nil
}

func (chain *ChainCacheSet) Refresh(ctx context.Context, tokens *TokenPair) (*TokenPair, error) {
	if chain.next == nil {
		return nil, errors.New("auth: invalid token")
	}

	pairs, err := chain.next.Refresh(ctx, tokens)
	if err != nil {
		return nil, err
	}

	// @TODO: add some log here
	bytes, err := xcache.Encode(pairs)
	if err != nil {
		return pairs, nil
	}

	// @TODO: add some log here
	if key, err := utils.MD5(tokens); err == nil {
		if err := chain.cache.Set(ctx, key, bytes); err == nil {
			return pairs, nil
		}
	}

	return pairs, nil
}
