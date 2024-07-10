package repositories

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	Title       string `gorm:"not null;"`
	Description string
	TutorID     uint
	Tutor       Tutor
	Chapters    []Chapter
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
	CourseID   uint
	CourseName string
	Visibility bool `gorm:"not null;"`
	ShowSubset bool `gorm:"not null;"`
	SubsetSize int
	//NumberOfSessions int         `gorm:"not null;"`
	ScoreSum  float32     `gorm:"not null;"`
	MaxScore  int         `gorm:"not null;"`
	Questions []*Question `gorm:"many2many:quiz_has_question;"`
}

type Question struct {
	gorm.Model
	Text        string `gorm:"not null;"`
	Explanation string `gorm:"not null;"`
	//Chapter                *Chapter
	ChapterID              uint
	Source                 string `gorm:"not null;"`
	MultipleCorrectAnswers bool   `gorm:"not null;"`
	NumberOfAnswers        int    `gorm:"not null;"`
	//TimesAnswered          int     `gorm:"not null;"`
	//TimesAnsweredCorrectly float32 `gorm:"not null;"`
	Answers []Answer
	//QuestionSessionID      *uint
	Quizs []*Quiz `gorm:"many2many:quiz_has_question;"`
}

type Answer struct {
	gorm.Model
	Text string `gorm:"not null;"`
	//Question   *Question
	QuestionID uint
	//AnswerSessionID *uint
	IsCorrect bool `gorm:"not null;"`
	//TimesGiven      int  `gorm:"not null;"`
}
