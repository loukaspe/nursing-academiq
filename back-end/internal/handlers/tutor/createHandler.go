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

type Tutor struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	BirthDate    string `json:"birth_date"`
	PhoneNumber  string `json:"phone_number"`
	Photo        string `json:"photo"`
	AcademicRank string `json:"academic_rank"`
}

type CreateTutorResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *CreateTutorHandler) CreateTutorController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &CreateTutorResponse{}
	tutorRequest := &Tutor{}

	err := json.NewDecoder(r.Body).Decode(tutorRequest)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating tutor")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed tutor data"
		json.NewEncoder(w).Encode(response)
		return
	}

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

	uid, err := handler.TutorService.CreateTutor(context.TODO(), domainTutor)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating tutor in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating tutor"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.CreatedTutorUid = uid

	w.WriteHeader(http.StatusCreated)
	return
}
