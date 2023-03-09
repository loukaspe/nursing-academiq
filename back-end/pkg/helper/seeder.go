package helper

import (
	"log"
	"time"

	"github.com/loukaspe/nursing-academiq/models"
	"gorm.io/gorm"
)

var users = []models.User{
	models.User{
		Username:    "user1",
		Password:    "password1",
		FirstName:   "firstName1",
		LastName:    "lastName1",
		Email:       "email1@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "1111111111",
		Photo:       "photo1",
	},
	models.User{
		Username:    "user2",
		Password:    "password2",
		FirstName:   "firstName2",
		LastName:    "lastName2",
		Email:       "email2@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "2222222222",
		Photo:       "photo2",
	},
	models.User{
		Username:    "user3",
		Password:    "password3",
		FirstName:   "firstName3",
		LastName:    "lastName3",
		Email:       "email3@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "33333333333",
		Photo:       "photo3",
	},
	models.User{
		Username:    "user4",
		Password:    "password4",
		FirstName:   "firstName4",
		LastName:    "lastName4",
		Email:       "email4@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "4444444444",
		Photo:       "photo4",
	},
	models.User{
		Username:    "user5",
		Password:    "password5",
		FirstName:   "firstName5",
		LastName:    "lastName5",
		Email:       "email5@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "5555555555",
		Photo:       "photo5",
	},
}

var tutor = models.Tutor{
	UserId:       "1",
	AcademicRank: "professoras",
}

var student = models.Student{
	UserId:             "2",
	RegistrationNumber: "1234556",
}

func LoadFakeData(db *gorm.DB) {
	CreateUsers(db)
	CreateTutors(db)
	CreateStudents(db)
}

func CreateUsers(db *gorm.DB) {
	var err error
	if db.Migrator().HasTable(&models.User{}) {
		err := db.Migrator().DropTable("users")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}

	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

func CreateTutors(db *gorm.DB) {
	var err error
	if db.Migrator().HasTable(&models.Tutor{}) {
		err := db.Migrator().DropTable("tutors")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}

	err = db.Debug().AutoMigrate(&models.Tutor{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Tutor{}).Create(&tutor).Error
	if err != nil {
		log.Fatalf("cannot seed tutors table: %v", err)
	}
}

func CreateStudents(db *gorm.DB) {
	var err error
	if db.Migrator().HasTable(&models.Student{}) {
		err := db.Migrator().DropTable("students")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}

	err = db.Debug().AutoMigrate(&models.Student{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Student{}).Create(&student).Error
	if err != nil {
		log.Fatalf("cannot seed students table: %v", err)
	}
}
