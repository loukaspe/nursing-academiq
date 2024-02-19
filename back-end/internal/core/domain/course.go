package domain

type Course struct {
	ID          uint32
	Title       string
	Description string
	Tutor       *Tutor
	Students    []Student
}
