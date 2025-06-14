package course

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UpdateCourseRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateCourseHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewUpdateCourseHandler(
	service *services.CourseService,
	logger *log.Logger,
) *UpdateCourseHandler {
	return &UpdateCourseHandler{
		CourseService: service,
		logger:        logger,
	}
}

type UpdateCourseResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateCourseHandler) UpdateCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateCourseResponse{}

	request := &UpdateCourseRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating course")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed course request"
		json.NewEncoder(w).Encode(response)
		return
	}

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

	domainCourse := &domain.Course{
		Title:       *request.Title,
		Description: *request.Description,
	}

	err = handler.CourseService.UpdateCourse(r.Context(), uint32(uid), domainCourse)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating course")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
