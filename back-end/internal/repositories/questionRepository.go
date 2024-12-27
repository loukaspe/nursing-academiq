package repositories

import (
	"context"
	"errors"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

	modelQuestions := make([]*Question, 0, len(questions))
	err = repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
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
			modelQuestion.CourseID = courseID

			err = tx.Create(&modelQuestion).Error
			if err != nil {
				return err
			}
			modelQuestions = append(modelQuestions, &modelQuestion)
		}

		return nil
	})

	if err != nil {
		return err
	}

	// TODO remove fake quiz create
	modelQuiz := Quiz{}
	modelQuiz.Title = "Test Quiz"
	modelQuiz.Description = "This is just a test quiz"
	modelQuiz.CourseID = courseID
	modelQuiz.Questions = modelQuestions

	return repo.db.WithContext(ctx).Create(&modelQuiz).Error
}

func (repo *QuestionRepository) CreateQuestion(
	ctx context.Context,
	question *domain.Question,
) (uint, error) {
	var err error

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
	modelQuestion.CourseID = uint(question.Course.ID)
	modelQuestion.ChapterID = uint(question.Chapter.ID)
	modelQuestion.Answers = modelAnswers

	err = repo.db.WithContext(ctx).Create(&modelQuestion).Error

	return modelQuestion.ID, err
}

func (repo *QuestionRepository) GetChapterAndQuestionsByCourseID(
	ctx context.Context,
	courseID uint32,
) (domain.Course, error) {
	var err error
	var modelCourse Course

	err = repo.db.WithContext(ctx).
		Preload("Chapters.Questions").
		First(&modelCourse, courseID).Error

	if err == gorm.ErrRecordNotFound {
		return domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(courseID)) + " not found"),
		}
	}
	if err != nil {
		return domain.Course{}, err
	}

	domainCourse := domain.Course{
		ID:          uint32(modelCourse.ID),
		Title:       modelCourse.Title,
		Description: modelCourse.Description,
	}
	domainChapters := make([]domain.Chapter, 0, len(modelCourse.Chapters))

	for _, modelChapter := range modelCourse.Chapters {
		domainQuestions := make([]domain.Question, 0)
		domainChapter := domain.Chapter{
			ID:    uint32(modelChapter.ID),
			Title: modelChapter.Title,
		}

		for _, modelQuestion := range modelChapter.Questions {
			domainAnswers := make([]domain.Answer, 0, modelQuestion.NumberOfAnswers)
			for _, modelAnswer := range modelQuestion.Answers {
				domainAnswer := &domain.Answer{
					Text:      modelAnswer.Text,
					IsCorrect: modelAnswer.IsCorrect,
				}

				domainAnswers = append(domainAnswers, *domainAnswer)
			}

			domainQuestion := &domain.Question{
				ID:                     uint32(modelQuestion.ID),
				Text:                   modelQuestion.Text,
				Explanation:            modelQuestion.Explanation,
				Source:                 modelQuestion.Source,
				MultipleCorrectAnswers: modelQuestion.MultipleCorrectAnswers,
				NumberOfAnswers:        modelQuestion.NumberOfAnswers,
				Answers:                domainAnswers,
				Chapter: &domain.Chapter{
					ID: uint32(modelQuestion.ChapterID),
				},
			}

			domainQuestions = append(domainQuestions, *domainQuestion)
		}

		domainChapter.Questions = domainQuestions
		domainChapters = append(domainChapters, domainChapter)
	}

	domainCourse.Chapters = domainChapters

	return domainCourse, err
}

func (repo *QuestionRepository) UpdateQuestion(
	ctx context.Context,
	uid uint32,
	question *domain.Question,
) error {
	modelQuestion := &Question{}

	err := repo.db.WithContext(ctx).Model(&Question{}).Where("id = ?", uid).First(modelQuestion).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questionID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	// Remove previous answers to add the new ones
	err = repo.db.WithContext(ctx).Model(&Answer{}).Where("question_id = ?", modelQuestion.ID).Delete(&Answer{}).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questionID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	modelAnswers := make([]Answer, 0)
	for _, answer := range question.Answers {
		modelAnswer := Answer{}
		modelAnswer.Text = answer.Text
		modelAnswer.IsCorrect = answer.IsCorrect
		modelAnswer.QuestionID = modelQuestion.ID

		modelAnswers = append(modelAnswers, modelAnswer)
	}

	modelQuestion.Text = question.Text
	modelQuestion.Text = question.Text
	modelQuestion.Explanation = question.Explanation
	modelQuestion.Source = question.Source
	modelQuestion.MultipleCorrectAnswers = question.MultipleCorrectAnswers
	modelQuestion.NumberOfAnswers = question.NumberOfAnswers
	modelQuestion.CourseID = uint(question.Course.ID)
	modelQuestion.ChapterID = uint(question.Chapter.ID)
	modelQuestion.Answers = modelAnswers

	err = repo.db.WithContext(ctx).Save(&modelQuestion).Error

	return err
}

