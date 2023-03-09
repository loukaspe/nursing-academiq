package repositories

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	Username    string `gorm:"not null;"`
	Password    string `gorm:"not null;"`
	FirstName   string `gorm:"not null;"`
	LastName    string `gorm:"not null;"`
	Email       string `gorm:"not null;"`
	BirthDate   time.Time
	PhoneNumber string
	Photo       string
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) prepare() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Photo = html.EscapeString(strings.TrimSpace(u.Photo))
	u.PhoneNumber = html.EscapeString(strings.TrimSpace(u.PhoneNumber))
}

func (u *User) validate() error {
	if u.Password == "" {
		return errors.New("Required Password")
	}
	if u.Username == "" {
		return errors.New("Required Username")
	}
	if u.FirstName == "" {
		return errors.New("Required FirstName")
	}
	if u.LastName == "" {
		return errors.New("Required LastName")
	}
	if u.Email == "" {
		return errors.New("Required Email")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Invalid Email")
	}

	return nil
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
		return &domain.User{}, errors.New("User Not Found")
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

	if db.Error != nil {
		return db.Error
	}
	return nil
}
