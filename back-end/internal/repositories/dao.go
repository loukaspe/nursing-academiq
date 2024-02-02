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
