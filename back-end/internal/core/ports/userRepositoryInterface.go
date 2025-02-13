package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type UserRepositoryInterface interface {
	Login(ctx context.Context, username, password string) (*domain.User, uint, error)
	ChangeUserPassword(ctx context.Context, userID uint32, oldPassword, newPassword string) error
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}
