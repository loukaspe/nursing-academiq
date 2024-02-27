package domain

import "time"

type QuizSessionByStudent struct {
	Quiz              *Quiz
	Student           *Student
	Date              time.Time
	DurationInSeconds int
	Score             float32
	MaxScore          int
}
