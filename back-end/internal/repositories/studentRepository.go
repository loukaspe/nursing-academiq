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
