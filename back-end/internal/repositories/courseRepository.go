package repositories

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	"gorm.io/gorm"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (repo *CourseRepository) GetCourse(
	ctx context.Context,
	uid uint32,
) (*domain.Course, error) {
	var err error
	var modelCourse *Course

	err = repo.db.WithContext(ctx).
		Preload("Tutor").
		Model(Course{}).
		Where("id = ?", uid).
		Take(&modelCourse).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Course{}, err
	}

	// TODO: preload Tutor if needed
	//domainUser := domain.User{
	//	Username:    modelCourse.User.Username,
	//	Password:    modelCourse.User.Password,
	//	FirstName:   modelCourse.User.FirstName,
	//	LastName:    modelCourse.User.LastName,
	//	Email:       modelCourse.User.Email,
	//	BirthDate:   modelCourse.User.BirthDate,
	//	PhoneNumber: modelCourse.User.PhoneNumber,
	//}

	return &domain.Course{
		Title:       modelCourse.Title,
		Description: modelCourse.Description,
		Tutor: &domain.Tutor{
			ID: modelCourse.TutorID,
		},
	}, err
}

func (repo *CourseRepository) GetExtendedCourse(
	ctx context.Context,
	uid uint32,
) (*domain.Course, error) {
	var err error
	var modelCourse *Course

	err = repo.db.WithContext(ctx).
		Preload("Quizs.Questions").
		Preload("Chapters").
		Preload("Tutor.User").
		Model(Course{}).
		Where("id = ?", uid).
		Take(&modelCourse).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Course{}, err
	}

	domainTutor := domain.Tutor{
		User: domain.User{
			FirstName: modelCourse.Tutor.User.FirstName,
			LastName:  modelCourse.Tutor.User.LastName,
		},
	}

	var domainQuizs []domain.Quiz

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
				Title: modelCourse.Title,
			},
		})
	}

	var domainChapters []domain.Chapter

	for _, modelChapter := range modelCourse.Chapters {
		domainChapters = append(domainChapters, domain.Chapter{
			ID:          uint32(modelChapter.ID),
			Title:       modelChapter.Title,
			Description: modelChapter.Description,
		})
	}

	return &domain.Course{
		Title:       modelCourse.Title,
		Description: modelCourse.Description,
		Quizzes:     domainQuizs,
		Chapters:    domainChapters,
		Tutor:       &domainTutor,
	}, err
}

func (repo *CourseRepository) GetCourseChapters(
	ctx context.Context,
	uid uint32,
) (*domain.Course, error) {
	var err error
	var modelCourse *Course

	err = repo.db.WithContext(ctx).
		Preload("Chapters").
		Model(Course{}).
		Where("id = ?", uid).
		Take(&modelCourse).Error

	if err == gorm.ErrRecordNotFound {
		return &domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}
	if err != nil {
		return &domain.Course{}, err
	}

	var domainChapters []domain.Chapter

	for _, modelChapter := range modelCourse.Chapters {
		domainChapters = append(domainChapters, domain.Chapter{
			ID:          uint32(modelChapter.ID),
			Title:       modelChapter.Title,
			Description: modelChapter.Description,
		})
	}

	return &domain.Course{
		ID:          uint32(modelCourse.ID),
		Title:       modelCourse.Title,
		Description: modelCourse.Description,
		Chapters:    domainChapters,
	}, err
}

