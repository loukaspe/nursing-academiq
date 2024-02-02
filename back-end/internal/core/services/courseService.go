package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
)

func NewCourseService(repository ports.CourseRepositoryInterface) *CourseService {
	return &CourseService{repository: repository}
}

type CourseService struct {
	repository ports.CourseRepositoryInterface
}

func (service CourseService) GetCourse(ctx context.Context, uid uint32) (*domain.Course, error) {
	return service.repository.GetCourse(ctx, uid)
}

func (service CourseService) GetCourseByStudentID(ctx context.Context, studentID uint32) ([]domain.Course, error) {
	return service.repository.GetCourseByStudentID(ctx, studentID)
}

//func (service CourseService) CreateCourse(ctx context.Context, course *domain.Course) (uint, error) {
//	return service.repository.CreateCourse(ctx, course)
//}
//
//func (service CourseService) UpdateCourse(ctx context.Context, uid uint32, course *domain.Course) error {
//	return service.repository.UpdateCourse(ctx, uid, course)
//}
//
//func (service CourseService) DeleteCourse(ctx context.Context, uid uint32) error {
//	return service.repository.DeleteCourse(ctx, uid)
//}
