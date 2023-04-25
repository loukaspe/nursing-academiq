package student

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CreateStudentHandler struct {
	StudentService *services.StudentService
	logger         *log.Logger
}

func NewCreateStudentHandler(
	service *services.StudentService,
	logger *log.Logger,
) *CreateStudentHandler {
	return &CreateStudentHandler{
		StudentService: service,
		logger:         logger,
	}
}

type Student struct {
	Username           string `json:"username"`
	Password           string `json:"password"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Email              string `json:"email"`
	BirthDate          string `json:"birth_date"`
	PhoneNumber        string `json:"phone_number"`
	Photo              string `json:"photo"`
	RegistrationNumber string `json:"registration_number"`
}

type CreateStudentResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *CreateStudentHandler) CreateStudentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &CreateStudentResponse{}
	studentRequest := &Student{}

	err := json.NewDecoder(r.Body).Decode(studentRequest)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating student")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed student data"
		json.NewEncoder(w).Encode(response)
		return
	}

	birthDate, err := time.Parse("01-02-2006", studentRequest.BirthDate)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating student birth date")

		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed student data: birth date"
		json.NewEncoder(w).Encode(response)
		return
	}

	domainUser := &domain.User{
		Username:    studentRequest.Username,
		Password:    studentRequest.Password,
		FirstName:   studentRequest.FirstName,
		LastName:    studentRequest.LastName,
		Email:       studentRequest.Email,
		BirthDate:   birthDate,
		PhoneNumber: studentRequest.PhoneNumber,
		Photo:       studentRequest.Photo,
	}

	domainStudent := &domain.Student{
		User:               *domainUser,
		RegistrationNumber: studentRequest.RegistrationNumber,
	}

	uid, err := handler.StudentService.CreateStudent(context.TODO(), domainStudent)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating student in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating student"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.CreatedStudentUid = uid

	w.WriteHeader(http.StatusCreated)
	return
}
