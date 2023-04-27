package tutor

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

//TODO: test if sql injection

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

// Response when we update Tutor
// swagger:model UpdateTutorResponse
type UpdateTutorResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// swagger:operation PUT /tutor/{tutorId} updateTutor
//
// # It updates a User
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
// responses:
//
//	"200":
//		description: Tutor updates successfully
//		schema:
//			$ref: "#/definitions/UpdateTutorResponse"
//	"400":
//		description: Bad request - request parameters are missing or invalid
//		schema:
//			$ref: "#/definitions/UpdateTutorResponse"
//	"404":
//		description: Requested Tutor not found
//		schema:
//			$ref: "#/definitions/UpdateTutorResponse"
//	"500":
//		description: Internal server error - check logs for details
//		schema:
//			$ref: "#/definitions/UpdateTutorResponse"
func (handler *UpdateTutorHandler) UpdateTutorController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateTutorResponse{}

	request := &TutorRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating tutor")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed tutor request"
		json.NewEncoder(w).Encode(response)
		return
	}

	tutorRequest := request.Tutor

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

	err = handler.TutorService.UpdateTutor(context.TODO(), uint32(uid), domainTutor)
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
