package course

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetExtendedCourseHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetExtendedCourseHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetExtendedCourseHandler {
	return &GetExtendedCourseHandler{
		CourseService: service,
		logger:        logger,
	}
}

type Quiz struct {
	ID                uint32
	Title             string
	NumberOfQuestions int
	CourseName        string
}

type Chapter struct {
	ID          uint32
	Title       string
	Description string
}

type GetExtendedCourseResponse struct {
	ErrorMessage string    `json:"errorMessage,omitempty"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Quizzes      []Quiz    `json:"quizzes,omitempty"`
	Chapters     []Chapter `json:"chapters,omitempty"`
	TutorName    string    `json:"tutorName"`
}

func (handler *GetExtendedCourseHandler) GetExtendedCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetExtendedCourseResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing course id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed course id"
		json.NewEncoder(w).Encode(response)
		return
	}

	extendedCourse, err := handler.CourseService.GetExtendedCourse(context.Background(), uint32(uid))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting extended course data")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, quiz := range extendedCourse.Quizzes {
		response.Quizzes = append(response.Quizzes, Quiz{
			ID:                quiz.ID,
			Title:             quiz.Title,
			NumberOfQuestions: quiz.NumberOfQuestions,
			CourseName:        quiz.Course.Title,
		})
	}

	for _, chapter := range extendedCourse.Chapters {
		response.Chapters = append(response.Chapters, struct {
			ID          uint32
			Title       string
			Description string
		}{
			ID:          chapter.ID,
			Title:       chapter.Title,
			Description: chapter.Description,
		})
	}

	response.Title = extendedCourse.Title
	response.Description = extendedCourse.Description
	response.TutorName = extendedCourse.Tutor.User.FirstName + " " + extendedCourse.Tutor.User.LastName

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
