package course

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetMostRecentCoursesHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetMostRecentCoursesHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetCourseHandler {
	return &GetCourseHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetMostRecentCoursesResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Courses      []struct {
		ID          uint32 `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"courses,omitempty"`
}

func (handler *GetCourseHandler) GetMostRecentCoursesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetMostRecentCoursesResponse{}

	var limit int
	var err error

	limitParam := r.URL.Query().Get("limit")
	if limitParam == "" {

	} else {
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.ErrorMessage = "malformed limit"
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	courses, err := handler.CourseService.GetMostRecentCourses(context.Background(), limit)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting most recent courses")

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
