package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type QuizRepositoryInterface interface {
	GetQuiz(ctx context.Context, uid uint32) (*domain.Quiz, error)
	GetQuizByTutorID(ctx context.Context, tutorID uint32) ([]domain.Quiz, error)
	GetQuizByCourseID(ctx context.Context, courseID uint32) ([]domain.Quiz, error)
	//CreateQuiz(context.Context, *domain.Quiz) (uint, error)
	//UpdateQuiz(context.Context, uint32, *domain.Quiz) error
	//DeleteQuiz(ctx context.Context, uid uint32) error
}
