package ports

import "github.com/loukaspe/nursing-academiq/internal/core/domain"

type UserRepositoryInterface interface {
	GetUser(uid uint32) (*domain.User, error)
	CreateUser(*domain.User) error
	UpdateUser(uint32, *domain.User) error
	DeleteUser(uid uint32) error
}
