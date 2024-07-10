package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type QuestionRepositoryInterface interface {
	ImportForCourse(ctx context.Context, questions []domain.Question, courseID uint) error
}
