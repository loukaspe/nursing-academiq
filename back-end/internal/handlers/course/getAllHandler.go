package course

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type GetCoursesHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetCoursesHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetCourseHandler {
	return &GetCourseHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetCoursesResponse struct {
	ErrorMessage string          `json:"errorMessage,omitempty"`
	Courses      []CourseRequest `json:"courses,omitempty"`
}

func (handler *GetCourseHandler) GetCoursesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetCoursesResponse{}

	courses, err := handler.CourseService.GetCourses(context.TODO())
	if dataNotFoundErrorWrapper, ok := err.(*apierrors.DataNotFoundErrorWrapper); ok {
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
		response.Courses = append(response.Courses, CourseRequest{
			Course: struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			}{
				Title:       course.Title,
				Description: course.Description,
			},
		})
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
