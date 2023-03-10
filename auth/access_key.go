package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/benbjohnson/clock"
	"github.com/golang-jwt/jwt/v4"
	"github.com/scrapnode/scrapcore/utils"
	"strings"
	"time"
)

type AccessKeyClaims struct {
	Workspaces []string `json:"workspaces"`
	Email      string   `json:"email"`
	jwt.RegisteredClaims
}

func NewAccessKey(id, secret string) *AccessKey {
	return &AccessKey{
		Clock:  clock.New(),
		id:     id,
		secret: secret,
		key:    []byte(fmt.Sprintf("%s:%s", id, secret)),
		algo:   jwt.SigningMethodHS512,
		issuer: "https://access_key.auth.scrapnode.com",
	}
}

type AccessKey struct {
	Clock  clock.Clock
	id     string
	secret string
	key    []byte
	algo   jwt.SigningMethod
	issuer string
}

func (auth *AccessKey) Connect(ctx context.Context) error {
	return nil
}

func (auth *AccessKey) Disconnect(ctx context.Context) error {
	return nil
}

func (auth *AccessKey) Sign(ctx context.Context, creds *SignCreds) (*TokenPair, error) {
	if ok := creds.Username == auth.id && creds.Password == auth.secret; !ok {
		return nil, errors.New("auth: incorrect username or password")
	}

	return auth.sign(creds.Username)
}

func (auth *AccessKey) sign(username string) (*TokenPair, error) {
	now := auth.Clock.Now().UTC()
	// access token
	acessToken := jwt.NewWithClaims(auth.algo, AccessKeyClaims{
		[]string{"*"},
		fmt.Sprintf("%s@scrapnode.com", username),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ACCESS_TOKEN_EXPIRE_HOURS) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    auth.issuer,
			Subject:   username,
			ID:        utils.NewId("jwt_access"),
		},
	})
	accessTokenStr, err := acessToken.SignedString(auth.key)
	if err != nil {
		return nil, err
	}

	// refresh token
	refreshToken := jwt.NewWithClaims(auth.algo, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(REFRESH_TOKEN_EXPIRE_HOURS) * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    auth.issuer,
		Subject:   username,
		ID:        utils.NewId("jwt_refresh"),
	})
	refreshTokenStr, err := refreshToken.SignedString(auth.key)
	if err != nil {
		return nil, err
	}

	return &TokenPair{AccessToken: accessTokenStr, RefreshToken: refreshTokenStr}, nil
}

func (auth *AccessKey) Verify(ctx context.Context, accessToken string) (*Account, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if ok := token.Method.Alg() == auth.algo.Alg(); !ok {
			return nil, errors.New("auth: unexpected signing method")
		}

		return auth.key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("auth: invalid token")
	}

	claims, err := auth.claims(token.Claims)
	if err != nil {
		return nil, err
	}
	if claims.Email == "" || len(claims.Workspaces) == 0 {
		return nil, errors.New("auth: invalid token claims")
	}
	if !strings.HasPrefix(claims.ID, "jwt_access") {
		return nil, errors.New("auth: invalid token type")
	}

	account := &Account{
		Workspaces: claims.Workspaces,
		Id:         claims.Subject,
		Name:       claims.Subject,
		Email:      claims.Email,
	}
	return account, nil
}

func (auth *AccessKey) claims(jwtclaims jwt.Claims) (*AccessKeyClaims, error) {
	mapclaims, ok := jwtclaims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("auth: invalid token claims")
	}
	claims := &AccessKeyClaims{
		[]string{},
		"",
		jwt.RegisteredClaims{
			Issuer:  "",
			ID:      "",
			Subject: "",
		},
	}
	ws, ok := mapclaims["workspaces"].([]interface{})
	if ok {
		for _, value := range ws {
			if workspace, ok := value.(string); ok {
				claims.Workspaces = append(claims.Workspaces, workspace)
			}
		}
	}

	if email, ok := mapclaims["email"].(string); ok {
		claims.Email = email
	}

	issuer, ok := mapclaims["iss"].(string)
	if !ok || issuer != auth.issuer {
		return nil, errors.New("auth: invalid token claims")
	}
	claims.Issuer = issuer

	id, ok := mapclaims["jti"].(string)
	if !ok {
		return nil, errors.New("auth: invalid token claims")
	}
	claims.ID = id

	subject, ok := mapclaims["sub"].(string)
	if !ok {
		return nil, errors.New("auth: invalid token claims")
	}
	claims.Subject = subject

	return claims, nil
}

func (auth *AccessKey) Refresh(ctx context.Context, tokens *TokenPair) (*TokenPair, error) {
	account, err := auth.Verify(ctx, tokens.AccessToken)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokens.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if ok := token.Method.Alg() == auth.algo.Alg(); !ok {
			return nil, errors.New("auth: unexpected signing method")
		}

		return auth.key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("auth: invalid token")
	}

	claims, err := auth.claims(token.Claims)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(claims.ID, "jwt_refresh") {
		return nil, errors.New("auth: invalid token type")
	}

	return auth.sign(account.Id)
}
