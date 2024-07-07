package repositories

import (
	"errors"
	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username  string `gorm:"not null;"`
	Password  string `gorm:"not null;"`
	FirstName string `gorm:"not null;"`
	LastName  string `gorm:"not null;"`
	Email     string `gorm:"not null;"`
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

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
