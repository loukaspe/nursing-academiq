package course

import (
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
) *GetCoursesHandler {
	return &GetCoursesHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetCoursesResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Courses      []struct {
		ID                uint32 `json:"id"`
		Title             string `json:"title"`
		Description       string `json:"description"`
		NumberOfQuestions int    `json:"numberOfQuestions"`
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
		response.Courses = append(response.Courses, struct {
			ID                uint32 `json:"id"`
			Title             string `json:"title"`
			Description       string `json:"description"`
			NumberOfQuestions int    `json:"numberOfQuestions"`
		}{
			ID:                course.ID,
			Title:             course.Title,
			Description:       course.Description,
			NumberOfQuestions: course.NumberOfQuestions,
		})
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
