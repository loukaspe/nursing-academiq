package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
	"github.com/loukaspe/nursing-academiq/internal/repositories"
)

func NewStudentService(repository *repositories.StudentRepository) *StudentService {
	return &StudentService{repository: repository}
}

type StudentService struct {
	repository ports.StudentRepositoryInterface
}

func (service StudentService) GetStudent(ctx context.Context, uid uint32) (*domain.Student, error) {
	return service.repository.GetStudent(ctx, uid)
}

func (service StudentService) CreateStudent(ctx context.Context, student *domain.Student) (uint, error) {
	return service.repository.CreateStudent(ctx, student)
}

func (service StudentService) UpdateStudent(ctx context.Context, uid uint32, student *domain.Student) error {
	return service.repository.UpdateStudent(ctx, uid, student)
}

func (service StudentService) DeleteStudent(ctx context.Context, uid uint32) error {
	return service.repository.DeleteStudent(ctx, uid)
}

func (service StudentService) RegisterCourses(ctx context.Context, studentID uint32, courses []domain.Course) (*domain.Student, error) {
	return service.repository.RegisterCourses(ctx, studentID, courses)
}