func (repo *CourseRepository) GetCourses(
	ctx context.Context,
) ([]domain.Course, error) {
	var err error
	var modelCourses []Course

	err = repo.db.WithContext(ctx).
		Preload("Questions").
		Preload("Tutor").
		Preload("Quizs.Questions").
		Preload("Chapters").
		Model(Course{}).
		Find(&modelCourses).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courses not found"),
		}
	}
	if err != nil {
		return []domain.Course{}, err
	}

	var domainCourses []domain.Course
	for _, modelCourse := range modelCourses {
		domainTutor := domain.Tutor{
			ID: modelCourse.TutorID,
			User: domain.User{
				FirstName: modelCourse.Tutor.User.FirstName,
				LastName:  modelCourse.Tutor.User.LastName,
			},
		}

		var domainQuizs []domain.Quiz

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
					Title: modelCourse.Title,
				},
			})
		}

		var domainChapters []domain.Chapter

		for _, modelChapter := range modelCourse.Chapters {
			domainChapters = append(domainChapters, domain.Chapter{
				ID:          uint32(modelChapter.ID),
				Title:       modelChapter.Title,
				Description: modelChapter.Description,
			})
		}

		domainCourses = append(domainCourses, domain.Course{
			ID:                uint32(modelCourse.ID),
			Title:             modelCourse.Title,
			Description:       modelCourse.Description,
			Quizzes:           domainQuizs,
			Chapters:          domainChapters,
			Tutor:             &domainTutor,
			NumberOfQuestions: len(modelCourse.Questions),
		})
	}

	return domainCourses, err
}

func (repo *CourseRepository) GetMostRecentCourses(
	ctx context.Context,
	limit int,
) ([]domain.Course, error) {
	var err error
	var modelCourses []Course

	if limit > 0 {
		err = repo.db.WithContext(ctx).
			Order("created_at DESC").
			Limit(limit).
			Model(Course{}).
			Find(&modelCourses).Error
	} else {
		err = repo.db.WithContext(ctx).
			Order("created_at DESC").
			Model(Course{}).
			Find(&modelCourses).Error
	}

	if err == gorm.ErrRecordNotFound {
		return []domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courses not found"),
		}
	}
	if err != nil {
		return []domain.Course{}, err
	}

	var domainCourses []domain.Course
	for _, modelCourse := range modelCourses {
		domainCourses = append(domainCourses, domain.Course{
			ID:          uint32(modelCourse.ID),
			Title:       modelCourse.Title,
			Description: modelCourse.Description,
		})
	}

	return domainCourses, err
}

func (repo *CourseRepository) GetCourseByTutorID(
	ctx context.Context,
	tutorID uint32,
) ([]domain.Course, error) {
	var err error
	var modelTutor Tutor

	err = repo.db.WithContext(ctx).
		Preload("Courses").
		First(&modelTutor, tutorID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("tutorID " + strconv.Itoa(int(tutorID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.Course{}, err
	}

	var domainCourses []domain.Course
	for _, modelCourse := range modelTutor.Courses {
		// TODO: preload Tutor, Students if needed
		domainCourses = append(domainCourses, domain.Course{
			ID:          uint32(modelCourse.ID),
			Title:       modelCourse.Title,
			Description: modelCourse.Description,
		})
	}

	return domainCourses, err
}

func (repo *CourseRepository) UpdateCourse(
	ctx context.Context,
	uid uint32,
	course *domain.Course,
) error {
	modelCourse := &Course{}

	// Only Title and Description are editable
	partialUpdates := make(map[string]interface{}, 2)
	if course.Title != "" {
		partialUpdates["title"] = course.Title
	}
	if course.Description != "" {
		partialUpdates["description"] = course.Description
	}

	err := repo.db.WithContext(ctx).Model(modelCourse).Where("id = ?", uid).Updates(partialUpdates).Error
	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return err
}

func (repo *CourseRepository) DeleteCourse(
	ctx context.Context,
	uid uint32,
) error {
	err := repo.db.WithContext(ctx).Model(&Course{}).
		Where("id = ?", uid).
		Take(&Course{}).
		Delete(&Course{}).Error

	if err == gorm.ErrRecordNotFound {
		return apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("courseID " + strconv.Itoa(int(uid)) + " not found"),
		}
	}

	return err
}

func (repo *CourseRepository) CreateCourse(
	ctx context.Context,
	course *domain.Course,
	tutorID uint,
) (uint, error) {
	var err error

	modelCourse := Course{}
	modelCourse.Title = course.Title
	modelCourse.Description = course.Description
	modelCourse.TutorID = tutorID

	err = repo.db.WithContext(ctx).Create(&modelCourse).Error

	return modelCourse.ID, err
}
