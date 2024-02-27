package helper

import (
	"github.com/loukaspe/nursing-academiq/internal/repositories"
	"log"
	"time"

	"gorm.io/gorm"
)

var users = []repositories.User{
	repositories.User{
		Username:    "user1",
		Password:    "password1",
		FirstName:   "firstName1",
		LastName:    "lastName1",
		Email:       "email1@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "1111111111",
		Photo:       "photo1",
	},
	repositories.User{
		Username:    "user2",
		Password:    "password2",
		FirstName:   "firstName2",
		LastName:    "lastName2",
		Email:       "email2@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "2222222222",
		Photo:       "photo2",
	},
	repositories.User{
		Username:    "user3",
		Password:    "password3",
		FirstName:   "firstName3",
		LastName:    "lastName3",
		Email:       "email3@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "33333333333",
		Photo:       "photo3",
	},
	repositories.User{
		Username:    "user4",
		Password:    "password4",
		FirstName:   "firstName4",
		LastName:    "lastName4",
		Email:       "email4@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "4444444444",
		Photo:       "photo4",
	},
	repositories.User{
		Username:    "user5",
		Password:    "password5",
		FirstName:   "firstName5",
		LastName:    "lastName5",
		Email:       "email5@email.com",
		BirthDate:   time.Time{},
		PhoneNumber: "5555555555",
		Photo:       "photo5",
	},
}

var tutor1 = repositories.Tutor{
	UserID:       1,
	AcademicRank: "professoras1",
}

var tutor2 = repositories.Tutor{
	UserID:       5,
	AcademicRank: "professoras2",
}

var student1 = repositories.Student{
	UserID:             2,
	RegistrationNumber: "§1",
	Courses: []repositories.Course{
		course1,
	},
}

var student2 = repositories.Student{
	UserID:             3,
	RegistrationNumber: "123",
	Courses: []repositories.Course{
		course3,
	},
}

var student3 = repositories.Student{
	UserID:             4,
	RegistrationNumber: "1234",
	Courses: []repositories.Course{
		course2,
	},
}

var course1 = repositories.Course{
	Title:       "math gen",
	Description: "mathimatika genikhs",
	TutorID:     1,
	//Students: []*repositories.Students{
	//	&student1,
	//},
}

var chapter1 = repositories.Chapter{
	Title:       "math gen kef 1",
	Description: "math gen kefalaio uno",
	CourseID:    1,
}

var chapter2 = repositories.Chapter{
	Title:       "math gen kef 2",
	Description: "math gen kefalaio dos",
	CourseID:    1,
}

var chapter3 = repositories.Chapter{
	Title:       "math gen kef 3",
	Description: "math gen kefalaio tres",
	CourseID:    1,
}

var course2 = repositories.Course{
	Title:       "math kat",
	Description: "mathimatika kateythynshs",
	TutorID:     2,
	//Students: []*repositories.Students{
	//	&student2, &student3,
	//},
}

var course3 = repositories.Course{
	Title:       "math 3",
	Description: "mathimatika 3",
	TutorID:     2,
	//Students: []*repositories.Students{
	//	&student2, &student3,
	//},
}

var question1 = repositories.Question{
	Text:                   "ti sko",
	Explanation:            "ti skotwse o ai giorgis",
	ChapterID:              1,
	Source:                 "wikipedia1",
	MultipleCorrectAnswers: true,
	NumberOfAnswers:        1,
	TimesAnswered:          2,
	TimesAnsweredCorrectly: 3,
	//Quizzes: []repositories.Quiz{
	//	quiz1,
	//},
}

var question2 = repositories.Question{
	Text:                   "me ti",
	Explanation:            "me ti ftiaxnoyme pastitsio",
	ChapterID:              2,
	Source:                 "google",
	MultipleCorrectAnswers: true,
	NumberOfAnswers:        1,
	TimesAnswered:          2,
	TimesAnsweredCorrectly: 3,
	//Quizzes: []repositories.Quiz{
	//	quiz2,
	//},
}

var question3 = repositories.Question{
	Text:                   "with what",
	Explanation:            "with what we kill the cat",
	ChapterID:              1,
	Source:                 "chatgpt",
	MultipleCorrectAnswers: true,
	NumberOfAnswers:        1,
	TimesAnswered:          2,
	TimesAnsweredCorrectly: 3,
	//Quizzes: []repositories.Quiz{
	//	quiz1, quiz2,
	//},
}

var answer1 = repositories.Answer{
	Text:       "apanthsh1",
	QuestionID: 1,
	IsCorrect:  false,
	TimesGiven: 0,
}

var answer2 = repositories.Answer{
	Text:       "apanthsh2",
	QuestionID: 1,
	IsCorrect:  true,
	TimesGiven: 0,
}

var answer3 = repositories.Answer{
	Text:       "apanthsh3",
	QuestionID: 2,
	IsCorrect:  false,
	TimesGiven: 0,
}

