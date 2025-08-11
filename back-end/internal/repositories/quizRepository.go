package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (repo *QuizRepository) GetQuiz(
	ctx context.Context,
	uid uint32,
) (*domain.Quiz, error) {
	var err error
	var modelQuiz *Quiz

	err = repo.db.WithContext(ctx).
		Model(Quiz{}).
		Preload("Questions.Answers").
		Preload("Course").
		Where("id = ?", uid).
		Take(&modelQuiz).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Quiz{}, err
	}

	domainQuestions := make([]domain.Question, 0, len(modelQuiz.Questions))
	for _, modelQuestion := range modelQuiz.Questions {

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
		}

		domainQuestions = append(domainQuestions, *domainQuestion)
	}

	return &domain.Quiz{
		ID:                uint32(modelQuiz.ID),
		Title:             modelQuiz.Title,
		Description:       modelQuiz.Description,
		Visibility:        modelQuiz.Visibility,
		ShowSubset:        modelQuiz.ShowSubset,
		SubsetSize:        modelQuiz.SubsetSize,
		ScoreSum:          modelQuiz.ScoreSum,
		MaxScore:          modelQuiz.MaxScore,
		NumberOfQuestions: len(modelQuiz.Questions),
		Questions:         domainQuestions,
		Course: &domain.Course{
			ID:          uint32(modelQuiz.CourseID),
			Title:       modelQuiz.Course.Title,
			Description: modelQuiz.Course.Description,
		},
	}, err
}

func (repo *QuizRepository) SearchQuiz(
	ctx context.Context,
	title string,
	courseName string, // optional filter
) ([]domain.Quiz, error) {
	var err error
	var modelQuizzes []*Quiz

	query := repo.db.WithContext(ctx).
		Model(&Quiz{}).
		Preload("Questions.Answers").
		Preload("Course").
		Where("quizzes.title LIKE ?", "%"+title+"%")

	if courseName != "" {
		query = query.Joins("JOIN courses c ON c.id = quizzes.course_id").
			Where("c.title = ?", courseName)
	}

	err = query.Find(&modelQuizzes).Error
	if err != nil {
		return nil, err
	}

	if len(modelQuizzes) == 0 {
		return nil, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      fmt.Errorf("no quizzes found with title LIKE '%s'", title),
		}
	}

	var domainQuizzes []domain.Quiz

	for _, modelQuiz := range modelQuizzes {
		domainQuestions := make([]domain.Question, 0, len(modelQuiz.Questions))
		for _, modelQuestion := range modelQuiz.Questions {

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
			}

			domainQuestions = append(domainQuestions, *domainQuestion)
		}

		domainQuizzes = append(domainQuizzes, domain.Quiz{
			ID:                uint32(modelQuiz.ID),
			Title:             modelQuiz.Title,
			Description:       modelQuiz.Description,
			Visibility:        modelQuiz.Visibility,
			ShowSubset:        modelQuiz.ShowSubset,
			SubsetSize:        modelQuiz.SubsetSize,
			ScoreSum:          modelQuiz.ScoreSum,
			MaxScore:          modelQuiz.MaxScore,
			NumberOfQuestions: len(modelQuiz.Questions),
			Questions:         domainQuestions,
			Course: &domain.Course{
				ID:          uint32(modelQuiz.CourseID),
				Title:       modelQuiz.Course.Title,
				Description: modelQuiz.Course.Description,
			},
		})
	}

	return domainQuizzes, nil
}

