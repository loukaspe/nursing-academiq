package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	"net/http"
	"strconv"
	"time"
)

type UpdateUserHandler struct {
	UserService *services.UserService
}

func NewUpdateUserHandler(service *services.UserService) *UpdateUserHandler {
	return &UpdateUserHandler{
		UserService: service,
	}
}

type UpdateUserResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateUserHandler) UpdateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateUserResponse{}

	userRequest := &User{}

	err := json.NewDecoder(r.Body).Decode(userRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

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

	birthDate, err := time.Parse("17-03-2023", userRequest.BirthDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	domainUser := &domain.User{
		Username:    userRequest.Username,
		Password:    userRequest.Password,
		FirstName:   userRequest.FirstName,
		LastName:    userRequest.LastName,
		Email:       userRequest.Email,
		BirthDate:   birthDate,
		PhoneNumber: userRequest.PhoneNumber,
		Photo:       userRequest.Photo,
	}

	err = handler.UserService.UpdateUser(uint32(uid), domainUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
