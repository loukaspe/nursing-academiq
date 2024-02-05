package repositories

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	Title       string `gorm:"not null;"`
	Description string
	TutorID     uint
	Tutor       Tutor
	Chapters    []Chapter
	Students    []Student `gorm:"many2many:student_takes_course;"`
	Quizs       []Quiz
}

type Quiz struct {
	gorm.Model
	Title       string `gorm:"not null;"`
	Description string
	//Course           Course
	CourseID         uint
	CourseName       string
	QuizSessionID    *uint
	Visibility       bool `gorm:"not null;"`
	ShowSubset       bool `gorm:"not null;"`
	SubsetSize       int
	NumberOfSessions int         `gorm:"not null;"`
	ScoreSum         float32     `gorm:"not null;"`
	MaxScore         int         `gorm:"not null;"`
	Questions        []*Question `gorm:"many2many:quiz_has_question;"`
}
