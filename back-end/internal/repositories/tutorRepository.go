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

type TutorRepository struct {
	db *gorm.DB
}

func NewTutorRepository(db *gorm.DB) *TutorRepository {
	return &TutorRepository{db: db}
}

func (repo *TutorRepository) CreateTutor(tutor *domain.Tutor) error {
	var err error

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
	// create custom error checking like Akis
	if err != nil {
		return err
	}

	modelTutor := Tutor{
		AcademicRank: tutor.AcademicRank,
		User:         modelUser,
	}

	err = repo.db.Debug().Create(&modelTutor).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *TutorRepository) GetTutor(uid uint32) (*domain.Tutor, error) {
	var err error
	var modelTutor *Tutor

	err = repo.db.Debug().Model(Tutor{}).Where("id = ?", uid).Take(&modelTutor).Error
	if err != nil {
		return &domain.Tutor{}, err
	}
	if err == gorm.ErrRecordNotFound {
		return &domain.Tutor{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("uid " + strconv.Itoa(int(uid)) + " not found"),
		}
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

func (repo *TutorRepository) UpdateTutor(uid uint32, tutor *domain.Tutor) error {
	modelTutor := Tutor{
		AcademicRank: tutor.AcademicRank,
	}

	db := repo.db.Debug().Model(&Tutor{}).
		Where("id = ?", uid).
		Take(&Tutor{}).
		UpdateColumns(
			map[string]interface{}{
				"academic_rank": modelTutor.AcademicRank,
				"updated_at":    time.Now(),
			},
		)
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

func (repo *TutorRepository) DeleteTutor(uid uint32) error {
	db := repo.db.Debug().Model(&Tutor{}).
		Where("id = ?", uid).
		Take(&Tutor{}).
		Delete(&Tutor{})

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
