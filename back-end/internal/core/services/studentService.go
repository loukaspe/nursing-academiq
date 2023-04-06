package services

import (
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

func (service StudentService) GetStudent(uid uint32) (*domain.Student, error) {
	return service.repository.GetStudent(uid)
}

func (service StudentService) CreateStudent(student *domain.Student) error {
	return service.repository.CreateStudent(student)
}

func (service StudentService) UpdateStudent(uid uint32, student *domain.Student) error {
	return service.repository.UpdateStudent(uid, student)
}

func (service StudentService) DeleteStudent(uid uint32) error {
	return service.repository.DeleteStudent(uid)
}
