package repositories

import (
	"context"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (repo *QuestionRepository) ImportForCourse(
	ctx context.Context,
	questions []domain.Question,
	courseID uint,
) error {
	var err error

	return repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, question := range questions {
			var chapter Chapter
			result := tx.Where("title = ?", question.Chapter.Title).First(&chapter)
			if result.Error == gorm.ErrRecordNotFound {
				// Chapter does not exist, create it.
				chapter = Chapter{
					Title:    question.Chapter.Title,
					CourseID: courseID,
				}
				err = tx.Create(&chapter).Error
				if err != nil {
					return err
				}
			}

			modelAnswers := make([]Answer, 0)
			for _, answer := range question.Answers {
				modelAnswer := Answer{}
				modelAnswer.Text = answer.Text
				modelAnswer.IsCorrect = answer.IsCorrect

				modelAnswers = append(modelAnswers, modelAnswer)
			}

			modelQuestion := Question{}
			modelQuestion.Text = question.Text
			modelQuestion.Explanation = question.Explanation
			modelQuestion.Source = question.Source
			modelQuestion.MultipleCorrectAnswers = question.MultipleCorrectAnswers
			modelQuestion.NumberOfAnswers = question.NumberOfAnswers
			modelQuestion.Answers = modelAnswers
			modelQuestion.ChapterID = chapter.ID

			err = tx.Create(&modelQuestion).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}
