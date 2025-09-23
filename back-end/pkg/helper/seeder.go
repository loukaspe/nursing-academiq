package helper

import (
	"log"

	"github.com/loukaspe/nursing-academiq/internal/repositories"

	"gorm.io/gorm"
)

var users = []repositories.User{
	repositories.User{
		Username:  "sarah.johnson",
		Password:  "password123",
		FirstName: "Sarah",
		LastName:  "Johnson",
		Email:     "sarah.johnson@nursing.edu",
	},
	repositories.User{
		Username:  "michael.chen",
		Password:  "password123",
		FirstName: "Michael",
		LastName:  "Chen",
		Email:     "michael.chen@nursing.edu",
	},
	repositories.User{
		Username:  "emma.rodriguez",
		Password:  "password123",
		FirstName: "Emma",
		LastName:  "Rodriguez",
		Email:     "emma.rodriguez@nursing.edu",
	},
	repositories.User{
		Username:  "david.kim",
		Password:  "password123",
		FirstName: "David",
		LastName:  "Kim",
		Email:     "david.kim@nursing.edu",
	},
	repositories.User{
		Username:  "lisa.anderson",
		Password:  "password123",
		FirstName: "Lisa",
		LastName:  "Anderson",
		Email:     "lisa.anderson@nursing.edu",
	},
	repositories.User{
		Username:  "dr.korelli",
		Password:  "korelli123",
		FirstName: "Dr. Maria",
		LastName:  "Korelli",
		Email:     "maria.korelli@nursing.edu",
	},
}

var tutor1 = repositories.Tutor{
	UserID:       1,
	AcademicRank: "Assistant Professor",
}

var tutor2 = repositories.Tutor{
	UserID:       2,
	AcademicRank: "Associate Professor",
}

var course1 = repositories.Course{
	Title:       "Fundamentals of Nursing",
	Description: "Introduction to basic nursing principles, patient care, and professional practice",
	TutorID:     2,
}

var chapter1 = repositories.Chapter{
	Title:       "Introduction to Nursing Practice",
	Description: "Overview of nursing as a profession and basic care principles",
	CourseID:    1,
}

var chapter2 = repositories.Chapter{
	Title:       "Patient Assessment and Vital Signs",
	Description: "Learning to assess patients and monitor vital signs effectively",
	CourseID:    1,
}

var chapter3 = repositories.Chapter{
	Title:       "Infection Control and Safety",
	Description: "Understanding infection prevention and maintaining patient safety",
	CourseID:    1,
}

var course2 = repositories.Course{
	Title:       "Medical-Surgical Nursing",
	Description: "Advanced nursing care for medical and surgical patients",
	TutorID:     2,
}

var course3 = repositories.Course{
	Title:       "Pharmacology for Nurses",
	Description: "Understanding medications, drug interactions, and safe administration",
	TutorID:     2,
}

var course4 = repositories.Course{
	Title:       "Pediatric Nursing",
	Description: "Specialized care for infants, children, and adolescents",
	TutorID:     2,
}

var course5 = repositories.Course{
	Title:       "Mental Health Nursing",
	Description: "Caring for patients with mental health conditions and psychiatric disorders",
	TutorID:     2,
}

var question1 = repositories.Question{
	Text:                   "What is the normal range for adult blood pressure?",
	Explanation:            "Normal adult blood pressure is typically defined as systolic pressure less than 120 mmHg and diastolic pressure less than 80 mmHg",
	ChapterID:              1,
	CourseID:               1,
	Source:                 "American Heart Association Guidelines",
	MultipleCorrectAnswers: false,
	NumberOfAnswers:        1,
}

var question2 = repositories.Question{
	Text:                   "Which of the following is the most effective method for preventing healthcare-associated infections?",
	Explanation:            "Hand hygiene is the single most effective method for preventing healthcare-associated infections and should be performed before and after patient contact",
	ChapterID:              2,
	CourseID:               1,
	Source:                 "CDC Infection Control Guidelines",
	MultipleCorrectAnswers: false,
	NumberOfAnswers:        1,
}

var question3 = repositories.Question{
	Text:                   "What is the primary purpose of the nursing process?",
	Explanation:            "The nursing process is a systematic method used by nurses to provide individualized, patient-centered care through assessment, diagnosis, planning, implementation, and evaluation",
	ChapterID:              1,
	CourseID:               1,
	Source:                 "Nursing Fundamentals Textbook",
	MultipleCorrectAnswers: false,
	NumberOfAnswers:        1,
}

var answer1 = repositories.Answer{
	Text:       "140/90 mmHg",
	QuestionID: 1,
	IsCorrect:  false,
	//TimesGiven: 0,
}

var answer2 = repositories.Answer{
	Text:       "Less than 120/80 mmHg",
	QuestionID: 1,
	IsCorrect:  true,
	//TimesGiven: 0,
}

var answer3 = repositories.Answer{
	Text:       "Wearing gloves only",
	QuestionID: 2,
	IsCorrect:  false,
	//TimesGiven: 0,
}

var answer4 = repositories.Answer{
	Text:       "Proper hand hygiene",
	QuestionID: 2,
	IsCorrect:  true,
	//TimesGiven: 0,
}

var answer5 = repositories.Answer{
	Text:       "To document patient information",
	QuestionID: 3,
	IsCorrect:  false,
	//TimesGiven: 0,
}

var answer6 = repositories.Answer{
	Text:       "To provide systematic, patient-centered care",
	QuestionID: 3,
	IsCorrect:  true,
	//TimesGiven: 0,
}

