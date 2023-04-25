package services

import (
	"context"
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

func (service TutorService) GetTutor(ctx context.Context, uid uint32) (*domain.Tutor, error) {
	return service.repository.GetTutor(ctx, uid)
}

func (service TutorService) CreateTutor(ctx context.Context, tutor *domain.Tutor) (uint, error) {
	return service.repository.CreateTutor(ctx, tutor)
}

func (service TutorService) UpdateTutor(ctx context.Context, uid uint32, tutor *domain.Tutor) error {
	return service.repository.UpdateTutor(ctx, uid, tutor)
}

func (service TutorService) DeleteTutor(ctx context.Context, uid uint32) error {
	return service.repository.DeleteTutor(ctx, uid)
}
