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

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (repo *QuestionRepository) ImportForCourse(
	ctx context.Context,
	questions []domain.Question,
	courseID uint,
) error {
	var err error

	modelQuestions := make([]*Question, 0, len(questions))
	err = repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, question := range questions {
			var chapter Chapter
			result := tx.Where("title = ?", question.Chapter.Title).First(&chapter)
			if result.Error == gorm.ErrRecordNotFound {
				// Chapter does not exist, create it.
				chapter = Chapter{
					Title:    question.Chapter.Title,
					CourseID: courseID,
				}
				err = tx.Create(&chapter).Error
				if err != nil {
					return err
				}
			}

			modelAnswers := make([]Answer, 0)
			for _, answer := range question.Answers {
				modelAnswer := Answer{}
				modelAnswer.Text = answer.Text
				modelAnswer.IsCorrect = answer.IsCorrect

				modelAnswers = append(modelAnswers, modelAnswer)
			}

			modelQuestion := Question{}
			modelQuestion.Text = question.Text
			modelQuestion.Explanation = question.Explanation
			modelQuestion.Source = question.Source
			modelQuestion.MultipleCorrectAnswers = question.MultipleCorrectAnswers
			modelQuestion.NumberOfAnswers = question.NumberOfAnswers
			modelQuestion.Answers = modelAnswers
			modelQuestion.ChapterID = chapter.ID

			err = tx.Create(&modelQuestion).Error
			if err != nil {
				return err
			}
			modelQuestions = append(modelQuestions, &modelQuestion)
		}

		return nil
	})

	if err != nil {
		return err
	}

	// TODO remove fake quiz create
	modelQuiz := Quiz{}
	modelQuiz.Title = "Test Quiz"
	modelQuiz.Description = "This is just a test quiz"
	modelQuiz.CourseID = courseID
	modelQuiz.Questions = modelQuestions

	return repo.db.WithContext(ctx).Create(&modelQuiz).Error
}

func (repo *QuestionRepository) CreateQuestion(
	ctx context.Context,
	question *domain.Question,
) (uint, error) {
	var err error

	modelQuestion := Question{}
	modelQuestion.Text = question.Text
	modelQuestion.Explanation = question.Explanation
	modelQuestion.Source = question.Source
	modelQuestion.MultipleCorrectAnswers = question.MultipleCorrectAnswers
	modelQuestion.NumberOfAnswers = question.NumberOfAnswers
	modelQuestion.CourseID = uint(question.Course.ID)
	modelQuestion.ChapterID = uint(question.Chapter.ID)

	err = repo.db.WithContext(ctx).Create(&modelQuestion).Error

	return modelQuestion.ID, err
}

