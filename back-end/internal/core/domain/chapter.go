package domain

type Chapter struct {
	ID          uint32
	Title       string
	Description string
	Course      *Course
	Questions   []Question
	// A chapter is considered to have a Chapter when even one of its questions belong to that Chapter
	Quizzes []Quiz
}
