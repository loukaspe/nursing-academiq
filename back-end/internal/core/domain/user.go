package domain

type User struct {
	ID         uint
	Username   string
	Password   string
	FirstName  string
	LastName   string
	Email      string
	UserType   string
	SpecificID uint
}
