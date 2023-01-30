package auth_test

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessKey(t *testing.T) {
	creds := &auth.SignCreds{
		Username: gofakeit.UUID(),
		Password: gofakeit.Password(true, true, true, false, false, 64),
	}
	authenticator := auth.NewAccessKey(creds.Username, creds.Password)

	tokens, err := authenticator.Sign(context.Background(), creds)
	assert.Nil(t, err)
	assert.NotNil(t, tokens)

	newtokens, err := authenticator.Refresh(context.Background(), tokens)
	assert.Nil(t, err)
	assert.NotNil(t, newtokens)
	assert.NotEqual(t, tokens.AccessToken, newtokens.AccessToken)
	assert.NotEqual(t, tokens.RefreshToken, newtokens.RefreshToken)

	account, err := authenticator.Verify(context.Background(), tokens.AccessToken)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, creds.Username, account.Id)
	assert.Equal(t, "*", account.WorkspaceId)
	assert.Equal(t, []string{"*"}, account.Workspaces)
	assert.Equal(t, fmt.Sprintf("%s@scrapnode.com", creds.Username), account.Email)
}
