package auth

import "context"

var EXPIRE_HOURS = 1

type Auth interface {
	Sign(ctx context.Context, creds *SignCreds) (*Tokens, error)
	Verify(ctx context.Context, token string) (*Account, error)
	Refresh(ctx context.Context, tokens *Tokens) (*Tokens, error)
}

type SignCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Account struct {
	Workspaces  []string `json:"workspaces"`
	WorkspaceId string   `json:"workspace_id"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
}
