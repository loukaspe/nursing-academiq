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

func (service QuestionService) GetQuestions(ctx context.Context) ([]domain.Question, error) {
	return service.repository.GetQuestions(ctx)
}

func (service QuestionService) GetQuestionsByCourseID(ctx context.Context, courseID uint32) ([]domain.Question, error) {
	return service.repository.GetQuestionsByCourseID(ctx, courseID)
}

func (service QuestionService) CreateQuestion(ctx context.Context, question *domain.Question) (uint, error) {
	return service.repository.CreateQuestion(ctx, question)
}

func (service QuestionService) GetQuestion(ctx context.Context, uid uint32) (*domain.Question, error) {
	return service.repository.GetQuestion(ctx, uid)
}

func (service QuestionService) UpdateQuestion(ctx context.Context, uid uint32, question *domain.Question) error {
	return service.repository.UpdateQuestion(ctx, uid, question)
}

func (service QuestionService) DeleteQuestion(ctx context.Context, uid uint32) error {
	return service.repository.DeleteQuestion(ctx, uid)
}
