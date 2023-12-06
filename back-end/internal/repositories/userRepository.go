package repositories

import (
	"context"
	"errors"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"gorm.io/gorm"
	"net/http"
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
) (*domain.User, error) {
	var err error
	var modelUser *User

	err = repo.db.WithContext(ctx).
		Model(User{}).
		Where("username = ?", username).
		Take(&modelUser).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.User{}, &apierrors.LoginError{
			ReturnedStatusCode: http.StatusUnauthorized,
			OriginalError:      errors.New("user with name " + username + " not found"),
		}
	}
	if err != nil {
		return &domain.User{}, err
	}

	err = VerifyPassword(modelUser.Password, password)
	if err != nil {
		return &domain.User{}, &apierrors.LoginError{
			ReturnedStatusCode: http.StatusUnauthorized,
			OriginalError:      err,
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
