package tutor

import (
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

type UpdateTutorHandler struct {
	TutorService *services.TutorService
	logger       *log.Logger
}

func NewUpdateTutorHandler(
	service *services.TutorService,
	logger *log.Logger,
) *UpdateTutorHandler {
	return &UpdateTutorHandler{
		TutorService: service,
		logger:       logger,
	}
}

type UpdateTutorResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateTutorHandler) UpdateTutorController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateTutorResponse{}

	tutorRequest := &Tutor{}

	err := json.NewDecoder(r.Body).Decode(tutorRequest)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating tutor")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed tutor request"
		json.NewEncoder(w).Encode(response)
		return
	}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing tutor id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed tutor id"
		json.NewEncoder(w).Encode(response)
		return
	}

	birthDate, err := time.Parse("17-03-2023", tutorRequest.BirthDate)
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

	err = handler.TutorService.UpdateTutor(uint32(uid), domainTutor)
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
