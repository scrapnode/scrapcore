package auth

import (
	"context"
	"github.com/samber/lo"
)

var (
	ACCESS_TOKEN_EXPIRE_HOURS  = 1
	REFRESH_TOKEN_EXPIRE_HOURS = 720 // 30 days
)

type Auth interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Sign(ctx context.Context, creds *SignCreds) (*TokenPair, error)
	Verify(ctx context.Context, token string) (*Account, error)
	Refresh(ctx context.Context, tokens *TokenPair) (*TokenPair, error)
}

type SignCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Account struct {
	Workspaces []string `json:"workspaces"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
}

func (account *Account) OwnWorkspace(id string) bool {
	// can access all workspaces
	if lo.Contains(account.Workspaces, "*") {
		return true
	}
	return account.Workspaces != nil && lo.Contains(account.Workspaces, id)
}
