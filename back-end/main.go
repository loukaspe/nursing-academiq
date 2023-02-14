package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/loukaspe/nursing-academiq/models"
	"github.com/loukaspe/nursing-academiq/pkg/helper"
	"github.com/loukaspe/nursing-academiq/pkg/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	router := mux.NewRouter()

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Europe/Athens",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database: ", err)
	}

	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("database migration error", err)
	}
	helper.LoadFakeData(db)

	server := server.NewServer(db, router)

	server.Run(":8080")
}
