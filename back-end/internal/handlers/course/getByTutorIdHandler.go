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

type GetCourseByTutorIDHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetCourseByTutorIDHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetCourseByTutorIDHandler {
	return &GetCourseByTutorIDHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetCourseByTutorIDResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Courses      []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"courses,omitempty"`
}

func (handler *GetCourseByTutorIDHandler) GetCourseByTutorIDController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetCourseByTutorIDResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing tutor id"
		json.NewEncoder(w).Encode(response)
		return
	}
	tutorId, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed tutor id"
		json.NewEncoder(w).Encode(response)
		return
	}

	courses, err := handler.CourseService.GetCourseByTutorID(context.TODO(), uint32(tutorId))
	if dataNotFoundErrorWrapper, ok := err.(*apierrors.DataNotFoundErrorWrapper); ok {
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
			Title       string `json:"title"`
			Description string `json:"description"`
		}{
			Title:       course.Title,
			Description: course.Description,
		})
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
