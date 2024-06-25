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

type Chapter struct {
	gorm.Model
	Title       string `gorm:"not null;"`
	Description string
	CourseID    uint
	Course      Course
	Questions   []Question
}

type Quiz struct {
	gorm.Model
	Title       string `gorm:"not null;"`
	Description string
	//Course           Course
	CourseID         uint
	CourseName       string
	Visibility       bool `gorm:"not null;"`
	ShowSubset       bool `gorm:"not null;"`
	SubsetSize       int
	NumberOfSessions int         `gorm:"not null;"`
	ScoreSum         float32     `gorm:"not null;"`
	MaxScore         int         `gorm:"not null;"`
	Questions        []*Question `gorm:"many2many:quiz_has_question;"`
}

type QuizSession struct {
	gorm.Model
	//QuizID            uint
	Quiz              Quiz
	QuizID            uint
	StudentID         uint
	DateTime          time.Time `gorm:"not null;"`
	DurationInSeconds int       `gorm:"not null;"`
	Score             float32   `gorm:"not null;"`
	MaxScore          int       `gorm:"not null;"`
	QuestionSessions  []QuestionSession
	AnswerSessions    []AnswerSession
}
