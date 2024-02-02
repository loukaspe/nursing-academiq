package domain

import "time"

type User struct {
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	BirthDate   time.Time
	PhoneNumber string
	Photo       string
	// There are either "student"/"tutor" and their ID
	UserType   string
	SpecificID uint
}
