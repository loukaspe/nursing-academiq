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
	"time"
)

type UpdateStudentHandler struct {
	StudentService *services.StudentService
	logger         *log.Logger
}

func NewUpdateStudentHandler(
	service *services.StudentService,
	logger *log.Logger,
) *UpdateStudentHandler {
	return &UpdateStudentHandler{
		StudentService: service,
		logger:         logger,
	}
}

type UpdateStudentResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateStudentHandler) UpdateStudentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateStudentResponse{}

	request := &StudentRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating student")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed student request"
		json.NewEncoder(w).Encode(response)
		return
	}

	studentRequest := request.Student

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

	birthDate, err := time.Parse("17-03-2023", studentRequest.BirthDate)
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

	err = handler.StudentService.UpdateStudent(context.TODO(), uint32(uid), domainStudent)
	if dataNotFoundErrorWrapper, ok := err.(*apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating solar panel data")

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