var answer4 = repositories.Answer{
	Text:       "apanthsh4",
	QuestionID: 2,
	IsCorrect:  true,
	TimesGiven: 0,
}

var answer5 = repositories.Answer{
	Text:       "apanthsh5",
	QuestionID: 3,
	IsCorrect:  false,
	TimesGiven: 0,
}

var answer6 = repositories.Answer{
	Text:       "apanthsh6",
	QuestionID: 3,
	IsCorrect:  true,
	TimesGiven: 0,
}

var quiz1 = repositories.Quiz{
	Title:            "first quiz",
	Description:      "the first quiz of the chapter",
	CourseID:         1,
	Visibility:       true,
	ShowSubset:       false,
	SubsetSize:       1,
	NumberOfSessions: 2,
	ScoreSum:         3,
	MaxScore:         4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz11 = repositories.Quiz{
	Title:            "first and one quiz",
	Description:      "the first quiz of the chapter",
	CourseID:         1,
	Visibility:       true,
	ShowSubset:       false,
	SubsetSize:       1,
	NumberOfSessions: 2,
	ScoreSum:         3,
	MaxScore:         4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz12 = repositories.Quiz{
	Title:            "first and two quiz",
	Description:      "the first quiz of the chapter",
	CourseID:         1,
	Visibility:       true,
	ShowSubset:       false,
	SubsetSize:       1,
	NumberOfSessions: 2,
	ScoreSum:         3,
	MaxScore:         4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz13 = repositories.Quiz{
	Title:            "first and three quiz",
	Description:      "the first quiz of the chapter",
	CourseID:         1,
	Visibility:       true,
	ShowSubset:       false,
	SubsetSize:       1,
	NumberOfSessions: 2,
	ScoreSum:         3,
	MaxScore:         4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz2 = repositories.Quiz{
	Title:            "second quiz",
	Description:      "the second quiz of the chapter",
	CourseID:         2,
	Visibility:       true,
	ShowSubset:       false,
	SubsetSize:       1,
	NumberOfSessions: 2,
	ScoreSum:         3,
	MaxScore:         4,
	Questions: []*repositories.Question{
		&question2, &question3,
	},
}

var quizSession1 = repositories.QuizSession{
	DateTime:          time.Now(),
	DurationInSeconds: 1,
	Score:             2,
	MaxScore:          3,
	QuestionSessions: []repositories.QuestionSession{
		questionSession1, questionSession2,
	},
	AnswerSessions: []repositories.AnswerSession{
		answerSession1, answerSession2,
	},
	StudentID: 1,
	QuizID:    1,
}

var quizSessionID1 uint = 1

var quizSession2 = repositories.QuizSession{
	DateTime:          time.Now(),
	DurationInSeconds: 11,
	Score:             22,
	MaxScore:          33,
	QuestionSessions:  []repositories.QuestionSession{},
	AnswerSessions:    []repositories.AnswerSession{},
	StudentID:         2,
	QuizID:            2,
}

var questionSession1 = repositories.QuestionSession{
	QuizSessionID: &quizSessionID1,
	Question:      question1,
	Position:      1,
}

var questionSession2 = repositories.QuestionSession{
	QuizSessionID: &quizSessionID1,
	Question:      question3,
	Position:      2,
}

var answerSession1 = repositories.AnswerSession{
	QuizSessionID: &quizSessionID1,
	Answer:        answer1,
}

var answerSession2 = repositories.AnswerSession{
	QuizSessionID: &quizSessionID1,
	Answer:        answer2,
}

func LoadFakeData(db *gorm.DB) {
	DropTables(db)
	CreateUsers(db)
	CreateTutors(db)
	CreateStudents(db)
	//CreateCourses(db)
	CreateChapters(db)
	CreateQuestions(db)
	CreateAnswers(db)
	CreateQuizzes(db)
	CreateQuizSessions(db)
	CreateQuestionSession(db)
	CreateAnswerSession(db)
}

func CreateUsers(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = users[i].BeforeSave()
		if err != nil {
			log.Fatalf("cannot hash in seed users table: %v", err)
		}
		err = db.Debug().Model(&repositories.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

func CreateTutors(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Tutor{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.Tutor{}).Create(&tutor1).Error
	if err != nil {
		log.Fatalf("cannot seed tutors table: %v", err)
	}
	err = db.Debug().Model(&repositories.Tutor{}).Create(&tutor2).Error
	if err != nil {
		log.Fatalf("cannot seed tutors table: %v", err)
	}
}

func CreateStudents(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Student{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.Student{}).Create(&student1).Error
	if err != nil {
		log.Fatalf("cannot seed students table: %v", err)
	}
	err = db.Debug().Model(&repositories.Student{}).Create(&student2).Error
	if err != nil {
		log.Fatalf("cannot seed students table: %v", err)
	}
	err = db.Debug().Model(&repositories.Student{}).Create(&student3).Error
	if err != nil {
		log.Fatalf("cannot seed students table: %v", err)
	}
}

func CreateCourses(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Course{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.Course{}).Create(&course1).Error
	if err != nil {
		log.Fatalf("cannot seed courses table: %v", err)
	}
	err = db.Debug().Model(&repositories.Course{}).Create(&course2).Error
	if err != nil {
		log.Fatalf("cannot seed courses table: %v", err)
	}
}

func CreateChapters(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Chapter{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.Chapter{}).Create(&chapter1).Error
	if err != nil {
		log.Fatalf("cannot seed chapters table: %v", err)
	}
	err = db.Debug().Model(&repositories.Chapter{}).Create(&chapter2).Error
	if err != nil {
		log.Fatalf("cannot seed chapters table: %v", err)
	}
	err = db.Debug().Model(&repositories.Chapter{}).Create(&chapter3).Error
	if err != nil {
		log.Fatalf("cannot seed chapters table: %v", err)
	}
}

func CreateQuestions(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Question{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.Question{}).Create(&question1).Error
	if err != nil {
		log.Fatalf("cannot seed questions table: %v", err)
	}
	err = db.Debug().Model(&repositories.Question{}).Create(&question2).Error
	if err != nil {
		log.Fatalf("cannot seed questions table: %v", err)
	}
	err = db.Debug().Model(&repositories.Question{}).Create(&question3).Error
	if err != nil {
		log.Fatalf("cannot seed questions table: %v", err)
	}
}

func CreateAnswers(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Answer{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.Answer{}).Create(&answer1).Error
	if err != nil {
		log.Fatalf("cannot seed answers table: %v", err)
	}
	err = db.Debug().Model(&repositories.Answer{}).Create(&answer2).Error
	if err != nil {
		log.Fatalf("cannot seed answers table: %v", err)
	}
	err = db.Debug().Model(&repositories.Answer{}).Create(&answer3).Error
	if err != nil {
		log.Fatalf("cannot seed answers table: %v", err)
	}
	err = db.Debug().Model(&repositories.Answer{}).Create(&answer4).Error
	if err != nil {
		log.Fatalf("cannot seed answers table: %v", err)
	}
	err = db.Debug().Model(&repositories.Answer{}).Create(&answer5).Error
	if err != nil {
		log.Fatalf("cannot seed answers table: %v", err)
	}
	err = db.Debug().Model(&repositories.Answer{}).Create(&answer6).Error
	if err != nil {
		log.Fatalf("cannot seed answers table: %v", err)
	}
}

func CreateQuizzes(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Quiz{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz1).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz11).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz12).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz13).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz2).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
}

func CreateQuizSessions(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.QuizSession{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.QuizSession{}).Create(&quizSession1).Error
	if err != nil {
		log.Fatalf("cannot seed quiz sessions table: %v", err)
	}
	err = db.Debug().Model(&repositories.QuizSession{}).Create(&quizSession2).Error
	if err != nil {
		log.Fatalf("cannot seed quiz sessions table: %v", err)
	}
}

func CreateQuestionSession(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.QuestionSession{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.QuestionSession{}).Create(&questionSession1).Error
	if err != nil {
		log.Fatalf("cannot seed question sessions table: %v", err)
	}
	err = db.Debug().Model(&repositories.QuestionSession{}).Create(&questionSession2).Error
	if err != nil {
		log.Fatalf("cannot seed question sessions table: %v", err)
	}
}

func CreateAnswerSession(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.AnswerSession{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&repositories.AnswerSession{}).Create(&answerSession1).Error
	if err != nil {
		log.Fatalf("cannot seed answer sessions table: %v", err)
	}
	err = db.Debug().Model(&repositories.AnswerSession{}).Create(&answerSession2).Error
	if err != nil {
		log.Fatalf("cannot seed answer sessions table: %v", err)
	}
}

func DropTables(db *gorm.DB) {
	if db.Migrator().HasTable(&repositories.User{}) {
		err := db.Migrator().DropTable("users")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.Tutor{}) {
		err := db.Migrator().DropTable("tutors")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.Student{}) {
		err := db.Migrator().DropTable("students")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.Course{}) {
		err := db.Migrator().DropTable("courses")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.Chapter{}) {
		err := db.Migrator().DropTable("chapters")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.Question{}) {
		err := db.Migrator().DropTable("questions")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.Answer{}) {
		err := db.Migrator().DropTable("answers")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.Quiz{}) {
		err := db.Migrator().DropTable("quizzes")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.QuizSession{}) {
		err := db.Migrator().DropTable("quiz_sessions")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.QuestionSession{}) {
		err := db.Migrator().DropTable("question_sessions")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	if db.Migrator().HasTable(&repositories.AnswerSession{}) {
		err := db.Migrator().DropTable("answer_sessions")
		if err != nil {
			log.Fatalf("cannot drop table: %v", err)
		}
	}
	err := db.Migrator().DropTable("quiz_has_question")
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Migrator().DropTable("student_takes_course")
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
}
