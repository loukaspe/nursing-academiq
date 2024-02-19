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

type GetCourseByStudentIDHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetCourseByStudentIDHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetCourseByStudentIDHandler {
	return &GetCourseByStudentIDHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetCourseByStudentIDResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Courses      []struct {
		ID          uint32 `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"courses,omitempty"`
}

func (handler *GetCourseByStudentIDHandler) GetCourseByStudentIDController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetCourseByStudentIDResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing student id"
		json.NewEncoder(w).Encode(response)
		return
	}
	studentId, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed student id"
		json.NewEncoder(w).Encode(response)
		return
	}

	courses, err := handler.CourseService.GetCourseByStudentID(context.TODO(), uint32(studentId))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting user data")

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
		response.Courses = append(response.Courses, struct {
			ID          uint32 `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
		}{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
		})
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
