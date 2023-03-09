package services

import (
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
	"github.com/loukaspe/nursing-academiq/internal/repositories"
)

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

type UserService struct {
	repository ports.UserRepositoryInterface
}

func (service UserService) GetUser(uid uint32) (*domain.User, error) {
	return service.repository.GetUser(uid)
}

func (service UserService) CreateUser(user *domain.User) error {
	return service.repository.CreateUser(user)
}

func (service UserService) UpdateUser(uid uint32, user *domain.User) error {
	return service.repository.UpdateUser(uid, user)
}

func (service UserService) DeleteUser(uid uint32) error {
	return service.repository.DeleteUser(uid)
}
