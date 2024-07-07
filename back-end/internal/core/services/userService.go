package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
)

type UserService struct {
	repository ports.UserRepositoryInterface
}

func NewUserService(repository ports.UserRepositoryInterface) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) ChangeUserPassword(ctx context.Context, userId uint32, oldPassword, newPassword string) error {
	return s.repository.ChangeUserPassword(ctx, userId, oldPassword, newPassword)
}
