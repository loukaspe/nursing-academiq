package user

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

type GetUserPhotoHandler struct {
	UserService *services.UserService
	logger      *log.Logger
}

func NewGetUserPhotoHandler(
	service *services.UserService,
	logger *log.Logger,
) *GetUserPhotoHandler {
	return &GetUserPhotoHandler{
		UserService: service,
		logger:      logger,
	}
}

type GetUserPhotoResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Photo        string `json:"path"`
}

func (handler *GetUserPhotoHandler) GetUserPhotoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetUserPhotoResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing user id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed user id"
		json.NewEncoder(w).Encode(response)
		return
	}

	photo, err := handler.UserService.GetUserPhoto(context.Background(), uint32(uid))
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

	response.Photo = "/uploads/" + photo

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
