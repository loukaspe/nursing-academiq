package repositories

import (
	"context"
	"errors"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ChapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) *ChapterRepository {
	return &ChapterRepository{db: db}
}

func (repo *ChapterRepository) GetChapter(
	ctx context.Context,
	uid uint32,
) (*domain.Chapter, error) {
	var err error
	var modelChapter *Chapter

	err = repo.db.WithContext(ctx).
		Preload("Course").
		Model(Chapter{}).
		Where("id = ?", uid).
		Take(&modelChapter).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Chapter{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("chapterID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Chapter{}, err
	}

	return &domain.Chapter{
		Title:       modelChapter.Title,
		Description: modelChapter.Description,
		Course: &domain.Course{
			ID:    uint32(modelChapter.CourseID),
			Title: modelChapter.Course.Title,
		},
	}, err
}

func (repo *ChapterRepository) GetChapterByTitle(
	ctx context.Context,
	title string,
) (*domain.Chapter, error) {
	var err error
	var modelChapter *Chapter

	err = repo.db.WithContext(ctx).
		Preload("Course").
		Model(Chapter{}).
		Where("title = ?", title).
		Take(&modelChapter).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Chapter{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("chapterTitle " + title + " not found"),
		}
	}
	if err != nil {
		return &domain.Chapter{}, err
	}

	return &domain.Chapter{
		Title:       modelChapter.Title,
		Description: modelChapter.Description,
		Course: &domain.Course{
			ID:    uint32(modelChapter.CourseID),
			Title: modelChapter.Course.Title,
		},
	}, err
}

//func (repo *ChapterRepository) GetExtendedChapter(
//	ctx context.Context,
//	uid uint32,
//) (*domain.Chapter, error) {
//	var err error
//	var modelChapter *Chapter
//
//	err = repo.db.WithContext(ctx).
//		Preload("Quizs.Questions").
//		Preload("Chapters").
//		Preload("Tutor.User").
//		Model(Chapter{}).
//		Where("id = ?", uid).
//		Take(&modelChapter).Error
//
//	if err == gorm.ErrRecordNotFound {
//		return &domain.Chapter{}, apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("chapterID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//	if err != nil {
//		return &domain.Chapter{}, err
//	}
//
//	domainTutor := domain.Tutor{
//		User: domain.User{
//			FirstName: modelChapter.Tutor.User.FirstName,
//			LastName:  modelChapter.Tutor.User.LastName,
//		},
//	}
//
//	var domainQuizs []domain.Quiz
//
//	for _, modelQuiz := range modelChapter.Quizzes {
//		var numberOfQuestions int
//		for _, _ = range modelQuiz.Questions {
//			numberOfQuestions++
//		}
//
//		domainQuizs = append(domainQuizs, domain.Quiz{
//			Title:             modelQuiz.Title,
//			Description:       modelQuiz.Description,
//			Visibility:        modelQuiz.Visibility,
//			ShowSubset:        modelQuiz.ShowSubset,
//			SubsetSize:        modelQuiz.SubsetSize,
//			ScoreSum:          modelQuiz.ScoreSum,
//			MaxScore:          modelQuiz.MaxScore,
//			NumberOfQuestions: numberOfQuestions,
//			Chapter: &domain.Chapter{
//				Title: modelChapter.Title,
//			},
//		})
//	}
//
//	var domainChapters []domain.Chapter
//
//	for _, modelChapter := range modelChapter.Chapters {
//		domainChapters = append(domainChapters, domain.Chapter{
//			Title:       modelChapter.Title,
//			Description: modelChapter.Description,
//		})
//	}
//
//	return &domain.Chapter{
//		Title:       modelChapter.Title,
//		Description: modelChapter.Description,
//		Quizzes:     domainQuizs,
//		Chapters:    domainChapters,
//		Tutor:       &domainTutor,
//	}, err
//}
//
//func (repo *ChapterRepository) GetChapters(
//	ctx context.Context,
//) ([]domain.Chapter, error) {
//	var err error
//	var modelChapters []Chapter
//
//	err = repo.db.WithContext(ctx).
//		Model(Chapter{}).
//		Find(&modelChapters).Error
//
//	if err == gorm.ErrRecordNotFound {
//		return []domain.Chapter{}, apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("chapters not found"),
//		}
//	}
//	if err != nil {
//		return []domain.Chapter{}, err
//	}
//
//	var domainChapters []domain.Chapter
//	for _, modelChapter := range modelChapters {
//		domainChapters = append(domainChapters, domain.Chapter{
//			ID:          uint32(modelChapter.ID),
//			Title:       modelChapter.Title,
//			Description: modelChapter.Description,
//		})
//	}
//
//	return domainChapters, err
//}
//
//func (repo *ChapterRepository) GetChapterByTutorID(
//	ctx context.Context,
//	tutorID uint32,
//) ([]domain.Chapter, error) {
//	var err error
//	var modelTutor Tutor
//
//	err = repo.db.WithContext(ctx).
//		Preload("Chapters").
//		First(&modelTutor, tutorID).Error
//
//	if err == gorm.ErrRecordNotFound {
//		return []domain.Chapter{}, apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("tutorID " + strconv.Itoa(int(tutorID)) + " not found"),
//		}
//	}
//	if err != nil {
//		return []domain.Chapter{}, err
//	}
//
//	var domainChapters []domain.Chapter
//	for _, modelChapter := range modelTutor.Chapters {
//		// TODO: preload Tutor, Students if needed
//		domainChapters = append(domainChapters, domain.Chapter{
//			ID:          uint32(modelChapter.ID),
//			Title:       modelChapter.Title,
//			Description: modelChapter.Description,
//		})
//	}
//
//	return domainChapters, err
//}

func (repo *ChapterRepository) UpdateChapter(
	ctx context.Context,
	uid uint32,
	chapter *domain.Chapter,
) error {
	modelChapter := &Chapter{}

	// Only Title and Description are editable
	partialUpdates := make(map[string]interface{}, 2)
	if chapter.Title != "" {
		partialUpdates["title"] = chapter.Title
	}
	if chapter.Description != "" {
		partialUpdates["description"] = chapter.Description
	}

	err := repo.db.WithContext(ctx).Model(modelChapter).Where("id = ?", uid).Updates(partialUpdates).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("chapterID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return err
}

func (repo *ChapterRepository) DeleteChapter(
	ctx context.Context,
	uid uint32,
) error {
	err := repo.db.WithContext(ctx).Model(&Chapter{}).
		Where("id = ?", uid).
		Take(&Chapter{}).
		Delete(&Chapter{}).Error

	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("chapterID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return err
}

func (repo *ChapterRepository) CreateChapter(
	ctx context.Context,
	chapter *domain.Chapter,
) (uint, error) {
	var err error

	modelChapter := Chapter{}
	modelChapter.Title = chapter.Title
	modelChapter.Description = chapter.Description
	modelChapter.CourseID = uint(chapter.Course.ID)

	err = repo.db.WithContext(ctx).Create(&modelChapter).Error

	return modelChapter.ID, err
}
