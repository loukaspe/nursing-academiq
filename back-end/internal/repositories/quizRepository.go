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
		Title:       modelQuiz.Title,
		Description: modelQuiz.Description,
		Visibility:  modelQuiz.Visibility,
		ShowSubset:  modelQuiz.ShowSubset,
		SubsetSize:  modelQuiz.SubsetSize,
		ScoreSum:    modelQuiz.ScoreSum,
		MaxScore:    modelQuiz.MaxScore,
		//Questions:        nil,
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
				Title:             modelQuiz.Title,
				Description:       modelQuiz.Description,
				Visibility:        modelQuiz.Visibility,
				ShowSubset:        modelQuiz.ShowSubset,
				SubsetSize:        modelQuiz.SubsetSize,
				ScoreSum:          modelQuiz.ScoreSum,
				MaxScore:          modelQuiz.MaxScore,
				NumberOfQuestions: numberOfQuestions,
				Course: &domain.Course{
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
		var numberOfQuestions int
		for _, _ = range modelQuiz.Questions {
			numberOfQuestions++
		}

		domainQuizs = append(domainQuizs, domain.Quiz{
			Title:             modelQuiz.Title,
			Description:       modelQuiz.Description,
			Visibility:        modelQuiz.Visibility,
			ShowSubset:        modelQuiz.ShowSubset,
			SubsetSize:        modelQuiz.SubsetSize,
			ScoreSum:          modelQuiz.ScoreSum,
			MaxScore:          modelQuiz.MaxScore,
			NumberOfQuestions: numberOfQuestions,
			Course: &domain.Course{
				Title: modelCourse.Title,
			},
		})
	}

	return domainQuizs, err
}

//func (repo *QuizRepository) UpdateQuiz(
//	ctx context.Context,
//	uid uint32,
//	quiz *domain.Quiz,
//) error {
//	modelQuiz := &Quiz{}
//
//	// TODO: handle Tutor preload if needed err := repo.db.WithContext(ctx).Preload("User").First(modelQuiz).Error
//	err := repo.db.WithContext(ctx).Model(modelQuiz).Where("id = ?", uid).Error
//	if err == gorm.ErrRecordNotFound {
//		return apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//	if err != nil {
//		return err
//	}
//
//	modelQuiz.Title = quiz.Title
//	modelQuiz.Description = quiz.Description
//
//	err = repo.db.WithContext(ctx).Save(&modelQuiz).Error
//
//	return err
//}
//
//func (repo *QuizRepository) DeleteQuiz(
//	ctx context.Context,
//	uid uint32,
//) error {
//	db := repo.db.WithContext(ctx).Model(&Quiz{}).
//		Where("id = ?", uid).
//		Take(&Quiz{}).
//		Delete(&Quiz{})
//
//	if db.Error == gorm.ErrRecordNotFound {
//		return apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//
//	return db.Error
//}

//func (repo *QuizRepository) CreateQuiz(
//	ctx context.Context,
//	quiz *domain.Quiz,
//	tutorID uint,
//) (uint, error) {
//	var err error
//
//	// TODO: add quiz tutor if needed
//	modelQuiz := Quiz{}
//	modelQuiz.Title = quiz.Title
//	modelQuiz.Description = quiz.Description
//	modelQuiz.TutorID = tutorID
//
//	err = repo.db.WithContext(ctx).Create(&modelQuiz).Error
//
//	return modelQuiz.ID, err
//}
