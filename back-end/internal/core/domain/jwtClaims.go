package domain

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

type JwtSubject struct {
	User   *User
	UserID uint
}

type JwtClaimsInterface interface {
	CreateAccessToken(sub string, userInfo interface{}) (string, error)
	CreateRefreshToken(sub string, userInfo interface{}) (string, error)
	GetClaimsFromAccessToken(tokenString string) (jwt.MapClaims, error)
	GetClaimsFromRefreshToken(tokenString string) (jwt.MapClaims, error)
	SetAccessJWTClaimsContext(ctx context.Context, claims jwt.MapClaims) context.Context
}