func (repo *QuestionRepository) DeleteQuestion(
	ctx context.Context,
	uid uint32,
) error {
	db := repo.db.WithContext(ctx).Model(&Question{}).
		Where("id = ?", uid).
		Take(&Question{}).
		Delete(&Question{})

	if db.Error == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questionID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return db.Error
}

func (repo *QuestionRepository) BulkDeleteQuestions(
	ctx context.Context,
	uids []uint32,
) error {
	questionsToBeDeleted := make([]Question, 0, len(uids))
	for _, uid := range uids {
		question := Question{}
		question.ID = uint(uid)
		questionsToBeDeleted = append(questionsToBeDeleted, question)
	}

	db := repo.db.WithContext(ctx).Model(&Question{}).
		Delete(questionsToBeDeleted)

	if db.Error == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("question not found"),
		}
	}

	return db.Error
}

func (repo *QuestionRepository) GetQuestion(
	ctx context.Context,
	uid uint32,
) (*domain.Question, error) {
	var err error
	var modelQuestion *Question

	err = repo.db.WithContext(ctx).
		//Preload("Tutor").
		Model(Question{}).
		Preload("Answers").
		Preload("Chapter").
		Where("id = ?", uid).
		Take(&modelQuestion).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Question{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questionID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Question{}, err
	}

	domainAnswers := make([]domain.Answer, 0, modelQuestion.NumberOfAnswers)
	for _, modelAnswer := range modelQuestion.Answers {
		domainAnswer := &domain.Answer{
			Text:      modelAnswer.Text,
			IsCorrect: modelAnswer.IsCorrect,
		}

		domainAnswers = append(domainAnswers, *domainAnswer)
	}

	return &domain.Question{
		ID:                     uint32(modelQuestion.ID),
		Text:                   modelQuestion.Text,
		Explanation:            modelQuestion.Explanation,
		Source:                 modelQuestion.Source,
		MultipleCorrectAnswers: modelQuestion.MultipleCorrectAnswers,
		NumberOfAnswers:        modelQuestion.NumberOfAnswers,
		Answers:                domainAnswers,
		Chapter: &domain.Chapter{
			ID:    uint32(modelQuestion.ChapterID),
			Title: modelQuestion.Chapter.Title,
		},
		Course: &domain.Course{
			ID: uint32(modelQuestion.CourseID),
		},
	}, err
}

func (repo *QuestionRepository) GetQuestions(
	ctx context.Context,
) ([]domain.Question, error) {
	var err error
	var modelQuestions []Question

	err = repo.db.WithContext(ctx).
		Model(Question{}).
		Find(&modelQuestions).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Question{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("questions not found"),
		}
	}
	if err != nil {
		return []domain.Question{}, err
	}

	var domainQuestions []domain.Question
	for _, modelQuestion := range modelQuestions {
		domainAnswers := make([]domain.Answer, 0, modelQuestion.NumberOfAnswers)
		for _, modelAnswer := range modelQuestion.Answers {
			domainAnswer := &domain.Answer{
				Text:      modelAnswer.Text,
				IsCorrect: modelAnswer.IsCorrect,
			}

			domainAnswers = append(domainAnswers, *domainAnswer)
		}

		domainQuestions = append(domainQuestions, domain.Question{
			ID:                     uint32(modelQuestion.ID),
			Text:                   modelQuestion.Text,
			Explanation:            modelQuestion.Explanation,
			Source:                 modelQuestion.Source,
			MultipleCorrectAnswers: modelQuestion.MultipleCorrectAnswers,
			NumberOfAnswers:        modelQuestion.NumberOfAnswers,
			Answers:                domainAnswers,
			Chapter: &domain.Chapter{
				ID: uint32(modelQuestion.ChapterID),
			},
			Course: &domain.Course{
				ID: uint32(modelQuestion.CourseID),
			},
		})
	}

	return domainQuestions, err
}
