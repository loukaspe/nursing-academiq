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

// the student model needed to create a new student
// it also creates a new user
//
// swagger:parameters createStudent
type CreateStudentRequest struct {
	// in:body
	Student struct {
		// Required: true
		Username string `json:"username"`
		// Required: true
		Password string `json:"password"`
		// Required: true
		FirstName string `json:"first_name"`
		// Required: true
		LastName string `json:"last_name"`
		// Required: true
		Email string `json:"email"`
		// Required: true
		BirthDate string `json:"birth_date"`
		// Required: true
		PhoneNumber string `json:"phone_number"`
		// Required: true
		Photo string `json:"photo"`
		// Required: true
		RegistrationNumber string `json:"registration_number"`
	} `json:""`
}

// Response when we insert a new Student
// swagger:model CreateStudentResponse
type CreateStudentResponse struct {
	// inserted student uid
	//
	// Required: true
	CreatedStudentUid uint `json:"insertedUid"`
	// possible error message
	//
	// Required: false
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// swagger:operation POST /student createStudent
//
// # It creates a new student along with a new user
//
// ---
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes:
//	- http
//	- https
//
//	responses:
//		"200":
//			description: Student created successfully
//			schema:
//				$ref: "#/definitions/CreateStudentResponse"
//		"400":
//			description: Bad request - request parameters are missing or invalid
//			schema:
//				$ref: "#/definitions/CreateStudentResponse"
//		"500":
//			description: Internal server error - check logs for details
//			schema:
//				$ref: "#/definitions/CreateStudentResponse"
func (handler *CreateStudentHandler) CreateStudentController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &CreateStudentResponse{}
	request := &CreateStudentRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating student")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed student data"
		json.NewEncoder(w).Encode(response)
		return
	}

	studentRequest := request.Student

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
