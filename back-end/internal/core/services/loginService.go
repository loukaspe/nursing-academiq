package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
)

type LoginService struct {
	repository ports.UserRepositoryInterface
}

func NewLoginService(repository ports.UserRepositoryInterface) *LoginService {
	return &LoginService{repository: repository}
}

func (j *LoginService) Login(ctx context.Context, username, password string) (*domain.User, uint, error) {
	return j.repository.Login(ctx, username, password)
}
