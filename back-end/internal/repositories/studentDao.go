package repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	RegistrationNumber string `gorm:"not null;"`
	UserId             uint   `gorm:"not null"`
	User               User   `gorm:"foreignKey:UserId"`
}

func (s *Student) validate() error {
	if s.RegistrationNumber == "" {
		return errors.New("Required RegistrationNumber")
	}

	return nil
}
