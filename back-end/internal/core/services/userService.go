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

func (s *UserService) SetUserPhoto(ctx context.Context, userId uint32, photo string) error {
	return s.repository.SetUserPhoto(ctx, userId, photo)
}

func (s *UserService) ChangeUserPassword(ctx context.Context, userId uint32, oldPassword, newPassword string) error {
	return s.repository.ChangeUserPassword(ctx, userId, oldPassword, newPassword)
}

func (s *UserService) GetUserPhoto(ctx context.Context, userId uint32) (string, error) {
	return s.repository.GetUserPhoto(ctx, userId)
}
