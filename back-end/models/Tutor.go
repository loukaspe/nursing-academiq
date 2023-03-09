package models

import (
	"errors"
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	UserId       string
	AcademicRank string `gorm:"not null;"`
	User         User   `gorm:"foreignKey:UserId"`
}

func (u *Tutor) Validate() error {
	if u.AcademicRank == "" {
		return errors.New("Required AcademicRank")
	}

	return nil
}

func (u *Tutor) SaveTutor(db *gorm.DB) (*Tutor, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Tutor{}, err
	}
	return u, nil
}

func (u *Tutor) FindAllTutors(db *gorm.DB) (*[]Tutor, error) {
	var err error
	var Tutors []Tutor
	err = db.Debug().Model(&Tutor{}).Limit(100).Find(&Tutors).Error
	if err != nil {
		return &[]Tutor{}, err
	}
	return &Tutors, err
}

func (u *Tutor) FindTutorByID(db *gorm.DB, uid uint32) (*Tutor, error) {
	var err error
	err = db.Debug().Model(Tutor{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Tutor{}, err
	}
	if err == gorm.ErrRecordNotFound {
		return &Tutor{}, errors.New("Tutor Not Found")
	}
	return u, err
}

func (u *Tutor) UpdateATutor(db *gorm.DB, uid uint32) (*Tutor, error) {
	db = db.Debug().
		Model(&Tutor{}).
		Where("id = ?", uid).
		Take(&Tutor{}).
		UpdateColumns(
			map[string]interface{}{
				"academic_rank": u.AcademicRank,
			},
		)
	if db.Error != nil {
		return &Tutor{}, db.Error
	}
	// This is the display the updated Tutor
	err := db.Debug().Model(&Tutor{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Tutor{}, err
	}
	return u, nil
}

func (u *Tutor) DeleteATutor(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().
		Model(&Tutor{}).
		Where("id = ?", uid).
		Take(&Tutor{}).
		Delete(&Tutor{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
