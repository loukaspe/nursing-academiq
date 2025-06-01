package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"time"
)

type claimskey int

var claimsKey claimskey

type AuthMechanism struct {
	accessSecretKey  []byte
	refreshSecretKey []byte
	signingMethod    string
}

func NewAuthMechanism(
	accessSecret string,
	refreshSecret string,
	signingMethod string,
) *AuthMechanism {
	return &AuthMechanism{
		accessSecretKey:  []byte(accessSecret),
		refreshSecretKey: []byte(refreshSecret),
		signingMethod:    signingMethod,
	}
}
func (j *AuthMechanism) CreateAccessToken(sub string, userInfo interface{}) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(j.signingMethod))
	expiration := time.Now().Add(time.Hour)
	token.Claims = &domain.JwtClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			Subject:   sub,
		}, userInfo,
	}
	val, err := token.SignedString(j.accessSecretKey)
	if err != nil {

		return "", err
	}
	return val, nil
}

func (j *AuthMechanism) CreateRefreshToken(sub string, userInfo interface{}) (string, error) {
	const weekTime = time.Hour * 24 * 7

	token := jwt.New(jwt.GetSigningMethod(j.signingMethod))
	expiration := time.Now().Add(weekTime)
	token.Claims = &domain.JwtClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			Subject:   sub,
		}, userInfo,
	}
	val, err := token.SignedString(j.refreshSecretKey)
	if err != nil {

		return "", err
	}
	return val, nil
}

func (j *AuthMechanism) GetClaimsFromAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.accessSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func (j *AuthMechanism) GetClaimsFromRefreshToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.refreshSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func (j *AuthMechanism) SetAccessJWTClaimsContext(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}
