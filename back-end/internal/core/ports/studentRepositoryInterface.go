package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type StudentRepositoryInterface interface {
	GetStudent(ctx context.Context, uid uint32) (*domain.Student, error)
	CreateStudent(context.Context, *domain.Student) (uint, error)
	UpdateStudent(context.Context, uint32, *domain.Student) error
	DeleteStudent(ctx context.Context, uid uint32) error
}
