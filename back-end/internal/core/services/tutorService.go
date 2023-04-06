package services

import (
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
	"github.com/loukaspe/nursing-academiq/internal/repositories"
)

func NewTutorService(repository *repositories.TutorRepository) *TutorService {
	return &TutorService{repository: repository}
}

type TutorService struct {
	repository ports.TutorRepositoryInterface
}

func (service TutorService) GetTutor(uid uint32) (*domain.Tutor, error) {
	return service.repository.GetTutor(uid)
}

func (service TutorService) CreateTutor(tutor *domain.Tutor) error {
	return service.repository.CreateTutor(tutor)
}

func (service TutorService) UpdateTutor(uid uint32, tutor *domain.Tutor) error {
	return service.repository.UpdateTutor(uid, tutor)
}

func (service TutorService) DeleteTutor(uid uint32) error {
	return service.repository.DeleteTutor(uid)
}
