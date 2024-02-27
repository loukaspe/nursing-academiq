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

type QuizSessionRepository struct {
	db *gorm.DB
}

func NewQuizSessionRepository(db *gorm.DB) *QuizSessionRepository {
	return &QuizSessionRepository{db: db}
}

//func (repo *QuizSessionRepository) GetQuizSession(
//	ctx context.Context,
//	uid uint32,
//) (*domain.QuizSessionByStudent, error) {
//	var err error
//	var modelQuizSession *QuizSession
//
//	err = repo.db.WithContext(ctx).
//		Preload("Tutor").
//		Model(QuizSession{}).
//		Where("id = ?", uid).
//		Take(&modelQuizSession).Error
//
//	if err == gorm.ErrRecordNotFound {
//		return &domain.QuizSessionByStudent{}, apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("quizSessionID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//	if err != nil {
//		return &domain.QuizSessionByStudent{}, err
//	}
//
//	// TODO: preload Tutor if needed
//	//domainUser := domain.User{
//	//	Username:    modelQuizSession.User.Username,
//	//	Password:    modelQuizSession.User.Password,
//	//	FirstName:   modelQuizSession.User.FirstName,
//	//	LastName:    modelQuizSession.User.LastName,
//	//	Email:       modelQuizSession.User.Email,
//	//	BirthDate:   modelQuizSession.User.BirthDate,
//	//	PhoneNumber: modelQuizSession.User.PhoneNumber,
//	//	Photo:       modelQuizSession.User.Photo,
//	//}
//
//	return &domain.QuizSessionByStudent{
//		Title:       modelQuizSession.Title,
//		Description: modelQuizSession.Description,
//		Students:    nil,
//	}, err
//}

//func (repo *QuizSessionRepository) GetQuizSessions(
//	ctx context.Context,
//) ([]domain.QuizSessionByStudent, error) {
//	var err error
//	var modelQuizSessions []QuizSession
//
//	err = repo.db.WithContext(ctx).
//		Model(QuizSession{}).
//		Find(&modelQuizSessions).Error
//
//	if err == gorm.ErrRecordNotFound {
//		return []domain.QuizSessionByStudent{}, apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("quizSessions not found"),
//		}
//	}
//	if err != nil {
//		return []domain.QuizSessionByStudent{}, err
//	}
//
//	var domainQuizSessions []domain.QuizSessionByStudent
//	for _, modelQuizSession := range modelQuizSessions {
//		domainQuizSessions = append(domainQuizSessions, domain.QuizSessionByStudent{
//			ID:          uint32(modelQuizSession.ID),
//			Title:       modelQuizSession.Title,
//			Description: modelQuizSession.Description,
//		})
//	}
//
//	return domainQuizSessions, err
//}

func (repo *QuizSessionRepository) GetQuizSessionByStudentID(
	ctx context.Context,
	studentID uint32,
) ([]domain.QuizSessionByStudent, error) {
	var err error
	//var modelQuizSessions []QuizSession
	var modelStudent Student

	//err = repo.db.WithContext(ctx).
	//	//Preload("QuizSessions").
	//	Where("id = ?", studentID).
	//	Take(&modelStudent).
	//	Association("QuizSessions").
	//	Find(&modelQuizSessions)

	//err = repo.db.WithContext(ctx).
	//	Joins("QuizSessions").
	//	Model(Student{}).
	//	Where("id = ?", studentID).
	//	Take(&modelStudent).Error

	err = repo.db.WithContext(ctx).
		Preload("QuizSessions.Quiz").
		First(&modelStudent, studentID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.QuizSessionByStudent{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("studentID " + strconv.Itoa(int(studentID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.QuizSessionByStudent{}, err
	}

	var domainQuizSessions []domain.QuizSessionByStudent
	for _, modelQuizSession := range modelStudent.QuizSessions {
		// TODO: preload Tutor, Students if needed
		domainQuizSessions = append(domainQuizSessions, domain.QuizSessionByStudent{
			Quiz: &domain.Quiz{
				Title: modelQuizSession.Quiz.Title,
			},
			Date:              modelQuizSession.DateTime,
			DurationInSeconds: modelQuizSession.DurationInSeconds,
			Score:             modelQuizSession.Score,
			MaxScore:          modelQuizSession.MaxScore,
		})
	}

	return domainQuizSessions, err
}

//func (repo *QuizSessionRepository) GetQuizSessionByTutorID(
//	ctx context.Context,
//	tutorID uint32,
//) ([]domain.QuizSessionByStudent, error) {
//	var err error
//	var modelTutor Tutor
//
//	err = repo.db.WithContext(ctx).
//		Preload("QuizSessions").
//		First(&modelTutor, tutorID).Error
//
//	if err == gorm.ErrRecordNotFound {
//		return []domain.QuizSessionByStudent{}, apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("tutorID " + strconv.Itoa(int(tutorID)) + " not found"),
//		}
//	}
//	if err != nil {
//		return []domain.QuizSessionByStudent{}, err
//	}
//
//	var domainQuizSessions []domain.QuizSessionByStudent
//	for _, modelQuizSession := range modelTutor.QuizSessions {
//		// TODO: preload Tutor, Students if needed
//		domainQuizSessions = append(domainQuizSessions, domain.QuizSessionByStudent{
//			ID:          uint32(modelQuizSession.ID),
//			Title:       modelQuizSession.Title,
//			Description: modelQuizSession.Description,
//		})
//	}
//
//	return domainQuizSessions, err
//}

//func (repo *QuizSessionRepository) UpdateQuizSession(
//	ctx context.Context,
//	uid uint32,
//	quizSession *domain.QuizSessionByStudent,
//) error {
//	modelQuizSession := &QuizSession{}
//
//	// TODO: handle Tutor preload if needed err := repo.db.WithContext(ctx).Preload("User").First(modelQuizSession).Error
//	err := repo.db.WithContext(ctx).Model(modelQuizSession).Where("id = ?", uid).Error
//	if err == gorm.ErrRecordNotFound {
//		return apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("quizSessionID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//	if err != nil {
//		return err
//	}
//
//	modelQuizSession.Title = quizSession.Title
//	modelQuizSession.Description = quizSession.Description
//
//	err = repo.db.WithContext(ctx).Save(&modelQuizSession).Error
//
//	return err
//}
//
//func (repo *QuizSessionRepository) DeleteQuizSession(
//	ctx context.Context,
//	uid uint32,
//) error {
//	db := repo.db.WithContext(ctx).Model(&QuizSession{}).
//		Where("id = ?", uid).
//		Take(&QuizSession{}).
//		Delete(&QuizSession{})
//
//	if db.Error == gorm.ErrRecordNotFound {
//		return apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("quizSessionID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//
//	return db.Error
//}

//func (repo *QuizSessionRepository) CreateQuizSession(
//	ctx context.Context,
//	quizSession *domain.QuizSessionByStudent,
//	tutorID uint,
//) (uint, error) {
//	var err error
//
//	// TODO: add quizSession tutor if needed
//	modelQuizSession := QuizSession{}
//	modelQuizSession.Title = quizSession.Title
//	modelQuizSession.Description = quizSession.Description
//	modelQuizSession.TutorID = tutorID
//
//	err = repo.db.WithContext(ctx).Create(&modelQuizSession).Error
//
//	return modelQuizSession.ID, err
//}
