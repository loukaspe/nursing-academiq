package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (repo *StudentRepository) CreateStudent(
	ctx context.Context,
	student *domain.Student,
) (uint, error) {
	var err error

	modelStudent := Student{}

	modelUser := User{
		Username:    student.User.Username,
		Password:    student.User.Password,
		FirstName:   student.User.FirstName,
		LastName:    student.User.LastName,
		Email:       student.User.Email,
		BirthDate:   student.User.BirthDate,
		PhoneNumber: student.User.PhoneNumber,
		Photo:       student.User.Photo,
	}

	modelUser.prepare()
	err = modelUser.validate()
	if err != nil {
		return modelStudent.ID, apierrors.UserValidationError{
			ReturnedStatusCode: http.StatusBadRequest,
			OriginalError:      err,
		}
	}
	err = modelUser.BeforeSave()
	if err != nil {
		return 0, err
	}

	modelStudent.RegistrationNumber = student.RegistrationNumber
	modelStudent.User = modelUser

	err = repo.db.WithContext(ctx).Create(&modelStudent).Error

	return modelStudent.ID, err
}

func (repo *StudentRepository) GetStudent(
	ctx context.Context,
	uid uint32,
) (*domain.Student, error) {
	var err error
	var modelStudent *Student

	err = repo.db.WithContext(ctx).
		Preload("User").
		Model(Student{}).
		Where("id = ?", uid).
		Take(&modelStudent).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Student{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("studentID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Student{}, err
	}

	domainUser := domain.User{
		Username:    modelStudent.User.Username,
		Password:    modelStudent.User.Password,
		FirstName:   modelStudent.User.FirstName,
		LastName:    modelStudent.User.LastName,
		Email:       modelStudent.User.Email,
		BirthDate:   modelStudent.User.BirthDate,
		PhoneNumber: modelStudent.User.PhoneNumber,
		Photo:       modelStudent.User.Photo,
	}

	return &domain.Student{
		User:               domainUser,
		RegistrationNumber: modelStudent.RegistrationNumber,
	}, err
}

func (repo *StudentRepository) GetExtendedStudent(
	ctx context.Context,
	uid uint32,
) (*domain.Student, error) {
	var err error
	var modelStudent *Student

	err = repo.db.WithContext(ctx).
		Preload("User").
		Preload("QuizSessions").
		//Preload("QuizSessions.Quiz").
		Model(Student{}).
		Where("id = ?", uid).
		Take(&modelStudent).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Student{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("studentID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Student{}, err
	}

	completedQuizzesIDs := make(map[uint]struct{})
	totalQuizSessionsScore := float32(0)
	totalQuizSessionsMaxScore := 0

	for _, quizSession := range modelStudent.QuizSessions {
		if _, ok := completedQuizzesIDs[quizSession.QuizID]; !ok {
			completedQuizzesIDs[quizSession.QuizID] = struct{}{}
		}

		totalQuizSessionsScore += quizSession.Score
		totalQuizSessionsMaxScore += quizSession.MaxScore
	}

	domainUser := domain.User{
		Username:    modelStudent.User.Username,
		Password:    modelStudent.User.Password,
		FirstName:   modelStudent.User.FirstName,
		LastName:    modelStudent.User.LastName,
		Email:       modelStudent.User.Email,
		BirthDate:   modelStudent.User.BirthDate,
		PhoneNumber: modelStudent.User.PhoneNumber,
		Photo:       modelStudent.User.Photo,
	}

	questionsScore := fmt.Sprintf("%.1f/%d", totalQuizSessionsScore, totalQuizSessionsMaxScore)
	percentageOfCorrectAnswers := fmt.Sprintf("%.2f%%", (totalQuizSessionsScore/float32(totalQuizSessionsMaxScore))*100)

	return &domain.Student{
		User:                       domainUser,
		RegistrationNumber:         modelStudent.RegistrationNumber,
		CompletedQuizzes:           len(completedQuizzesIDs),
		QuestionsScore:             questionsScore,
		PercentageOfCorrectAnswers: percentageOfCorrectAnswers,
	}, err
}

func (repo *StudentRepository) UpdateStudent(
	ctx context.Context,
	uid uint32,
	student *domain.Student,
) error {
	modelUser := &User{}
	modelStudent := &Student{}

	err := repo.db.WithContext(ctx).Preload("User").Model(modelStudent).Where("id = ?", uid).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("studentID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	modelUser.Username = student.User.Username
	modelUser.Password = student.User.Password
	modelUser.FirstName = student.User.FirstName
	modelUser.LastName = student.User.LastName
	modelUser.Email = student.User.Email
	modelUser.BirthDate = student.User.BirthDate
	modelUser.PhoneNumber = student.User.PhoneNumber
	modelUser.Photo = student.User.Photo

	modelUser.prepare()
	err = modelUser.validate()
	if err != nil {
		return apierrors.UserValidationError{
			ReturnedStatusCode: http.StatusBadRequest,
			OriginalError:      err,
		}
	}

	modelStudent.RegistrationNumber = student.RegistrationNumber
	modelStudent.User = *modelUser

	err = repo.db.WithContext(ctx).Save(&modelStudent).Error

	return err
}

func (repo *StudentRepository) DeleteStudent(
	ctx context.Context,
	uid uint32,
) error {
	db := repo.db.WithContext(ctx).Model(&Student{}).
		Where("id = ?", uid).
		Take(&Student{}).
		Delete(&Student{})

	if db.Error == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("studentID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return db.Error
}

// RegisterCourses This takes a list of courses and add them to the student
// Any previous relationships are removed
func (repo *StudentRepository) RegisterCourses(
	ctx context.Context,
	studentID uint32,
	courses []domain.Course,
) (*domain.Student, error) {
	var err error
	var modelStudent Student
	var modelCourses []Course

	// firstly we check if the domain courses exist in the database
	for _, course := range courses {
		var modelCourse Course
		err = repo.db.WithContext(ctx).
			Where("id = ?", course.ID).
			First(&modelCourse).Error

		if err == gorm.ErrRecordNotFound {
			return nil, apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New("course " + strconv.Itoa(int(course.ID)) + " not found"),
			}
		}
		if err != nil {
			return nil, err
		}

		modelCourses = append(modelCourses, modelCourse)
	}

	err = repo.db.WithContext(ctx).
		Preload("Courses").
		Preload("User").
		First(&modelStudent, studentID).Error

	if err == gorm.ErrRecordNotFound {
		return nil, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("studentID " + strconv.Itoa(int(studentID)) + " not found"),
		}
	}
	if err != nil {
		return nil, err
	}

	if err := repo.db.WithContext(ctx).Model(&modelStudent).Association("Courses").Clear(); err != nil {
		return nil, err
	}

	for _, course := range modelCourses {
		modelStudent.Courses = append(modelStudent.Courses, course)
	}

	if err := repo.db.WithContext(ctx).Save(&modelStudent).Error; err != nil {
		return nil, err
	}

	var domainCourses []domain.Course
	for _, modelCourse := range modelStudent.Courses {
		domainCourses = append(domainCourses, domain.Course{
			ID:          uint32(modelCourse.ID),
			Title:       modelCourse.Title,
			Description: modelCourse.Description,
		})
	}

	return &domain.Student{
		User: domain.User{
			Username: modelStudent.User.Username,
			//Password:    modelStudent.User.Password,
			FirstName:   modelStudent.User.FirstName,
			LastName:    modelStudent.User.LastName,
			Email:       modelStudent.User.Email,
			BirthDate:   modelStudent.User.BirthDate,
			PhoneNumber: modelStudent.User.PhoneNumber,
			Photo:       modelStudent.User.Photo,
		},
		RegistrationNumber: modelStudent.RegistrationNumber,
		Courses:            domainCourses,
	}, err
}
