package helper

import (
	"log"

	"github.com/loukaspe/nursing-academiq/internal/repositories"

	"gorm.io/gorm"
)

var users = []repositories.User{
	repositories.User{
		Username:  "maria.papadopoulou",
		Password:  "password123",
		FirstName: "Μαρία",
		LastName:  "Παπαδοπούλου",
		Email:     "maria.papadopoulou@nursing.edu",
	},
	repositories.User{
		Username:  "dimitris.kostas",
		Password:  "password123",
		FirstName: "Δημήτρης",
		LastName:  "Κώστας",
		Email:     "dimitris.kostas@nursing.edu",
	},
	repositories.User{
		Username:  "eleni.nikolaou",
		Password:  "password123",
		FirstName: "Ελένη",
		LastName:  "Νικολάου",
		Email:     "eleni.nikolaou@nursing.edu",
	},
	repositories.User{
		Username:  "kostas.petrou",
		Password:  "password123",
		FirstName: "Κώστας",
		LastName:  "Πέτρου",
		Email:     "kostas.petrou@nursing.edu",
	},
	repositories.User{
		Username:  "sofia.andrea",
		Password:  "password123",
		FirstName: "Σοφία",
		LastName:  "Ανδρέα",
		Email:     "sofia.andrea@nursing.edu",
	},
	repositories.User{
		Username:  "dkorelli",
		Password:  "korelli123",
		FirstName: "Αλεξάνδρα",
		LastName:  "Κορέλλη",
		Email:     "maria.korelli@nursing.edu",
	},
}

var tutor1 = repositories.Tutor{
	UserID:       1,
	AcademicRank: "Αναπληρωτής Καθηγητής",
}

var tutor2 = repositories.Tutor{
	UserID:       2,
	AcademicRank: "Επίκουρος Καθηγητής",
}

var course1 = repositories.Course{
	Title:       "Βασικές Αρχές Νοσηλευτικής",
	Description: "Εισαγωγή στις βασικές αρχές νοσηλευτικής, φροντίδα ασθενών και επαγγελματική πρακτική",
	TutorID:     1,
}

var chapter1 = repositories.Chapter{
	Title:       "Εισαγωγή στην Νοσηλευτική Πρακτική",
	Description: "Επισκόπηση της νοσηλευτικής ως επαγγέλματος και βασικών αρχών φροντίδας",
	CourseID:    1,
}

var chapter2 = repositories.Chapter{
	Title:       "Αξιολόγηση Ασθενούς και Σημεία Ζωής",
	Description: "Μάθηση αξιολόγησης ασθενών και παρακολούθησης σημείων ζωής",
	CourseID:    1,
}

var chapter3 = repositories.Chapter{
	Title:       "Έλεγχος Λοιμώξεων και Ασφάλεια",
	Description: "Κατανόηση πρόληψης λοιμώξεων και διατήρησης ασφάλειας ασθενών",
	CourseID:    1,
}

var course2 = repositories.Course{
	Title:       "Ιατρική-Χειρουργική Νοσηλευτική",
	Description: "Προχωρημένη νοσηλευτική φροντίδα για ιατρικούς και χειρουργικούς ασθενείς",
	TutorID:     1,
}

var course3 = repositories.Course{
	Title:       "Φαρμακολογία για Νοσηλευτές",
	Description: "Κατανόηση φαρμάκων, φαρμακευτικών αλληλεπιδράσεων και ασφαλούς χορήγησης",
	TutorID:     1,
}

var course4 = repositories.Course{
	Title:       "Παιδιατρική Νοσηλευτική",
	Description: "Εξειδικευμένη φροντίδα για βρέφη, παιδιά και εφήβους",
	TutorID:     1,
}

var course5 = repositories.Course{
	Title:       "Ψυχιατρική Νοσηλευτική",
	Description: "Φροντίδα ασθενών με ψυχικές διαταραχές και ψυχιατρικές παθήσεις",
	TutorID:     2,
}

var question1 = repositories.Question{
	Text:                   "Ποια είναι η φυσιολογική αξία της αρτηριακής πίεσης σε ενήλικες;",
	Explanation:            "Η φυσιολογική αρτηριακή πίεση ενήλικων ορίζεται ως συστολική πίεση κάτω από 120 mmHg και διαστολική πίεση κάτω από 80 mmHg",
	ChapterID:              1,
	CourseID:               1,
	Source:                 "Οδηγίες Αμερικανικής Καρδιολογικής Εταιρείας",
	MultipleCorrectAnswers: false,
	NumberOfAnswers:        1,
}

