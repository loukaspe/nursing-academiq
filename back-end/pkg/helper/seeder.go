package helper

import (
	"github.com/loukaspe/nursing-academiq/internal/repositories"
	"log"
	"time"

	"gorm.io/gorm"
)

var users = []repositories.User{
	repositories.User{
		Username:    "user1",
		Password:    "password1",
		FirstName:   "firstName1",
		LastName:    "lastName1",
		Email:       "email1@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "1111111111",
		Photo:       "photo1",
	},
	repositories.User{
		Username:    "user2",
		Password:    "password2",
		FirstName:   "firstName2",
		LastName:    "lastName2",
		Email:       "email2@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "2222222222",
		Photo:       "photo2",
	},
	repositories.User{
		Username:    "user3",
		Password:    "password3",
		FirstName:   "firstName3",
		LastName:    "lastName3",
		Email:       "email3@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "33333333333",
		Photo:       "photo3",
	},
	repositories.User{
		Username:    "user4",
		Password:    "password4",
		FirstName:   "firstName4",
		LastName:    "lastName4",
		Email:       "email4@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "4444444444",
		Photo:       "photo4",
	},
	repositories.User{
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

//var tutor = repositories.Tutor{
//	UserId:       "1",
//	AcademicRank: "professoras",
//}
//
//var student = repositories.Student{
//	UserId:             "2",
//	RegistrationNumber: "1234556",
//}

func LoadFakeData(db *gorm.DB) {
	CreateUsers(db)
	//CreateTutors(db)
	//CreateStudents(db)
}

func CreateUsers(db *gorm.DB) {
	var err error
	if db.Migrator().HasTable(&repositories.User{}) {
		err := db.Migrator().DropTable("users")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}

	err = db.Debug().AutoMigrate(&repositories.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&repositories.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

//func CreateTutors(db *gorm.DB) {
//	var err error
//	if db.Migrator().HasTable(&repositories.Tutor{}) {
//		err := db.Migrator().DropTable("tutors")
//		if err != nil {
//			log.Fatalf("cannot drop table: %v", err)
//		}
//	}
//
//	err = db.Debug().AutoMigrate(&repositories.Tutor{})
//	if err != nil {
//		log.Fatalf("cannot migrate table: %v", err)
//	}
//
//	err = db.Debug().Model(&repositories.Tutor{}).Create(&tutor).Error
//	if err != nil {
//		log.Fatalf("cannot seed tutors table: %v", err)
//	}
//}
//
//func CreateStudents(db *gorm.DB) {
//	var err error
//	if db.Migrator().HasTable(&repositories.Student{}) {
//		err := db.Migrator().DropTable("students")
//		if err != nil {
//			log.Fatalf("cannot drop table: %v", err)
//		}
//	}
//
//	err = db.Debug().AutoMigrate(&repositories.Student{})
//	if err != nil {
//		log.Fatalf("cannot migrate table: %v", err)
//	}
//
//	err = db.Debug().Model(&repositories.Student{}).Create(&student).Error
//	if err != nil {
//		log.Fatalf("cannot seed students table: %v", err)
//	}
//}
