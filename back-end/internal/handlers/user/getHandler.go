package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetUserHandler struct {
	UserService *services.UserService
	logger      *log.Logger
}

func NewGetUserHandler(
	service *services.UserService,
	logger *log.Logger,
) *GetUserHandler {
	return &GetUserHandler{
		UserService: service,
		logger:      logger,
	}
}

type GetUserResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	User         *User  `json:"user,omitempty"`
}

func (handler *GetUserHandler) GetUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetUserResponse{}

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

	user, err := handler.UserService.GetUser(uint32(uid))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	response.User = &User{
		Username:    user.Username,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		BirthDate:   user.BirthDate.String(),
		PhoneNumber: user.PhoneNumber,
		Photo:       user.Photo,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