var question2 = repositories.Question{
	Text:                   "Ποια από τις παρακάτω είναι η πιο αποτελεσματική μέθοδος πρόληψης νοσοκομειακών λοιμώξεων;",
	Explanation:            "Η υγιεινή των χεριών είναι η πιο αποτελεσματική μέθοδος πρόληψης νοσοκομειακών λοιμώξεων και πρέπει να πραγματοποιείται πριν και μετά την επαφή με ασθενή",
	ChapterID:              2,
	CourseID:               1,
	Source:                 "Οδηγίες Έλεγχου Λοιμώξεων CDC",
	MultipleCorrectAnswers: false,
	NumberOfAnswers:        1,
}

var question3 = repositories.Question{
	Text:                   "Ποιος είναι ο κύριος σκοπός της νοσηλευτικής διαδικασίας;",
	Explanation:            "Η νοσηλευτική διαδικασία είναι μια συστηματική μέθοδος που χρησιμοποιούν οι νοσηλευτές για την παροχή εξατομικευμένης, κεντρωμένης στον ασθενή φροντίδας μέσω αξιολόγησης, διάγνωσης, σχεδιασμού, εφαρμογής και αξιολόγησης",
	ChapterID:              1,
	CourseID:               1,
	Source:                 "Βιβλίο Βασικών Αρχών Νοσηλευτικής",
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
	Text:       "Κάτω από 120/80 mmHg",
	QuestionID: 1,
	IsCorrect:  true,
	//TimesGiven: 0,
}

var answer3 = repositories.Answer{
	Text:       "Φοράω μόνο γάντια",
	QuestionID: 2,
	IsCorrect:  false,
	//TimesGiven: 0,
}

var answer4 = repositories.Answer{
	Text:       "Σωστή υγιεινή των χεριών",
	QuestionID: 2,
	IsCorrect:  true,
	//TimesGiven: 0,
}

var answer5 = repositories.Answer{
	Text:       "Να καταγράφω πληροφορίες ασθενούς",
	QuestionID: 3,
	IsCorrect:  false,
	//TimesGiven: 0,
}

var answer6 = repositories.Answer{
	Text:       "Να παρέχω συστηματική, κεντρωμένη στον ασθενή φροντίδα",
	QuestionID: 3,
	IsCorrect:  true,
	//TimesGiven: 0,
}

var quiz1 = repositories.Quiz{
	Title:       "Αξιολόγηση Βασικών Αρχών Νοσηλευτικής",
	Description: "Ολοκληρωμένο κουίζ που καλύπτει τις βασικές αρχές νοσηλευτικής και φροντίδας ασθενών",
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
	Title:       "Κουίζ Αξιολόγησης Ασθενούς",
	Description: "Δοκιμάστε τις γνώσεις σας για τεχνικές αξιολόγησης ασθενούς και σημεία ζωής",
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
	Title:       "Επανάληψη Νοσηλευτικής Διαδικασίας",
	Description: "Αξιολογήστε την κατανόησή σας για τη νοσηλευτική διαδικασία και τον σχεδιασμό φροντίδας",
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
	Title:       "Κουίζ Επαγγελματικής Πρακτικής",
	Description: "Αξιολόγηση επαγγελματικών προτύπων νοσηλευτικής και ηθικής πρακτικής",
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
	Title:       "Αξιολόγηση Ελέγχου Λοιμώξεων",
	Description: "Δοκιμάστε τις γνώσεις σας για πρόληψη και έλεγχο λοιμώξεων",
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
	Title:       "Κουίζ Πρωτοκόλλων Ασφάλειας",
	Description: "Αξιολόγηση πρωτοκόλλων ασφάλειας ασθενών και προτύπων υγειονομικής περίθαλψης",
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
	Title:       "Επανάληψη Ιατρικής-Χειρουργικής Νοσηλευτικής",
	Description: "Ολοκληρωμένη επανάληψη αρχών και πρακτικών ιατρικής-χειρουργικής νοσηλευτικής",
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
