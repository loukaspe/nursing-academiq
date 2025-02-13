package domain

import "time"

type User struct {
	ID          uint
	Username    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	BirthDate   time.Time
	PhoneNumber string
	Photo       string
	UserType    string
	SpecificID  uint
}
