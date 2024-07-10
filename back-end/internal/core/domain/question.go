package domain

type Question struct {
	Chapter                *Chapter
	Text                   string
	Explanation            string
	Source                 string
	MultipleCorrectAnswers bool
	NumberOfAnswers        int
	//TimesAnswered          int
	//TimesAnsweredCorrectly float32
	Answers []Answer
}
