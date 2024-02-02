package domain

type Student struct {
	User
	RegistrationNumber string
	Courses            []Course
}
