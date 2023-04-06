package repositories

import (
	"errors"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (repo *StudentRepository) CreateStudent(student *domain.Student) error {
	var err error

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
	// create custom error checking like Akis
	if err != nil {
		return err
	}

	modelStudent := Student{
		RegistrationNumber: student.RegistrationNumber,
		User:               modelUser,
	}

	err = repo.db.Debug().Create(&modelStudent).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *StudentRepository) GetStudent(uid uint32) (*domain.Student, error) {
	var err error
	var modelStudent *Student

	err = repo.db.Debug().Model(Student{}).Where("id = ?", uid).Take(&modelStudent).Error
	if err != nil {
		return &domain.Student{}, err
	}
	if err == gorm.ErrRecordNotFound {
		return &domain.Student{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("uid " + strconv.Itoa(int(uid)) + " not found"),
		}
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

func (repo *StudentRepository) UpdateStudent(uid uint32, student *domain.Student) error {
	modelStudent := Student{
		RegistrationNumber: student.RegistrationNumber,
	}

	db := repo.db.Debug().Model(&Student{}).
		Where("id = ?", uid).
		Take(&Student{}).
		UpdateColumns(
			map[string]interface{}{
				"registration_number": modelStudent.RegistrationNumber,
				"updated_at":          time.Now(),
			},
		)

	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (repo *StudentRepository) DeleteStudent(uid uint32) error {
	db := repo.db.Debug().Model(&Student{}).
		Where("id = ?", uid).
		Take(&Student{}).
		Delete(&Student{})

	if db.Error == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("uid " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	if db.Error != nil {
		return db.Error
	}
	return nil
}
