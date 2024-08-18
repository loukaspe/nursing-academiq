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

type GetCourseHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetCourseHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetCourseHandler {
	return &GetCourseHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetCourseResponse struct {
	ErrorMessage string         `json:"errorMessage,omitempty"`
	Course       *CourseRequest `json:"course,omitempty"`
}

func (handler *GetCourseHandler) GetCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetCourseResponse{}

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

	course, err := handler.CourseService.GetCourse(context.Background(), uint32(uid))
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

	response.Course = &CourseRequest{
		Course: struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			TutorID     uint   `json:"tutorID"`
		}{
			Title:       course.Title,
			Description: course.Description,
			TutorID:     course.Tutor.ID,
		},
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
