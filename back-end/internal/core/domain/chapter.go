package domain

type Chapter struct {
	Title       string
	Description string
	Course      *Course
}
