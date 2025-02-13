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

func (j *JwtService) CreateAccessJwtToken(user *domain.JwtSubject) (string, error) {
	tokenValue, err := j.jwtDomain.CreateAccessToken(user.User.Username, user)
	if err != nil {
		return "", err
	}
	return tokenValue, nil
}

func (j *JwtService) CreateRefreshJwtToken(user *domain.JwtSubject) (string, error) {
	tokenValue, err := j.jwtDomain.CreateRefreshToken(user.User.Username, nil)
	if err != nil {
		return "", err
	}
	return tokenValue, nil
}

func (j *JwtService) AccessClaimsFromJwtToken(token string) (jwt.MapClaims, error) {
	claims, err := j.jwtDomain.GetClaimsFromAccessToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func (j *JwtService) RefreshClaimsFromJwtToken(token string) (jwt.MapClaims, error) {
	claims, err := j.jwtDomain.GetClaimsFromRefreshToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
