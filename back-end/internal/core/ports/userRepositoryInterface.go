package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type UserRepositoryInterface interface {
	Login(ctx context.Context, username, password string) (*domain.User, uint, error)
}
