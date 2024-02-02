package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type JwtService struct {
	jwtDomain domain.JwtClaimsInterface
}

func NewJwtService(domain domain.JwtClaimsInterface) *JwtService {
	return &JwtService{jwtDomain: domain}
}

func (j *JwtService) CreateJwtTokenService(user *domain.JwtSubject) (string, error) {
	tokenValue, err := j.jwtDomain.CreateToken(user.User.Username, user)
	if err != nil {
		return "", err
	}
	return tokenValue, nil
}

func (j *JwtService) ClaimsFromJwtTokenService(token string) (jwt.MapClaims, error) {
	claims, err := j.jwtDomain.GetClaimsFromToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
