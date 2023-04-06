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

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CreateUser(user *domain.User) error {
	var err error

	modelUser := User{
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		BirthDate:   user.BirthDate,
		PhoneNumber: user.PhoneNumber,
		Photo:       user.Photo,
	}

	modelUser.prepare()
	err = modelUser.validate()
	// create custom error checking like Akis
	if err != nil {
		return err
	}

	err = modelUser.BeforeSave()
	if err != nil {
		return err
	}

	err = repo.db.Debug().Create(&modelUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUser(uid uint32) (*domain.User, error) {
	var err error
	var modelUser *User

	err = repo.db.Debug().Model(User{}).Where("id = ?", uid).Take(&modelUser).Error
	if err != nil {
		return &domain.User{}, err
	}
	if err == gorm.ErrRecordNotFound {
		return &domain.User{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("uid " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return &domain.User{
		Username:    modelUser.Username,
		Password:    modelUser.Password,
		FirstName:   modelUser.FirstName,
		LastName:    modelUser.LastName,
		Email:       modelUser.Email,
		BirthDate:   modelUser.BirthDate,
		PhoneNumber: modelUser.PhoneNumber,
		Photo:       modelUser.Photo,
	}, err
}

func (repo *UserRepository) UpdateUser(uid uint32, user *domain.User) error {
	var err error

	modelUser := User{
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		BirthDate:   user.BirthDate,
		PhoneNumber: user.PhoneNumber,
		Photo:       user.Photo,
	}

	err = modelUser.BeforeSave()
	if err != nil {
		return err
	}

	err = modelUser.validate()
	// create custom error checking like Akis
	if err != nil {
		return err
	}

	db := repo.db.Debug().Model(&User{}).
		Where("id = ?", uid).
		Take(&User{}).
		UpdateColumns(
			map[string]interface{}{
				"password":     modelUser.Password,
				"email":        modelUser.Email,
				"username":     modelUser.Username,
				"first_name":   modelUser.FirstName,
				"last_name":    modelUser.LastName,
				"phone_number": modelUser.PhoneNumber,
				"birth_date":   modelUser.BirthDate,
				"photo":        modelUser.Photo,
				"updated_at":   time.Now(),
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

func (repo *UserRepository) DeleteUser(uid uint32) error {
	db := repo.db.Debug().Model(&User{}).
		Where("id = ?", uid).
		Take(&User{}).
		Delete(&User{})

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
