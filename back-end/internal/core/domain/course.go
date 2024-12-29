package domain

type Course struct {
	ID                uint32
	Title             string
	Description       string
	Chapters          []Chapter
	Quizzes           []Quiz
	Tutor             *Tutor
	NumberOfQuestions int
}
