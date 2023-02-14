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
}

func LoadFakeData(db *gorm.DB) {
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
