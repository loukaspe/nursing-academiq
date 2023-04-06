package user

import (
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CreateUserHandler struct {
	UserService *services.UserService
	logger      *log.Logger
}

func NewCreateUserHandler(
	service *services.UserService,
	logger *log.Logger,
) *CreateUserHandler {
	return &CreateUserHandler{
		UserService: service,
		logger:      logger,
	}
}

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	BirthDate   string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Photo       string `json:"photo"`
}

type CreateUserResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *CreateUserHandler) CreateUserController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &CreateUserResponse{}
	userRequest := &User{}

	err := json.NewDecoder(r.Body).Decode(userRequest)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating user")

		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed user data"
		json.NewEncoder(w).Encode(response)
		return
	}

	birthDate, err := time.Parse("01-02-2006", userRequest.BirthDate)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating user birth date")

		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed user data: birth date"
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

	err = handler.UserService.CreateUser(domainUser)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating user in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating user"
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
