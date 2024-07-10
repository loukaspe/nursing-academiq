package domain

type Answer struct {
	Question  *Question
	Text      string
	IsCorrect bool
	//TimesGiven int
}
