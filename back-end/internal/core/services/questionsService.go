package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
)

func NewQuestionService(repository ports.QuestionRepositoryInterface) *QuestionService {
	return &QuestionService{repository: repository}
}

type QuestionService struct {
	repository ports.QuestionRepositoryInterface
}

func (service QuestionService) ImportForCourse(ctx context.Context, questions []domain.Question, courseID uint) error {
	return service.repository.ImportForCourse(ctx, questions, courseID)
}
