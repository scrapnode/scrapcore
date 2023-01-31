package auth

import (
	"context"
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

func NewAccessKey(id, key string) *AccessKey {
	return &AccessKey{
		Clock:  clock.New(),
		id:     id,
		key:    key,
		secret: []byte(fmt.Sprintf("%s:%s", id, key)),
		algo:   jwt.SigningMethodHS512,
		issuer: "https://auth.scrapnode.com",
	}
}

type AccessKey struct {
	Clock  clock.Clock
	id     string
	key    string
	secret []byte
	algo   jwt.SigningMethod
	issuer string
}

func (auth *AccessKey) Connect(ctx context.Context) error {
	return nil
}

func (auth *AccessKey) Disconnect(ctx context.Context) error {
	return nil
}

func (auth *AccessKey) Sign(ctx context.Context, creds *SignCreds) (*Tokens, error) {
	if ok := creds.Username == auth.id && creds.Password == auth.key; !ok {
		return nil, ErrSignFailed
	}

	return auth.sign(creds.Username)
}

func (auth *AccessKey) sign(username string) (*Tokens, error) {
	now := auth.Clock.Now().UTC()
	// access token
	attoken := jwt.NewWithClaims(auth.algo, AccessKeyClaims{
		[]string{"*"},
		fmt.Sprintf("%s@scrapnode.com", username),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(EXPIRE_HOURS) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    auth.issuer,
			Subject:   username,
			ID:        utils.NewId("jwt_access"),
		},
	})
	at, err := attoken.SignedString(auth.secret)
	if err != nil {
		return nil, err
	}

	// refresh token
	rttoken := jwt.NewWithClaims(auth.algo, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(EXPIRE_HOURS*2) * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		Issuer:    auth.issuer,
		Subject:   username,
		ID:        utils.NewId("jwt_refresh"),
	})
	rt, err := rttoken.SignedString(auth.secret)
	if err != nil {
		return nil, err
	}

	return &Tokens{AccessToken: at, RefreshToken: rt}, err
}

func (auth *AccessKey) Verify(ctx context.Context, accessToken string) (*Account, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if ok := token.Method.Alg() == auth.algo.Alg(); !ok {
			return nil, ErrInvalidSignMethod
		}

		return auth.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, err := auth.claims(token.Claims)
	if err != nil {
		return nil, err
	}
	if claims.Email == "" || len(claims.Workspaces) == 0 {
		return nil, ErrInvalidTokenClaims
	}
	if !strings.HasPrefix(claims.ID, "jwt_access") {
		return nil, ErrInvalidTokenType
	}

	account := &Account{
		Workspaces: claims.Workspaces,
		Id:         claims.Subject,
		Name:       claims.Subject,
		Email:      claims.Email,
	}
	return account, err
}

func (auth *AccessKey) claims(jwtclaims jwt.Claims) (*AccessKeyClaims, error) {
	mapclaims, ok := jwtclaims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidTokenClaims
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
		return nil, ErrInvalidTokenClaims
	}
	claims.Issuer = issuer

	id, ok := mapclaims["jti"].(string)
	if !ok {
		return nil, ErrInvalidTokenClaims
	}
	claims.ID = id

	subject, ok := mapclaims["sub"].(string)
	if !ok {
		return nil, ErrInvalidTokenClaims
	}
	claims.Subject = subject

	return claims, nil
}

func (auth *AccessKey) Refresh(ctx context.Context, tokens *Tokens) (*Tokens, error) {
	account, err := auth.Verify(ctx, tokens.AccessToken)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokens.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if ok := token.Method.Alg() == auth.algo.Alg(); !ok {
			return nil, ErrInvalidSignMethod
		}

		return auth.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, err := auth.claims(token.Claims)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(claims.ID, "jwt_refresh") {
		return nil, ErrInvalidTokenType
	}

	return auth.sign(account.Id)
}
