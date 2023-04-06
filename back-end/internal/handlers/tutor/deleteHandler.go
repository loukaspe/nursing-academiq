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

type DeleteTutorHandler struct {
	TutorService *services.TutorService
	logger       *log.Logger
}

func NewDeleteTutorHandler(
	service *services.TutorService,
	logger *log.Logger,
) *DeleteTutorHandler {
	return &DeleteTutorHandler{
		TutorService: service,
		logger:       logger,
	}
}

type DeleteTutorResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *DeleteTutorHandler) DeleteTutorController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &DeleteTutorResponse{}

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

	err = handler.TutorService.DeleteTutor(uint32(uid))
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

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
