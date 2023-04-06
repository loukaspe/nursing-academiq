package ports

import "github.com/loukaspe/nursing-academiq/internal/core/domain"

type TutorRepositoryInterface interface {
	GetTutor(uid uint32) (*domain.Tutor, error)
	CreateTutor(*domain.Tutor) error
	UpdateTutor(uint32, *domain.Tutor) error
	DeleteTutor(uid uint32) error
}
