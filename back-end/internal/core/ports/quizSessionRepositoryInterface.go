package ports

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
)

type QuizSessionByStudentRepositoryInterface interface {
	//GetQuizSessionByStudent(ctx context.Context, uid uint32) (*domain.QuizSessionByStudent, error)
	//GetQuizSessionsByStudent(ctx context.Context) ([]domain.QuizSessionByStudent, error)
	GetQuizSessionByStudentID(ctx context.Context, studentID uint32) ([]domain.QuizSessionByStudent, error)
	//GetQuizSessionByStudentByTutorID(ctx context.Context, tutorID uint32) ([]domain.QuizSessionByStudent, error)
	//CreateQuizSessionByStudent(context.Context, *domain.QuizSessionByStudent) (uint, error)
	//UpdateQuizSessionByStudent(context.Context, uint32, *domain.QuizSessionByStudent) error
	//DeleteQuizSessionByStudent(ctx context.Context, uid uint32) error
}