func (repo *QuestionRepository) GetQuestionsByCourseID(
	ctx context.Context,
	courseID uint32,
) ([]domain.Question, error) {
	var err error
	var modelCourse Course

	err = repo.db.WithContext(ctx).
		Preload("Quizs.Questions").
		First(&modelCourse, courseID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Question{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(courseID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.Question{}, err
	}

	domainQuestions := make([]domain.Question, 0)

	for _, modelQuiz := range modelCourse.Quizs {
		for _, modelQuestion := range modelQuiz.Questions {
			domainAnswers := make([]domain.Answer, 0, modelQuestion.NumberOfAnswers)
			for _, modelAnswer := range modelQuestion.Answers {
				domainAnswer := &domain.Answer{
					Text:      modelAnswer.Text,
					IsCorrect: modelAnswer.IsCorrect,
				}

				domainAnswers = append(domainAnswers, *domainAnswer)
			}

			domainQuestion := &domain.Question{
				Text:                   modelQuestion.Text,
				Explanation:            modelQuestion.Explanation,
				Source:                 modelQuestion.Source,
				MultipleCorrectAnswers: modelQuestion.MultipleCorrectAnswers,
				NumberOfAnswers:        modelQuestion.NumberOfAnswers,
				Answers:                domainAnswers,
				Chapter: &domain.Chapter{
					ID: uint32(modelQuestion.ChapterID),
				},
			}

			domainQuestions = append(domainQuestions, *domainQuestion)
		}
	}

	return domainQuestions, err
}

func (repo *QuestionRepository) UpdateQuestion(
	ctx context.Context,
	uid uint32,
	question *domain.Question,
) error {
	modelQuestion := &Question{}

	err := repo.db.WithContext(ctx).Model(&Question{}).First(modelQuestion).Where("id = ?", uid).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questionID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	modelQuestion.Text = question.Text
	modelQuestion.Explanation = question.Explanation
	modelQuestion.Source = question.Source
	modelQuestion.MultipleCorrectAnswers = question.MultipleCorrectAnswers
	modelQuestion.NumberOfAnswers = question.NumberOfAnswers
	modelQuestion.CourseID = uint(question.Course.ID)
	modelQuestion.ChapterID = uint(question.Chapter.ID)
	err = repo.db.WithContext(ctx).Save(&modelQuestion).Error

	return err
}

func (repo *QuestionRepository) DeleteQuestion(
	ctx context.Context,
	uid uint32,
) error {
	db := repo.db.WithContext(ctx).Model(&Question{}).
		Where("id = ?", uid).
		Take(&Question{}).
		Delete(&Question{})

	if db.Error == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questionID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return db.Error
}

func (repo *QuestionRepository) GetQuestion(
	ctx context.Context,
	uid uint32,
) (*domain.Question, error) {
	var err error
	var modelQuestion *Question

	err = repo.db.WithContext(ctx).
		//Preload("Tutor").
		Model(Question{}).
		Preload("Answers").
		Where("id = ?", uid).
		Take(&modelQuestion).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Question{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questionID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Question{}, err
	}

	domainAnswers := make([]domain.Answer, 0, modelQuestion.NumberOfAnswers)
	for _, modelAnswer := range modelQuestion.Answers {
		domainAnswer := &domain.Answer{
			Text:      modelAnswer.Text,
			IsCorrect: modelAnswer.IsCorrect,
		}

		domainAnswers = append(domainAnswers, *domainAnswer)
	}

	return &domain.Question{
		Text:                   modelQuestion.Text,
		Explanation:            modelQuestion.Explanation,
		Source:                 modelQuestion.Source,
		MultipleCorrectAnswers: modelQuestion.MultipleCorrectAnswers,
		NumberOfAnswers:        modelQuestion.NumberOfAnswers,
		Answers:                domainAnswers,
		Chapter: &domain.Chapter{
			ID: uint32(modelQuestion.ChapterID),
		},
		Course: &domain.Course{
			ID: uint32(modelQuestion.CourseID),
		},
	}, err
}

func (repo *QuestionRepository) GetQuestions(
	ctx context.Context,
) ([]domain.Question, error) {
	var err error
	var modelQuestions []Question

	err = repo.db.WithContext(ctx).
		Model(Question{}).
		Find(&modelQuestions).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Question{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questions not found"),
		}
	}
	if err != nil {
		return []domain.Question{}, err
	}

	var domainQuestions []domain.Question
	for _, modelQuestion := range modelQuestions {
		domainAnswers := make([]domain.Answer, 0, modelQuestion.NumberOfAnswers)
		for _, modelAnswer := range modelQuestion.Answers {
			domainAnswer := &domain.Answer{
				Text:      modelAnswer.Text,
				IsCorrect: modelAnswer.IsCorrect,
			}

			domainAnswers = append(domainAnswers, *domainAnswer)
		}

		domainQuestions = append(domainQuestions, domain.Question{
			Text:                   modelQuestion.Text,
			Explanation:            modelQuestion.Explanation,
			Source:                 modelQuestion.Source,
			MultipleCorrectAnswers: modelQuestion.MultipleCorrectAnswers,
			NumberOfAnswers:        modelQuestion.NumberOfAnswers,
			Answers:                domainAnswers,
			Chapter: &domain.Chapter{
				ID: uint32(modelQuestion.ChapterID),
			},
			Course: &domain.Course{
				ID: uint32(modelQuestion.CourseID),
			},
		})
	}

	return domainQuestions, err
}
