package domain

import "time"

type Quiz struct {
	Title             string
	Description       string
	Course            *Course
	Visibility        bool
	ShowSubset        bool
	SubsetSize        int
	NumberOfSessions  int
	ScoreSum          float32
	MaxScore          int
	Questions         []Question
	NumberOfQuestions int
	CreatedAt         time.Time
}
