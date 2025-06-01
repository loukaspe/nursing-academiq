package course

import (
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Course struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TutorID     uint   `json:"tutorID"`
}

type CourseRequest struct {
	Course Course `json:""`
}

type CreateCourseResponse struct {
	CreatedCourseID uint   `json:"insertedID"`
	ErrorMessage    string `json:"errorMessage,omitempty"`
}

type CreateCourseHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewCreateCourseHandler(
	service *services.CourseService,
	logger *log.Logger,
) *CreateCourseHandler {
	return &CreateCourseHandler{
		CourseService: service,
		logger:        logger,
	}
}

func (handler *CreateCourseHandler) CreateCourseController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &CreateCourseResponse{}
	request := &CourseRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating course")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed course data"
		json.NewEncoder(w).Encode(response)
		return
	}

	courseRequest := request.Course

	domainCourse := &domain.Course{
		Title:       courseRequest.Title,
		Description: courseRequest.Description,
	}

	uid, err := handler.CourseService.CreateCourse(r.Context(), domainCourse, courseRequest.TutorID)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating course in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating course"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.CreatedCourseID = uid

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusCreated)
	return
}
