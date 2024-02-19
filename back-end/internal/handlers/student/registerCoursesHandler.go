package student

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type RegisterStudentCoursesHandler struct {
	StudentService *services.StudentService
	logger         *log.Logger
}

func NewRegisterStudentCoursesHandler(
	service *services.StudentService,
	logger *log.Logger,
) *RegisterStudentCoursesHandler {
	return &RegisterStudentCoursesHandler{
		StudentService: service,
		logger:         logger,
	}
}

type RegisterStudentCoursesRequest struct {
	Courses []uint32 `json:"courses"`
}

type RegisterStudentCoursesResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Student      *struct {
		Username           string `json:"username"`
		FirstName          string `json:"first_name"`
		LastName           string `json:"last_name"`
		Email              string `json:"email"`
		BirthDate          string `json:"birth_date"`
		PhoneNumber        string `json:"phone_number"`
		Photo              string `json:"photo"`
		RegistrationNumber string `json:"registration_number"`
		Courses            []struct {
			ID    uint32 `json:"id"`
			Title string `json:"title"`
		} `json:"courses,omitempty"`
	} `json:"student,omitempty"`
}

// expects a list of course titles in order to register the student to the courses
func (handler *RegisterStudentCoursesHandler) RegisterStudentCoursesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &RegisterStudentCoursesResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing student id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed student id"
		json.NewEncoder(w).Encode(response)
		return
	}

	var coursesRequest RegisterStudentCoursesRequest
	err = json.NewDecoder(r.Body).Decode(&coursesRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed course data"
		json.NewEncoder(w).Encode(response)
		return
	}

	var domainCourses []domain.Course
	for _, courseID := range coursesRequest.Courses {
		domainCourses = append(domainCourses, domain.Course{
			ID: courseID,
		})
	}

	student, err := handler.StudentService.RegisterCourses(context.Background(), uint32(uid), domainCourses)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in registering student courses")

		response.ErrorMessage = dataNotFoundErrorWrapper.Unwrap().Error()

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)
		json.NewEncoder(w).Encode(response)
		return

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	var studentCourses []struct {
		ID    uint32 `json:"id"`
		Title string `json:"title"`
	}

	for _, course := range student.Courses {
		studentCourses = append(studentCourses, struct {
			ID    uint32 `json:"id"`
			Title string `json:"title"`
		}{
			ID:    course.ID,
			Title: course.Title,
		})
	}

	response.Student = &struct {
		Username           string `json:"username"`
		FirstName          string `json:"first_name"`
		LastName           string `json:"last_name"`
		Email              string `json:"email"`
		BirthDate          string `json:"birth_date"`
		PhoneNumber        string `json:"phone_number"`
		Photo              string `json:"photo"`
		RegistrationNumber string `json:"registration_number"`
		Courses            []struct {
			ID    uint32 `json:"id"`
			Title string `json:"title"`
		} `json:"courses,omitempty"`
	}{
		Username:           student.Username,
		FirstName:          student.FirstName,
		LastName:           student.LastName,
		Email:              student.Email,
		BirthDate:          student.BirthDate.String(),
		PhoneNumber:        student.PhoneNumber,
		Photo:              student.Photo,
		RegistrationNumber: student.RegistrationNumber,
		Courses:            studentCourses,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
