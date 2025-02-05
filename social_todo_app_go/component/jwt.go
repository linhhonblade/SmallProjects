package component

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	sctx "github.com/linhhonblade/service-context"
	"time"
)

const (
	defaultSecret                = "very-important-please-change-it"
	defaultExpireTokenInSecond   = 60 * 60 * 24 * 7  // 7 days
	defaultExpireRefreshInSecond = 60 * 60 * 24 * 14 // 14 days
)

var (
	ErrSecretKeyNotValid     = errors.New("secret key must be in 32 bytes")
	ErrTokenLifeTimeTooShort = errors.New("token life time too short")
)

type jwtx struct {
	id                     string
	secret                 string
	expireTokenInSeconds   int
	expireRefreshInSeconds int
}

func NewJWTProvider(secret string, expireTokenInSeconds, expireRefreshInSeconds int) *jwtx {
	return &jwtx{
		secret:                 secret,
		expireTokenInSeconds:   expireTokenInSeconds,
		expireRefreshInSeconds: expireRefreshInSeconds,
	}
}

func NewJWT(id string) *jwtx {
	return &jwtx{id: id}
}

// Implement Component interface in service-context package
func (j *jwtx) ID() string {
	return j.id
}
func (j *jwtx) InitFlags() {
	flag.StringVar(
		&j.secret,
		"jwt-secret",
		defaultSecret, "Secret key to sign JWT")
	flag.IntVar(&j.expireTokenInSeconds, "jwt-exp-sec", defaultExpireTokenInSecond, "Number of seconds access token will expire")
	flag.IntVar(&j.expireRefreshInSeconds, "jwt-exp-refresh-sec", defaultExpireRefreshInSecond, "Number of seconds refresh token will expire")
}
func (j *jwtx) Activate(_ sctx.ServiceContext) error {
	if len(j.secret) != 32 {
		return fmt.Errorf("%w: %v", ErrSecretKeyNotValid, len(j.secret))
	}
	if j.expireTokenInSeconds < 60*60*24 {
		return fmt.Errorf("%w: %v", ErrTokenLifeTimeTooShort, j.expireTokenInSeconds)
	}
	return nil
}
func (j *jwtx) Stop() error {
	return nil
}

func (j *jwtx) IssueToken(ctx context.Context, id, sub string) (token string, err error) {
	now := time.Now().UTC()

	claims := jwt.RegisteredClaims{
		Subject:   sub,
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(j.expireTokenInSeconds))),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        id}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSignedStr, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenSignedStr, nil
}

func (j *jwtx) TokenExpiredInSeconds() int {
	return j.expireTokenInSeconds
}

func (j *jwtx) TokenRefreshInSeconds() int {
	return j.expireRefreshInSeconds
}

func (j *jwtx) ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error) {
	var rc jwt.RegisteredClaims
	token, err := jwt.ParseWithClaims(tokenString, &rc, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	return &rc, nil
}

type TokenProvider interface {
	IssueToken(ctx context.Context, id, sub string) (token string, err error)
	TokenExpiredInSeconds() int
	TokenRefreshInSeconds() int
	ParseToken(ctx context.Context, tokenString string) (claims *jwt.RegisteredClaims, err error)
}
