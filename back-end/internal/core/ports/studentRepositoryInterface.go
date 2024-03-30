package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type StudentRepositoryInterface interface {
	GetStudent(ctx context.Context, uid uint32) (*domain.Student, error)
	// GetExtendedStudent The difference here is that this function returns a Student with more fields,
	// that require more DB queries
	GetExtendedStudent(ctx context.Context, uid uint32) (*domain.Student, error)
	CreateStudent(context.Context, *domain.Student) (uint, error)
	UpdateStudent(context.Context, uint32, *domain.Student) error
	DeleteStudent(ctx context.Context, uid uint32) error
	RegisterCourses(ctx context.Context, studentID uint32, courses []domain.Course) (*domain.Student, error)
}
