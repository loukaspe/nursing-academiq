package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type CourseRepositoryInterface interface {
	GetCourse(ctx context.Context, uid uint32) (*domain.Course, error)
	GetExtendedCourse(ctx context.Context, uid uint32) (*domain.Course, error)
	GetCourseChapters(ctx context.Context, uid uint32) (*domain.Course, error)
	GetCourses(ctx context.Context) ([]domain.Course, error)
	GetMostRecentCourses(ctx context.Context, limit int) ([]domain.Course, error)
	GetCourseByTutorID(ctx context.Context, tutorID uint32) ([]domain.Course, error)
	CreateCourse(context.Context, *domain.Course, uint) (uint, error)
	UpdateCourse(context.Context, uint32, *domain.Course) error
	DeleteCourse(ctx context.Context, uid uint32) error
}
