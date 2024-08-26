package services

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/ports"
)

func NewChapterService(
	chapterRepository ports.ChapterRepositoryInterface,
	quizRepository ports.QuizRepositoryInterface) *ChapterService {
	return &ChapterService{
		chapterRepository: chapterRepository,
		quizRepository:    quizRepository,
	}
}

type ChapterService struct {
	chapterRepository ports.ChapterRepositoryInterface
	quizRepository    ports.QuizRepositoryInterface
}

func (service ChapterService) GetChapter(ctx context.Context, uid uint32) (*domain.Chapter, error) {
	domainChapter, err := service.chapterRepository.GetChapter(ctx, uid)
	if err != nil {
		return nil, err
	}

	chapterQuizs := []domain.Quiz{}

	quizzes, err := service.quizRepository.GetQuizByCourseID(ctx, domainChapter.Course.ID)
	for _, quiz := range quizzes {
		for _, question := range quiz.Questions {
			if question.Chapter.ID == uid {
				chapterQuizs = append(chapterQuizs, quiz)
				break
			}
		}
	}

	domainChapter.Quizzes = chapterQuizs

	return domainChapter, nil
}

func (service ChapterService) CreateChapter(ctx context.Context, course *domain.Chapter) (uint, error) {
	return service.chapterRepository.CreateChapter(ctx, course)
}

func (service ChapterService) UpdateChapter(ctx context.Context, uid uint32, course *domain.Chapter) error {
	return service.chapterRepository.UpdateChapter(ctx, uid, course)
}

func (service ChapterService) DeleteChapter(ctx context.Context, uid uint32) error {
	return service.chapterRepository.DeleteChapter(ctx, uid)
}

//
//func (service ChapterService) GetExtendedChapter(ctx context.Context, uid uint32) (*domain.Chapter, error) {
//	return service.repository.GetExtendedChapter(ctx, uid)
//}
