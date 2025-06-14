package tutor

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetTutorHandler struct {
	TutorService *services.TutorService
	logger       *log.Logger
}

func NewGetTutorHandler(
	service *services.TutorService,
	logger *log.Logger,
) *GetTutorHandler {
	return &GetTutorHandler{
		TutorService: service,
		logger:       logger,
	}
}

type TutorRequest struct {
	Tutor struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		BirthDate    string `json:"birth_date"`
		PhoneNumber  string `json:"phone_number"`
		AcademicRank string `json:"academic_rank"`
	} `json:""`
}

type GetTutorResponse struct {
	ErrorMessage string        `json:"errorMessage,omitempty"`
	Tutor        *TutorRequest `json:"tutor,omitempty"`
}

func (handler *GetTutorHandler) GetTutorController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetTutorResponse{}

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

	tutor, err := handler.TutorService.GetTutor(r.Context(), uint32(uid))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
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

	response.Tutor = &TutorRequest{
		Tutor: struct {
			Username     string `json:"username"`
			Password     string `json:"password"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Email        string `json:"email"`
			BirthDate    string `json:"birth_date"`
			PhoneNumber  string `json:"phone_number"`
			AcademicRank string `json:"academic_rank"`
		}{
			Username:     tutor.Username,
			Password:     tutor.Password,
			FirstName:    tutor.FirstName,
			LastName:     tutor.LastName,
			Email:        tutor.Email,
			BirthDate:    tutor.BirthDate.String(),
			PhoneNumber:  tutor.PhoneNumber,
			AcademicRank: tutor.AcademicRank,
		},
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
