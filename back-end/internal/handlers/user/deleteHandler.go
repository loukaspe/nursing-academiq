package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	"net/http"
	"strconv"
)

type DeleteUserHandler struct {
	UserService *services.UserService
}

func NewDeleteUserHandler(service *services.UserService) *DeleteUserHandler {
	return &DeleteUserHandler{
		UserService: service,
	}
}

type DeleteUserResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *DeleteUserHandler) DeleteUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &DeleteUserResponse{}

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

	err = handler.UserService.DeleteUser(uint32(uid))
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
