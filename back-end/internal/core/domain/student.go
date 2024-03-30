package domain

type Student struct {
	User
	RegistrationNumber string
	Courses            []Course
	QuizSessions       []QuizSessionByStudent
	// The following fields are used in the user profile page
	CompletedQuizzes int
	// a string in format "rightQuestions/totalQuestions"
	QuestionsScore             string
	PercentageOfCorrectAnswers string
}
