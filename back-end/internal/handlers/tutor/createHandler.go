package tutor

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CreateTutorHandler struct {
	TutorService *services.TutorService
	logger       *log.Logger
}

func NewCreateTutorHandler(
	service *services.TutorService,
	logger *log.Logger,
) *CreateTutorHandler {
	return &CreateTutorHandler{
		TutorService: service,
		logger:       logger,
	}
}

type TutorRequest struct {
	// in:body
	Tutor struct {
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
		AcademicRank string `json:"academic_rank"`
	} `json:""`
}

type CreateTutorResponse struct {
	CreatedTutorID uint   `json:"insertedID"`
	ErrorMessage   string `json:"errorMessage,omitempty"`
}

func (handler *CreateTutorHandler) CreateTutorController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &CreateTutorResponse{}
	request := &TutorRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating tutor")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed tutor data"
		json.NewEncoder(w).Encode(response)
		return
	}

	tutorRequest := request.Tutor

	birthDate, err := time.Parse("01-02-2006", tutorRequest.BirthDate)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating tutor birth date")

		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed tutor data: birth date"
		json.NewEncoder(w).Encode(response)
		return
	}

	domainUser := &domain.User{
		Username:    tutorRequest.Username,
		Password:    tutorRequest.Password,
		FirstName:   tutorRequest.FirstName,
		LastName:    tutorRequest.LastName,
		Email:       tutorRequest.Email,
		BirthDate:   birthDate,
		PhoneNumber: tutorRequest.PhoneNumber,
		Photo:       tutorRequest.Photo,
	}

	domainTutor := &domain.Tutor{
		User:         *domainUser,
		AcademicRank: tutorRequest.AcademicRank,
	}

	uid, err := handler.TutorService.CreateTutor(context.Background(), domainTutor)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating tutor in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating tutor"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.CreatedTutorID = uid

	w.WriteHeader(http.StatusCreated)
	return
}