func (repo *QuizRepository) GetQuizByTutorID(
	ctx context.Context,
	tutorID uint32,
) ([]domain.Quiz, error) {
	var err error
	var modelTutor Tutor

	err = repo.db.WithContext(ctx).
		Preload("Courses.Quizs.Questions").
		First(&modelTutor, tutorID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("tutorID " + strconv.Itoa(int(tutorID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.Quiz{}, err
	}

	var domainQuizs []domain.Quiz
	for _, modelCourse := range modelTutor.Courses {
		courseName := modelCourse.Title

		for _, modelQuiz := range modelCourse.Quizs {
			var numberOfQuestions int
			for _, _ = range modelQuiz.Questions {
				numberOfQuestions++
			}

			domainQuizs = append(domainQuizs, domain.Quiz{
				ID:                uint32(modelQuiz.ID),
				Title:             modelQuiz.Title,
				Description:       modelQuiz.Description,
				Visibility:        modelQuiz.Visibility,
				ShowSubset:        modelQuiz.ShowSubset,
				SubsetSize:        modelQuiz.SubsetSize,
				ScoreSum:          modelQuiz.ScoreSum,
				MaxScore:          modelQuiz.MaxScore,
				NumberOfQuestions: numberOfQuestions,
				Course: &domain.Course{
					ID:    uint32(modelCourse.ID),
					Title: courseName,
				},
			})
		}
	}

	return domainQuizs, err
}

func (repo *QuizRepository) GetQuizByCourseID(
	ctx context.Context,
	courseID uint32,
) ([]domain.Quiz, error) {
	var err error
	var modelCourse Course

	err = repo.db.WithContext(ctx).
		Preload("Quizs.Questions").
		Where("visibility = ?", true).
		First(&modelCourse, courseID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(courseID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.Quiz{}, err
	}

	var domainQuizs []domain.Quiz

	for _, modelQuiz := range modelCourse.Quizs {
		domainQuestions := make([]domain.Question, 0, len(modelQuiz.Questions))
		for _, modelQuestion := range modelQuiz.Questions {

			domainAnswers := make([]domain.Answer, 0, modelQuestion.NumberOfAnswers)
			for _, modelAnswer := range modelQuestion.Answers {
				domainAnswer := &domain.Answer{
					Text:      modelAnswer.Text,
					IsCorrect: modelAnswer.IsCorrect,
				}

				domainAnswers = append(domainAnswers, *domainAnswer)
			}

			domainQuestion := &domain.Question{
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

		domainQuizs = append(domainQuizs, domain.Quiz{
			ID:                uint32(modelQuiz.ID),
			Title:             modelQuiz.Title,
			Description:       modelQuiz.Description,
			Visibility:        modelQuiz.Visibility,
			ShowSubset:        modelQuiz.ShowSubset,
			SubsetSize:        modelQuiz.SubsetSize,
			ScoreSum:          modelQuiz.ScoreSum,
			MaxScore:          modelQuiz.MaxScore,
			NumberOfQuestions: len(modelQuiz.Questions),
			Course: &domain.Course{
				Title: modelCourse.Title,
			},
			Questions: domainQuestions,
		})
	}

	return domainQuizs, err
}

func (repo *QuizRepository) GetQuizzes(
	ctx context.Context,
) ([]domain.Quiz, error) {
	var err error
	var modelQuizzes []Quiz

	err = repo.db.WithContext(ctx).
		Model(Quiz{}).
		Where("visibility = ?", true).
		Preload("Questions.Answers").
		Preload("Course").
		Find(&modelQuizzes).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizzes not found"),
		}
	}
	if err != nil {
		return []domain.Quiz{}, err
	}

	var domainQuizzes []domain.Quiz
	for _, modelQuiz := range modelQuizzes {
		domainQuestions := make([]domain.Question, 0, len(modelQuiz.Questions))
		for _, modelQuestion := range modelQuiz.Questions {

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
			}

			domainQuestions = append(domainQuestions, *domainQuestion)
		}

		domainQuiz := &domain.Quiz{
			ID:                uint32(modelQuiz.ID),
			Title:             modelQuiz.Title,
			Description:       modelQuiz.Description,
			Visibility:        modelQuiz.Visibility,
			ShowSubset:        modelQuiz.ShowSubset,
			SubsetSize:        modelQuiz.SubsetSize,
			ScoreSum:          modelQuiz.ScoreSum,
			MaxScore:          modelQuiz.MaxScore,
			NumberOfQuestions: len(modelQuiz.Questions),
			Questions:         domainQuestions,
			Course: &domain.Course{
				ID:          uint32(modelQuiz.CourseID),
				Title:       modelQuiz.Course.Title,
				Description: modelQuiz.Course.Description,
			},
		}

		domainQuizzes = append(domainQuizzes, *domainQuiz)
	}

	return domainQuizzes, err
}

func (repo *QuizRepository) GetMostRecentQuizzes(
	ctx context.Context,
	limit int,
) ([]domain.Quiz, error) {
	var err error
	var modelQuizzes []Quiz

	if limit > 0 {
		err = repo.db.WithContext(ctx).
			Order("created_at DESC").
			Limit(limit).
			Model(Quiz{}).
			Where("visibility = ?", true).
			Preload("Questions.Answers").
			Preload("Course").
			Find(&modelQuizzes).Error
	} else {
		err = repo.db.WithContext(ctx).
			Order("created_at DESC").
			Model(Quiz{}).
			Where("visibility = ?", true).
			Preload("Questions.Answers").
			Preload("Course").
			Find(&modelQuizzes).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []domain.Quiz{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizzes not found"),
		}
	}
	if err != nil {
		return []domain.Quiz{}, err
	}

	var domainQuizzes []domain.Quiz
	for _, modelQuiz := range modelQuizzes {
		domainQuestions := make([]domain.Question, 0, len(modelQuiz.Questions))
		for _, modelQuestion := range modelQuiz.Questions {

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
			}

			domainQuestions = append(domainQuestions, *domainQuestion)
		}

		domainQuiz := &domain.Quiz{
			ID:                uint32(modelQuiz.ID),
			Title:             modelQuiz.Title,
			Description:       modelQuiz.Description,
			Visibility:        modelQuiz.Visibility,
			ShowSubset:        modelQuiz.ShowSubset,
			SubsetSize:        modelQuiz.SubsetSize,
			ScoreSum:          modelQuiz.ScoreSum,
			MaxScore:          modelQuiz.MaxScore,
			NumberOfQuestions: len(modelQuiz.Questions),
			Questions:         domainQuestions,
			Course: &domain.Course{
				ID:          uint32(modelQuiz.CourseID),
				Title:       modelQuiz.Course.Title,
				Description: modelQuiz.Course.Description,
			},
		}

		domainQuizzes = append(domainQuizzes, *domainQuiz)
	}

	return domainQuizzes, err
}

func (repo *QuizRepository) UpdateQuiz(
	ctx context.Context,
	uid uint32,
	quiz *domain.Quiz,
	questionsIDs []uint32,
) error {
	modelQuiz := &Quiz{}

	err := repo.db.WithContext(ctx).Model(&Quiz{}).Where("id = ?", uid).First(modelQuiz).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return err
	}

	modelQuiz.Title = quiz.Title
	modelQuiz.Description = quiz.Description
	modelQuiz.Visibility = quiz.Visibility
	modelQuiz.ShowSubset = quiz.ShowSubset
	modelQuiz.SubsetSize = quiz.SubsetSize

	var questions []*Question
	if len(questionsIDs) > 0 {
		if err := repo.db.WithContext(ctx).Find(&questions, questionsIDs).Error; err != nil {
			return apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New(fmt.Sprintf("questionIDs %v not found", questionsIDs)),
			}
		}

		if err := repo.db.WithContext(ctx).Model(&modelQuiz).Association("Questions").Replace(questions); err != nil {
			return fmt.Errorf("could not add questions for quiz with ID %d: %w", modelQuiz.ID, err)
		}
	}

	if len(questionsIDs) == 0 {
		err = repo.db.WithContext(ctx).Model(&modelQuiz).Association("Questions").Clear()
		if err != nil {
			return fmt.Errorf("could not update questions for quiz with ID %d: %w", uid, err)
		}

		return nil
	}

	err = repo.db.WithContext(ctx).Save(&modelQuiz).Error

	return err
}

