package ports

import "github.com/loukaspe/nursing-academiq/internal/core/domain"

type StudentRepositoryInterface interface {
	GetStudent(uid uint32) (*domain.Student, error)
	CreateStudent(*domain.Student) error
	UpdateStudent(uint32, *domain.Student) error
	DeleteStudent(uid uint32) error
}
