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

type ChangeUserPasswordHandler struct {
	UserService *services.UserService
	logger      *log.Logger
}

func NewChangeUserPasswordHandler(
	service *services.UserService,
	logger *log.Logger,
) *ChangeUserPasswordHandler {
	return &ChangeUserPasswordHandler{
		UserService: service,
		logger:      logger,
	}
}

type ChangeUserPasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangeUserPasswordResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *ChangeUserPasswordHandler) ChangeUserPasswordController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &ChangeUserPasswordResponse{}

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

	request := &ChangeUserPasswordRequest{}

	err = json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in changing user password")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed user password data"
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.NewPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing new password"
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.OldPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing old password"
		json.NewEncoder(w).Encode(response)
		return

	}

	err = handler.UserService.ChangeUserPassword(context.Background(), uint32(uid), request.OldPassword, request.NewPassword)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in changing user password")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if passwordMismatchError, ok := err.(apierrors.PasswordMismatchError); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": passwordMismatchError.Unwrap().Error(),
		}).Debug("Error in changing user password")

		w.WriteHeader(passwordMismatchError.ReturnedStatusCode)

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