func (repo *QuizRepository) UpdateQuizQuestions(
	ctx context.Context,
	uid uint32,
	questionsIDS []uint32,
) error {
	var quiz Quiz
	if err := repo.db.WithContext(ctx).Preload("Questions").First(&quiz, uid).Error; err != nil {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("quizID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	if len(questionsIDS) == 0 && len(quiz.Questions) == 0 {
		return nil
	}

	if len(questionsIDS) == 0 {
		if err := repo.db.WithContext(ctx).Model(&quiz).Association("Questions").Clear().Error; err != nil {
			return fmt.Errorf("could not update questions for quiz with ID %d: %w", questionsIDS, err)
		}

		return nil
	}

	var questions []*Question
	if err := repo.db.WithContext(ctx).Find(&questions, questionsIDS).Error; err != nil {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New(fmt.Sprintf("questionIDs %v not found", questionsIDS)),
		}
	}

	return nil
}

func (repo *QuizRepository) DeleteQuiz(
	ctx context.Context,
	uid uint32,
) error {
	var quiz Quiz

	// Load the quiz with questions
	if err := repo.db.WithContext(ctx).
		Preload("Questions").
		First(&quiz, uid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      fmt.Errorf("quizID %d not found", uid),
			}
		}
		return err
	}

	// Start transaction
	return repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Clear the join table to remove associations
		if err := tx.Model(&quiz).Association("Questions").Clear(); err != nil {
			return err
		}

		// Delete the quiz
		if err := tx.Delete(&quiz).Error; err != nil {
			return err
		}

		return nil
	})
}

func (repo *QuizRepository) CreateQuiz(
	ctx context.Context,
	quiz *domain.Quiz,
	questionsIDs []uint32,
) (uint, error) {
	var err error

	modelQuiz := Quiz{}
	modelQuiz.Title = quiz.Title
	modelQuiz.Description = quiz.Description
	modelQuiz.Visibility = quiz.Visibility
	modelQuiz.ShowSubset = quiz.ShowSubset
	modelQuiz.SubsetSize = quiz.SubsetSize
	modelQuiz.CourseID = uint(quiz.Course.ID)

	err = repo.db.WithContext(ctx).Create(&modelQuiz).Error

	var questions []*Question
	if len(questionsIDs) > 0 {
		if err := repo.db.WithContext(ctx).Find(&questions, questionsIDs).Error; err != nil {
			return 0, apierrors.DataNotFoundErrorWrapper{
				ReturnedStatusCode: http.StatusNotFound,
				OriginalError:      errors.New(fmt.Sprintf("questionIDs %v not found", questionsIDs)),
			}
		}

		if err := repo.db.WithContext(ctx).Model(&modelQuiz).Association("Questions").Replace(questions); err != nil {
			return 0, fmt.Errorf("could not add questions for quiz with ID %d: %w", modelQuiz.ID, err)
		}
	}

	return modelQuiz.ID, err
}
