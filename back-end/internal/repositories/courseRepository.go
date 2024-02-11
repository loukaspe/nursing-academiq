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
	//	Photo:       modelCourse.User.Photo,
	//}

	return &domain.Course{
		Title:       modelCourse.Title,
		Description: modelCourse.Description,
		Students:    nil,
	}, err
}

func (repo *CourseRepository) GetCourses(
	ctx context.Context,
) ([]domain.Course, error) {
	var err error
	var modelCourses []Course

	err = repo.db.WithContext(ctx).
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
		domainCourses = append(domainCourses, domain.Course{
			Title:       modelCourse.Title,
			Description: modelCourse.Description,
		})
	}

	return domainCourses, err
}

func (repo *CourseRepository) GetCourseByStudentID(
	ctx context.Context,
	studentID uint32,
) ([]domain.Course, error) {
	var err error
	//var modelCourses []Course
	var modelStudent Student

	//err = repo.db.WithContext(ctx).
	//	//Preload("Courses").
	//	Where("id = ?", studentID).
	//	Take(&modelStudent).
	//	Association("Courses").
	//	Find(&modelCourses)

	//err = repo.db.WithContext(ctx).
	//	Joins("Courses").
	//	Model(Student{}).
	//	Where("id = ?", studentID).
	//	Take(&modelStudent).Error

	err = repo.db.WithContext(ctx).
		Preload("Courses").
		First(&modelStudent, studentID).Error

	if err == gorm.ErrRecordNotFound {
		return []domain.Course{}, apierrors.DataNotFoundErrorWrapper{
			ReturnedStatusCode: http.StatusNotFound,
			OriginalError:      errors.New("studentID " + strconv.Itoa(int(studentID)) + " not found"),
		}
	}
	if err != nil {
		return []domain.Course{}, err
	}

	//var domainCourses []domain.Course
	//for _, modelCourse := range modelCourses {
	//	// TODO: preload Tutor, Students if needed
	//	domainCourses = append(domainCourses, domain.Course{
	//		Title:       modelCourse.Title,
	//		Description: modelCourse.Description,
	//	})
	//}

	var domainCourses []domain.Course
	for _, modelCourse := range modelStudent.Courses {
		// TODO: preload Tutor, Students if needed
		domainCourses = append(domainCourses, domain.Course{
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
			Title:       modelCourse.Title,
			Description: modelCourse.Description,
		})
	}

	return domainCourses, err
}

//func (repo *CourseRepository) UpdateCourse(
//	ctx context.Context,
//	uid uint32,
//	course *domain.Course,
//) error {
//	modelCourse := &Course{}
//
//	// TODO: handle Tutor preload if needed err := repo.db.WithContext(ctx).Preload("User").First(modelCourse).Error
//	err := repo.db.WithContext(ctx).Model(modelCourse).Where("id = ?", uid).Error
//	if err == gorm.ErrRecordNotFound {
//		return apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("courseID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//	if err != nil {
//		return err
//	}
//
//	modelCourse.Title = course.Title
//	modelCourse.Description = course.Description
//
//	err = repo.db.WithContext(ctx).Save(&modelCourse).Error
//
//	return err
//}
//
//func (repo *CourseRepository) DeleteCourse(
//	ctx context.Context,
//	uid uint32,
//) error {
//	db := repo.db.WithContext(ctx).Model(&Course{}).
//		Where("id = ?", uid).
//		Take(&Course{}).
//		Delete(&Course{})
//
//	if db.Error == gorm.ErrRecordNotFound {
//		return apierrors.DataNotFoundErrorWrapper{
//			ReturnedStatusCode: http.StatusNotFound,
//			OriginalError:      errors.New("courseID " + strconv.Itoa(int(uid)) + " not found"),
//		}
//	}
//
//	return db.Error
//}

//func (repo *CourseRepository) CreateCourse(
//	ctx context.Context,
//	course *domain.Course,
//	tutorID uint,
//) (uint, error) {
//	var err error
//
//	// TODO: add course tutor if needed
//	modelCourse := Course{}
//	modelCourse.Title = course.Title
//	modelCourse.Description = course.Description
//	modelCourse.TutorID = tutorID
//
//	err = repo.db.WithContext(ctx).Create(&modelCourse).Error
//
//	return modelCourse.ID, err
//}
