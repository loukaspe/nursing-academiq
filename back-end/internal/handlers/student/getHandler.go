package student

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

type GetStudentHandler struct {
	StudentService *services.StudentService
	logger         *log.Logger
}

func NewGetStudentHandler(
	service *services.StudentService,
	logger *log.Logger,
) *GetStudentHandler {
	return &GetStudentHandler{
		StudentService: service,
		logger:         logger,
	}
}

type GetStudentResponse struct {
	ErrorMessage string          `json:"errorMessage,omitempty"`
	Student      *StudentRequest `json:"student,omitempty"`
}

func (handler *GetStudentHandler) GetStudentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetStudentResponse{}

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

	student, err := handler.StudentService.GetStudent(context.TODO(), uint32(uid))
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

	response.Student = &StudentRequest{
		Student: struct {
			Username           string `json:"username"`
			Password           string `json:"password"`
			FirstName          string `json:"first_name"`
			LastName           string `json:"last_name"`
			Email              string `json:"email"`
			BirthDate          string `json:"birth_date"`
			PhoneNumber        string `json:"phone_number"`
			Photo              string `json:"photo"`
			RegistrationNumber string `json:"registration_number"`
		}{
			Username:           student.Username,
			Password:           student.Password,
			FirstName:          student.FirstName,
			LastName:           student.LastName,
			Email:              student.Email,
			BirthDate:          student.BirthDate.String(),
			PhoneNumber:        student.PhoneNumber,
			Photo:              student.Photo,
			RegistrationNumber: student.RegistrationNumber,
		},
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
