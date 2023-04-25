package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type TutorRepositoryInterface interface {
	GetTutor(ctx context.Context, uid uint32) (*domain.Tutor, error)
	CreateTutor(context.Context, *domain.Tutor) (uint, error)
	UpdateTutor(context.Context, uint32, *domain.Tutor) error
	DeleteTutor(ctx context.Context, uid uint32) error
}
