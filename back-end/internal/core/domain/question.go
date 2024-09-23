package domain

type Question struct {
	ID                     uint32
	Chapter                *Chapter
	Course                 *Course
	Text                   string
	Explanation            string
	Source                 string
	MultipleCorrectAnswers bool
	NumberOfAnswers        int
	//TimesAnswered          int
	//TimesAnsweredCorrectly float32
	Answers []Answer
}
