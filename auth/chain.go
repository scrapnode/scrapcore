package auth

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/xcache"
)

// Chain needs to be constructed with order
// Get -> Set -> Main

func NewChain(ctx context.Context, steps []*Chain) (Auth, error) {
	if len(steps) == 0 {
		return nil, errors.New("could not init chains without steps")
	}

	cache := xcache.FromContext(ctx)

	var cursor Auth
	cursor = steps[len(steps)-1]
	cursor = &ChainCacheSet{cache: cache, next: cursor}
	cursor = &ChainCacheGet{cache: cache, next: cursor}

	if len(steps) > 1 {
		for i := len(steps) - 2; i >= 0; i-- {
			cursor = steps[i]
			cursor = &ChainCacheSet{cache: cache, next: cursor}
			cursor = &ChainCacheGet{cache: cache, next: cursor}
		}
	}

	return cursor, nil
}

type Chain struct {
	Auth Auth

	next Auth
}

func (chain *Chain) Connect(ctx context.Context) error {
	if err := chain.Auth.Connect(ctx); err != nil {
		return err
	}

	if chain.next == nil {
		return nil
	}
	return chain.next.Connect(ctx)
}

func (chain *Chain) Disconnect(ctx context.Context) error {
	if err := chain.Auth.Disconnect(ctx); err != nil {
		return err
	}

	if chain.next != nil {
		return nil
	}
	return chain.next.Connect(ctx)
}

func (chain *Chain) Sign(ctx context.Context, creds *SignCreds) (*TokenPair, error) {
	pairs, err := chain.Auth.Sign(ctx, creds)
	if err == nil {
		return pairs, nil
	}

	if chain.next == nil {
		return nil, err
	}
	return chain.next.Sign(ctx, creds)
}

func (chain *Chain) Verify(ctx context.Context, token string) (*Account, error) {
	account, err := chain.Auth.Verify(ctx, token)
	if err == nil {
		return account, nil
	}

	if chain.next == nil {
		return nil, err
	}
	return chain.next.Verify(ctx, token)
}

func (chain *Chain) Refresh(ctx context.Context, tokens *TokenPair) (*TokenPair, error) {
	pairs, err := chain.Auth.Refresh(ctx, tokens)
	if err == nil {
		return pairs, nil
	}

	if chain.next == nil {
		return nil, err
	}
	return chain.next.Refresh(ctx, tokens)
}
