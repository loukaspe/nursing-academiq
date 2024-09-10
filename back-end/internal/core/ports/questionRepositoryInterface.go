package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type QuestionRepositoryInterface interface {
	ImportForCourse(ctx context.Context, questions []domain.Question, courseID uint) error
	GetQuestions(ctx context.Context) ([]domain.Question, error)
	GetQuestionsByCourseID(ctx context.Context, courseID uint32) ([]domain.Question, error)
	CreateQuestion(context.Context, *domain.Question) (uint, error)
	GetQuestion(ctx context.Context, uid uint32) (*domain.Question, error)
	UpdateQuestion(context.Context, uint32, *domain.Question) error
	DeleteQuestion(ctx context.Context, uid uint32) error
}
