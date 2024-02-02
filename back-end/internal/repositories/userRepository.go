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

	// Get either the Tutor or Student Information
	var modelStudent Student
	var modelTutor Tutor
	repo.db.WithContext(ctx).
		Where("user_id = ?", modelUser.ID).
		First(&modelStudent)
	if modelStudent.ID == 0 {
		repo.db.WithContext(ctx).
			Where("user_id = ?", modelUser.ID).
			First(&modelTutor)
	}

	err = VerifyPassword(modelUser.Password, password)
	if err != nil {
		return &domain.User{}, 0, &apierrors.LoginError{
			ReturnedStatusCode: http.StatusUnauthorized,
			OriginalError:      err,
		}
	}

	var userType string
	var specificID uint
	if modelStudent.ID != 0 {
		userType = "student"
		specificID = modelStudent.ID
	} else if modelTutor.ID != 0 {
		userType = "tutor"
		specificID = modelTutor.ID
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
		UserType:    userType,
		SpecificID:  specificID,
	}, modelUser.ID, err
}
