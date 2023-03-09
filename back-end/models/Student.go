package models

import (
	"errors"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	UserId             string
	RegistrationNumber string `gorm:"not null;"`
	User               User   `gorm:"foreignKey:UserId"`
}

func (u *Student) Validate() error {
	if u.RegistrationNumber == "" {
		return errors.New("Required RegistrationNumber")
	}

	return nil
}

func (u *Student) SaveStudent(db *gorm.DB) (*Student, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Student{}, err
	}
	return u, nil
}

func (u *Student) FindAllStudents(db *gorm.DB) (*[]Student, error) {
	var err error
	var Students []Student
	err = db.Debug().Model(&Student{}).Limit(100).Find(&Students).Error
	if err != nil {
		return &[]Student{}, err
	}
	return &Students, err
}

func (u *Student) FindStudentByID(db *gorm.DB, uid uint32) (*Student, error) {
	var err error
	err = db.Debug().Model(Student{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Student{}, err
	}
	if err == gorm.ErrRecordNotFound {
		return &Student{}, errors.New("Student Not Found")
	}
	return u, err
}

func (u *Student) UpdateAStudent(db *gorm.DB, uid uint32) (*Student, error) {
	db = db.Debug().
		Model(&Student{}).
		Where("id = ?", uid).
		Take(&Student{}).
		UpdateColumns(
			map[string]interface{}{
				"registration_number": u.RegistrationNumber,
			},
		)
	if db.Error != nil {
		return &Student{}, db.Error
	}
	// This is the display the updated Student
	err := db.Debug().Model(&Student{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Student{}, err
	}
	return u, nil
}

func (u *Student) DeleteAStudent(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().
		Model(&Student{}).
		Where("id = ?", uid).
		Take(&Student{}).
		Delete(&Student{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
