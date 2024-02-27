package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
)

func NewQuizSessionByStudentService(repository ports.QuizSessionByStudentRepositoryInterface) *QuizSessionByStudentService {
	return &QuizSessionByStudentService{repository: repository}
}

type QuizSessionByStudentService struct {
	repository ports.QuizSessionByStudentRepositoryInterface
}

//func (service QuizSessionByStudentService) GetQuizSessionByStudent(ctx context.Context, uid uint32) (*domain.QuizSessionByStudent, error) {
//	return service.repository.GetQuizSessionByStudent(ctx, uid)
//}
//
//func (service QuizSessionByStudentService) GetQuizSessionsByStudent(ctx context.Context) ([]domain.QuizSessionByStudent, error) {
//	return service.repository.GetQuizSessionsByStudent(ctx)
//}

func (service QuizSessionByStudentService) GetQuizSessionByStudentID(ctx context.Context, studentID uint32) ([]domain.QuizSessionByStudent, error) {
	return service.repository.GetQuizSessionByStudentID(ctx, studentID)
}

//func (service QuizSessionByStudentService) GetQuizSessionByStudentByTutorID(ctx context.Context, tutorID uint32) ([]domain.QuizSessionByStudent, error) {
//	return service.repository.GetQuizSessionByStudentByTutorID(ctx, tutorID)
//}

//func (service QuizSessionByStudentService) CreateQuizSessionByStudent(ctx context.Context, quizSessionByStudent *domain.QuizSessionByStudent) (uint, error) {
//	return service.repository.CreateQuizSessionByStudent(ctx, quizSessionByStudent)
//}
//
//func (service QuizSessionByStudentService) UpdateQuizSessionByStudent(ctx context.Context, uid uint32, quizSessionByStudent *domain.QuizSessionByStudent) error {
//	return service.repository.UpdateQuizSessionByStudent(ctx, uid, quizSessionByStudent)
//}
//
//func (service QuizSessionByStudentService) DeleteQuizSessionByStudent(ctx context.Context, uid uint32) error {
//	return service.repository.DeleteQuizSessionByStudent(ctx, uid)
//}
