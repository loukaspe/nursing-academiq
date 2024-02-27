package repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	RegistrationNumber string `gorm:"not null;"`
	UserID             uint   `gorm:"not null"`
	User               User
	QuizSessions       []QuizSession
	Courses            []Course `gorm:"many2many:student_takes_course;"`
}

func (s *Student) validate() error {
	if s.RegistrationNumber == "" {
		return errors.New("Required RegistrationNumber")
	}

	return nil
}
