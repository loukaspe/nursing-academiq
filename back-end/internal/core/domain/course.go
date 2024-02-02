package domain

type Course struct {
	Title       string
	Description string
	Tutor       *Tutor
	Students    []Student
}
