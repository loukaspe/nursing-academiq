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

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Login(
	ctx context.Context,
	username,
	password string,
) (*domain.User, uint, error) {
	var err error
	var modelUser *User

	err = repo.db.WithContext(ctx).
		Model(User{}).
		Where("username = ?", username).
		Take(&modelUser).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.User{}, 0, &apierrors.LoginError{
			ReturnedStatusCode: http.StatusUnauthorized,
			OriginalError:      errors.New("user with name " + username + " not found"),
		}
	}
	if err != nil {
		return &domain.User{}, 0, err
	}

	var modelTutor Tutor

	repo.db.WithContext(ctx).
		Where("user_id = ?", modelUser.ID).
		First(&modelTutor)

	err = VerifyPassword(modelUser.Password, password)
	if err != nil {
		return &domain.User{}, 0, &apierrors.LoginError{
			ReturnedStatusCode: http.StatusUnauthorized,
			OriginalError:      err,
		}
	}

	var userType string
	var specificID uint

	userType = "tutor"
	specificID = modelTutor.ID

	return &domain.User{
		Username:   modelUser.Username,
		Password:   modelUser.Password,
		FirstName:  modelUser.FirstName,
		LastName:   modelUser.LastName,
		Email:      modelUser.Email,
		UserType:   userType,
		SpecificID: specificID,
	}, modelUser.ID, err
}

func (repo *UserRepository) ChangeUserPassword(
	ctx context.Context,
	uid uint32,
	oldPassword,
	newPassword string,
) error {
	modelUser := &User{}

	err := repo.db.WithContext(ctx).Model(User{}).Where("id = ?", uid).Take(modelUser).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("UserID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	err = VerifyPassword(modelUser.Password, oldPassword)
	if err != nil {
		return &apierrors.PasswordMismatchError{
			ReturnedStatusCode: http.StatusInternalServerError,
			OriginalError:      err,
		}
	}

	modelUser.Password = newPassword
	err = modelUser.BeforeSave()
	if err != nil {
		return err
	}

	err = repo.db.WithContext(ctx).Save(&modelUser).Error

	return err
}
