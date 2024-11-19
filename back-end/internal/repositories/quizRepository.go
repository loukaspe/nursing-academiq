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

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (repo *QuizRepository) GetQuiz(
	ctx context.Context,
	uid uint32,
) (*domain.Quiz, error) {
	var err error
	var modelQuiz *Quiz

	err = repo.db.WithContext(ctx).
		//Preload("Tutor").
		Model(Quiz{}).
		Preload("Questions.Answers").
		Preload("Course").
		Where("id = ?", uid).
		Take(&modelQuiz).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Quiz{}, err
	}

	domainQuestions := make([]domain.Question, 0, len(modelQuiz.Questions))
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
		}

		domainQuestions = append(domainQuestions, *domainQuestion)
	}

	// TODO: preload Tutor if needed
	//domainUser := domain.User{
	//	Username:    modelQuiz.User.Username,
	//	Password:    modelQuiz.User.Password,
	//	FirstName:   modelQuiz.User.FirstName,
	//	LastName:    modelQuiz.User.LastName,
	//	Email:       modelQuiz.User.Email,
	//	BirthDate:   modelQuiz.User.BirthDate,
	//	PhoneNumber: modelQuiz.User.PhoneNumber,
	//	Photo:       modelQuiz.User.Photo,
	//}

	return &domain.Quiz{
		Title:             modelQuiz.Title,
		Description:       modelQuiz.Description,
		Visibility:        modelQuiz.Visibility,
		ShowSubset:        modelQuiz.ShowSubset,
		SubsetSize:        modelQuiz.SubsetSize,
		ScoreSum:          modelQuiz.ScoreSum,
		MaxScore:          modelQuiz.MaxScore,
		NumberOfQuestions: len(modelQuiz.Questions),
		Questions:         domainQuestions,
		Course: &domain.Course{
			ID:          uint32(modelQuiz.CourseID),
			Title:       modelQuiz.Course.Title,
			Description: modelQuiz.Course.Description,
		},
	}, err
}

func (repo *QuizRepository) GetQuizByTutorID(
	ctx context.Context,
	tutorID uint32,
) ([]domain.Quiz, error) {
	var err error
	var modelTutor Tutor

	err = repo.db.WithContext(ctx).
		Preload("Courses.Quizs.Questions").
		First(&modelTutor, tutorID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("tutorID " + strconv.Itoa(int(tutorID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.Quiz{}, err
	}

	var domainQuizs []domain.Quiz
	for _, modelCourse := range modelTutor.Courses {
		courseName := modelCourse.Title

		for _, modelQuiz := range modelCourse.Quizs {
			var numberOfQuestions int
			for _, _ = range modelQuiz.Questions {
				numberOfQuestions++
			}

			domainQuizs = append(domainQuizs, domain.Quiz{
				ID:                uint32(modelQuiz.ID),
				Title:             modelQuiz.Title,
				Description:       modelQuiz.Description,
				Visibility:        modelQuiz.Visibility,
				ShowSubset:        modelQuiz.ShowSubset,
				SubsetSize:        modelQuiz.SubsetSize,
				ScoreSum:          modelQuiz.ScoreSum,
				MaxScore:          modelQuiz.MaxScore,
				NumberOfQuestions: numberOfQuestions,
				Course: &domain.Course{
					ID:    uint32(modelCourse.ID),
					Title: courseName,
				},
			})
		}
	}

	return domainQuizs, err
}

func (repo *QuizRepository) GetQuizByCourseID(
	ctx context.Context,
	courseID uint32,
) ([]domain.Quiz, error) {
	var err error
	var modelCourse Course

	err = repo.db.WithContext(ctx).
		Preload("Quizs.Questions").
		First(&modelCourse, courseID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(courseID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.Quiz{}, err
	}

	var domainQuizs []domain.Quiz

	for _, modelQuiz := range modelCourse.Quizs {
		domainQuestions := make([]domain.Question, 0, len(modelQuiz.Questions))
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

		domainQuizs = append(domainQuizs, domain.Quiz{
			ID:                uint32(modelQuiz.ID),
			Title:             modelQuiz.Title,
			Description:       modelQuiz.Description,
			Visibility:        modelQuiz.Visibility,
			ShowSubset:        modelQuiz.ShowSubset,
			SubsetSize:        modelQuiz.SubsetSize,
			ScoreSum:          modelQuiz.ScoreSum,
			MaxScore:          modelQuiz.MaxScore,
			NumberOfQuestions: len(modelQuiz.Questions),
			Course: &domain.Course{
				Title: modelCourse.Title,
			},
			Questions: domainQuestions,
		})
	}

	return domainQuizs, err
}

func (repo *QuizRepository) UpdateQuiz(
	ctx context.Context,
	uid uint32,
	quiz *domain.Quiz,
) error {
	modelQuiz := &Quiz{}

	err := repo.db.WithContext(ctx).Model(&Quiz{}).First(modelQuiz).Where("id = ?", uid).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	modelQuiz.Title = quiz.Title
	modelQuiz.Description = quiz.Description
	modelQuiz.Visibility = quiz.Visibility
	modelQuiz.ShowSubset = quiz.ShowSubset
	modelQuiz.SubsetSize = quiz.SubsetSize

	err = repo.db.WithContext(ctx).Save(&modelQuiz).Error

	return err
}

func (repo *QuizRepository) DeleteQuiz(
	ctx context.Context,
	uid uint32,
) error {
	db := repo.db.WithContext(ctx).Model(&Quiz{}).
		Where("id = ?", uid).
		Take(&Quiz{}).
		Delete(&Quiz{})

	if db.Error == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return db.Error
}

func (repo *QuizRepository) CreateQuiz(
	ctx context.Context,
	quiz *domain.Quiz,
) (uint, error) {
	var err error

	modelQuiz := Quiz{}
	modelQuiz.Title = quiz.Title
	modelQuiz.Description = quiz.Description
	modelQuiz.Visibility = quiz.Visibility
	modelQuiz.ShowSubset = quiz.ShowSubset
	modelQuiz.SubsetSize = quiz.SubsetSize
	modelQuiz.CourseID = uint(quiz.Course.ID)

	err = repo.db.WithContext(ctx).Create(&modelQuiz).Error

	return modelQuiz.ID, err
}
