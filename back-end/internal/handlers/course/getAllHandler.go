package course

import (
	"encoding/json"
	"net/http"

	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type GetCoursesHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetCoursesHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetCoursesHandler {
	return &GetCoursesHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetCoursesResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Courses      []struct {
		ID                uint32    `json:"id"`
		Title             string    `json:"title"`
		Description       string    `json:"description"`
		Quizzes           []Quiz    `json:"quizzes,omitempty"`
		Chapters          []Chapter `json:"chapters,omitempty"`
		TutorName         string    `json:"tutorName"`
		NumberOfQuestions int       `json:"numberOfQuestions"`
		TutorID           uint      `json:"tutorID"`
	} `json:"courses,omitempty"`
}

func (handler *GetCoursesHandler) GetCoursesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &GetCoursesResponse{}

	courses, err := handler.CourseService.GetCourses(r.Context())
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting courses")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, course := range courses {
		responseCourse := struct {
			ID                uint32    `json:"id"`
			Title             string    `json:"title"`
			Description       string    `json:"description"`
			Quizzes           []Quiz    `json:"quizzes,omitempty"`
			Chapters          []Chapter `json:"chapters,omitempty"`
			TutorName         string    `json:"tutorName"`
			NumberOfQuestions int       `json:"numberOfQuestions"`
			TutorID           uint      `json:"tutorID"`
		}{
			ID:                course.ID,
			Title:             course.Title,
			Description:       course.Description,
			NumberOfQuestions: course.NumberOfQuestions,
			TutorName:         course.Tutor.User.FirstName + " " + course.Tutor.User.LastName,
			TutorID:           course.Tutor.ID,
		}

		for _, quiz := range course.Quizzes {
			responseCourse.Quizzes = append(responseCourse.Quizzes, Quiz{
				ID:                quiz.ID,
				Title:             quiz.Title,
				NumberOfQuestions: quiz.NumberOfQuestions,
				CourseName:        quiz.Course.Title,
				ShowSubset:        quiz.ShowSubset,
				SubsetSize:        quiz.SubsetSize,
				Visibility:        quiz.Visibility,
			})
		}

		for _, chapter := range course.Chapters {
			responseCourse.Chapters = append(responseCourse.Chapters, struct {
				ID          uint32
				Title       string
				Description string
			}{
				ID:          chapter.ID,
				Title:       chapter.Title,
				Description: chapter.Description,
			})
		}

		response.Courses = append(response.Courses, responseCourse)
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