var quiz1 = repositories.Quiz{
	Title:       "Nursing Fundamentals Assessment",
	Description: "Comprehensive quiz covering basic nursing principles and patient care",
	CourseID:    1,
	Visibility:  true,
	ShowSubset:  false,
	SubsetSize:  1,
	////NumberOfSessions: 2,
	ScoreSum: 3,
	MaxScore: 4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz11 = repositories.Quiz{
	Title:       "Patient Assessment Quiz",
	Description: "Test your knowledge of patient assessment techniques and vital signs",
	CourseID:    1,
	Visibility:  true,
	ShowSubset:  false,
	SubsetSize:  1,
	//NumberOfSessions: 2,
	ScoreSum: 3,
	MaxScore: 4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz12 = repositories.Quiz{
	Title:       "Nursing Process Review",
	Description: "Evaluate your understanding of the nursing process and care planning",
	CourseID:    1,
	Visibility:  true,
	ShowSubset:  false,
	SubsetSize:  1,
	//NumberOfSessions: 2,
	ScoreSum: 3,
	MaxScore: 4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz13 = repositories.Quiz{
	Title:       "Professional Practice Quiz",
	Description: "Assessment of nursing professional standards and ethical practice",
	CourseID:    1,
	Visibility:  true,
	ShowSubset:  false,
	SubsetSize:  1,
	//NumberOfSessions: 2,
	ScoreSum: 3,
	MaxScore: 4,
	Questions: []*repositories.Question{
		&question1, &question3,
	},
}

var quiz21 = repositories.Quiz{
	Title:       "Infection Control Assessment",
	Description: "Test your knowledge of infection prevention and control measures",
	CourseID:    2,
	Visibility:  true,
	ShowSubset:  false,
	SubsetSize:  1,
	//NumberOfSessions: 2,
	ScoreSum: 3,
	MaxScore: 4,
	Questions: []*repositories.Question{
		&question2, &question3,
	},
}

var quiz22 = repositories.Quiz{
	Title:       "Safety Protocols Quiz",
	Description: "Assessment of patient safety protocols and healthcare standards",
	CourseID:    2,
	Visibility:  true,
	ShowSubset:  false,
	SubsetSize:  1,
	//NumberOfSessions: 2,
	ScoreSum: 3,
	MaxScore: 4,
	Questions: []*repositories.Question{
		&question2, &question3,
	},
}

var quiz23 = repositories.Quiz{
	Title:       "Medical-Surgical Nursing Review",
	Description: "Comprehensive review of medical-surgical nursing principles and practices",
	CourseID:    2,
	Visibility:  true,
	ShowSubset:  false,
	SubsetSize:  1,
	//NumberOfSessions: 2,
	ScoreSum: 3,
	MaxScore: 4,
	Questions: []*repositories.Question{
		&question2, &question3,
	},
}

func LoadFakeData(db *gorm.DB) {
	CreateUsers(db)
	CreateAdminUser(db)
	CreateAdminTutor(db)
	CreateTutors(db)
	CreateCourses(db)
	CreateChapters(db)
	CreateQuestions(db)
	CreateAnswers(db)
	CreateQuizzes(db)
}

func PrepareDB(db *gorm.DB) {
	DropTables(db)
	CreateAdminUser(db)
	CreateAdminTutor(db)
	CreateCoursesTable(db)
	CreateChaptersTable(db)
	CreateQuestionsTable(db)
	CreateAnswersTable(db)
	CreateQuizzesTable(db)
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

func CreateAdminUser(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.User{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	admin := repositories.User{
		Username:  "loukas",
		Password:  "loukastest",
		FirstName: "Loukas",
		LastName:  "Peteinaris",
		Email:     "loukas@email.com",
	}

	err = admin.BeforeSave()
	if err != nil {
		log.Fatalf("cannot hash in seed users table: %v", err)
	}
	err = db.Debug().Model(&repositories.User{}).Create(&admin).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
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

func CreateAdminTutor(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Tutor{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	var admin = repositories.Tutor{
		UserID:       1,
		AcademicRank: "Admin",
	}

	err = db.Debug().Model(&repositories.Tutor{}).Create(&admin).Error
	if err != nil {
		log.Fatalf("cannot seed tutors table: %v", err)
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
	err = db.Debug().Model(&repositories.Course{}).Create(&course3).Error
	if err != nil {
		log.Fatalf("cannot seed courses table: %v", err)
	}
	err = db.Debug().Model(&repositories.Course{}).Create(&course4).Error
	if err != nil {
		log.Fatalf("cannot seed courses table: %v", err)
	}
	err = db.Debug().Model(&repositories.Course{}).Create(&course5).Error
	if err != nil {
		log.Fatalf("cannot seed courses table: %v", err)
	}
}

func CreateCoursesTable(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Course{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
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

func CreateChaptersTable(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Chapter{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
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

func CreateQuestionsTable(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Question{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
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

func CreateAnswersTable(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Answer{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
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
	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz21).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz22).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
	err = db.Debug().Model(&repositories.Quiz{}).Create(&quiz23).Error
	if err != nil {
		log.Fatalf("cannot seed quizs table: %v", err)
	}
}

func CreateQuizzesTable(db *gorm.DB) {
	var err error
	err = db.Debug().AutoMigrate(&repositories.Quiz{})
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
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
	err := db.Migrator().DropTable("quiz_has_question")
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
}
