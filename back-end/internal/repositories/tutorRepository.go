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

type TutorRepository struct {
	db *gorm.DB
}

func NewTutorRepository(db *gorm.DB) *TutorRepository {
	return &TutorRepository{db: db}
}

func (repo *TutorRepository) CreateTutor(
	ctx context.Context,
	tutor *domain.Tutor,
) (uint, error) {
	var err error

	modelTutor := Tutor{}

	modelUser := User{
		Username:    tutor.User.Username,
		Password:    tutor.User.Password,
		FirstName:   tutor.User.FirstName,
		LastName:    tutor.User.LastName,
		Email:       tutor.User.Email,
		BirthDate:   tutor.User.BirthDate,
		PhoneNumber: tutor.User.PhoneNumber,
		Photo:       tutor.User.Photo,
	}

	modelUser.prepare()
	err = modelUser.validate()
	if err != nil {
		return modelTutor.ID, apierrors.UserValidationError{
			ReturnedStatusCode: http.StatusBadRequest,
			OriginalError:      err,
		}
	}

	modelTutor.AcademicRank = tutor.AcademicRank
	modelTutor.User = modelUser

	err = repo.db.WithContext(ctx).Create(&modelTutor).Error

	return modelTutor.ID, err
}

func (repo *TutorRepository) GetTutor(
	ctx context.Context,
	uid uint32,
) (*domain.Tutor, error) {
	var err error
	var modelTutor *Tutor

	err = repo.db.WithContext(ctx).
		Preload("User").
		Model(Tutor{}).
		Where("id = ?", uid).
		Take(&modelTutor).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Tutor{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("uid " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Tutor{}, err
	}

	domainUser := domain.User{
		Username:    modelTutor.User.Username,
		Password:    modelTutor.User.Password,
		FirstName:   modelTutor.User.FirstName,
		LastName:    modelTutor.User.LastName,
		Email:       modelTutor.User.Email,
		BirthDate:   modelTutor.User.BirthDate,
		PhoneNumber: modelTutor.User.PhoneNumber,
		Photo:       modelTutor.User.Photo,
	}

	return &domain.Tutor{
		User:         domainUser,
		AcademicRank: modelTutor.AcademicRank,
	}, err
}

func (repo *TutorRepository) UpdateTutor(
	ctx context.Context,
	uid uint32,
	tutor *domain.Tutor,
) error {
	modelUser := &User{}
	modelTutor := &Tutor{}

	err := repo.db.WithContext(ctx).Preload("User").First(modelTutor).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("uid " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	modelUser.Username = tutor.User.Username
	modelUser.Password = tutor.User.Password
	modelUser.FirstName = tutor.User.FirstName
	modelUser.LastName = tutor.User.LastName
	modelUser.Email = tutor.User.Email
	modelUser.BirthDate = tutor.User.BirthDate
	modelUser.PhoneNumber = tutor.User.PhoneNumber
	modelUser.Photo = tutor.User.Photo

	modelUser.prepare()
	err = modelUser.validate()
	if err != nil {
		return apierrors.UserValidationError{
			ReturnedStatusCode: http.StatusBadRequest,
			OriginalError:      err,
		}
	}

	modelTutor.AcademicRank = tutor.AcademicRank
	modelTutor.User = *modelUser

	err = repo.db.WithContext(ctx).Save(&modelTutor).Error

	return err
}

func (repo *TutorRepository) DeleteTutor(
	ctx context.Context,
	uid uint32,
) error {
	db := repo.db.WithContext(ctx).Model(&Tutor{}).
		Where("id = ?", uid).
		Take(&Tutor{}).
		Delete(&Tutor{})

	if db.Error == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("uid " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return db.Error
}
