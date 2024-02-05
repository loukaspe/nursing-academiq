package repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	AcademicRank string `gorm:"not null;"`
	UserID       uint   `gorm:"not null"`
	User         User
	Courses      []Course
}

func (t *Tutor) validate() error {
	if t.AcademicRank == "" {
		return errors.New("Required AcademicRank")
	}

	return nil
}
