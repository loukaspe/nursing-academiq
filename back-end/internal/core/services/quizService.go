package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
)

func NewQuizService(repository ports.QuizRepositoryInterface) *QuizService {
	return &QuizService{repository: repository}
}

type QuizService struct {
	repository ports.QuizRepositoryInterface
}

func (service QuizService) GetQuiz(ctx context.Context, uid uint32) (*domain.Quiz, error) {
	return service.repository.GetQuiz(ctx, uid)
}

func (service QuizService) SearchQuiz(ctx context.Context, title string, courseName string) ([]domain.Quiz, error) {
	return service.repository.SearchQuiz(ctx, title, courseName)
}

func (service QuizService) GetQuizByTutorID(ctx context.Context, tutorID uint32) ([]domain.Quiz, error) {
	return service.repository.GetQuizByTutorID(ctx, tutorID)
}

func (service QuizService) GetQuizByCourseID(ctx context.Context, courseID uint32) ([]domain.Quiz, error) {
	return service.repository.GetQuizByCourseID(ctx, courseID)
}

func (service QuizService) GetQuizzes(ctx context.Context) ([]domain.Quiz, error) {
	return service.repository.GetQuizzes(ctx)
}

func (service QuizService) GetMostRecentQuizzes(ctx context.Context, limit int) ([]domain.Quiz, error) {
	return service.repository.GetMostRecentQuizzes(ctx, limit)
}

func (service QuizService) CreateQuiz(ctx context.Context, quiz *domain.Quiz) (uint, error) {
	return service.repository.CreateQuiz(ctx, quiz)
}

func (service QuizService) UpdateQuiz(ctx context.Context, uid uint32, quiz *domain.Quiz) error {
	return service.repository.UpdateQuiz(ctx, uid, quiz)
}

func (service QuizService) UpdateQuizQuestions(ctx context.Context, uid uint32, questionsIDs []uint32) error {
	return service.repository.UpdateQuizQuestions(ctx, uid, questionsIDs)
}

func (service QuizService) DeleteQuiz(ctx context.Context, uid uint32) error {
	return service.repository.DeleteQuiz(ctx, uid)
}
