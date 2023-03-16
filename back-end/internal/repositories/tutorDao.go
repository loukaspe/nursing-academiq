package repositories

import (
	"errors"
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	TutorId      string
	AcademicRank string `gorm:"not null;"`
	User         User   `gorm:"foreignKey:UserId"`
}

func (t *Tutor) validate() error {
	if t.AcademicRank == "" {
		return errors.New("Required AcademicRank")
	}

	return nil
}
