package handlers

import (
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type LoginHandler struct {
	jwtService   *services.JwtService
	loginService *services.LoginService
	logger       *log.Logger
}

func NewLoginHandler(
	jwtService *services.JwtService,
	loginService *services.LoginService,
	logger *log.Logger,
) *LoginHandler {
	return &LoginHandler{
		jwtService:   jwtService,
		loginService: loginService,
		logger:       logger,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *LoginHandler) LoginController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request LoginRequest
	var response LoginResponse

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handler.logger.Error("Error in login - request decode",
			map[string]interface{}{
				"errorMessage": err.Error(),
			})

		response.ErrorMessage = "malformed auth request"

		handler.JsonResponse(w, http.StatusInternalServerError, &response)

		return
	}

	if request.Username == "" || request.Password == "" {
		response.ErrorMessage = "empty username or password"

		handler.JsonResponse(w, http.StatusBadRequest, &response)

		return
	}

	loginUserResponse, userID, err := handler.loginService.Login(r.Context(), request.Username, request.Password)
	if loginErrorWrapper, ok := err.(*apierrors.LoginError); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": loginErrorWrapper.Unwrap().Error(),
		}).Debug("Error in login")

		response.ErrorMessage = "invalid credentials"
		handler.JsonResponse(w, loginErrorWrapper.ReturnedStatusCode, &response)

		return
	}
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Debug("Error in login")
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	jwtSubject := &domain.JwtSubject{
		User:   loginUserResponse,
		UserID: userID,
	}

	accessToken, err := handler.jwtService.CreateAccessJwtToken(jwtSubject)
	if err != nil {
		response.ErrorMessage = "error during creation of the access token"

		handler.JsonResponse(w, http.StatusInternalServerError, &response)

		return
	}

	refreshToken, err := handler.jwtService.CreateRefreshJwtToken(jwtSubject)
	if err != nil {
		response.ErrorMessage = "error during creation of the refresh token"

		handler.JsonResponse(w, http.StatusInternalServerError, &response)

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})

	response.AccessToken = accessToken
	handler.JsonResponse(w, http.StatusOK, &response)

	return
}

func (handler *LoginHandler) JsonResponse(
	w http.ResponseWriter,
	statusCode int,
	response *LoginResponse,
) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error in logging in - json response"

		handler.logger.Error("Error in logging in - json response",
			map[string]interface{}{
				"errorMessage": err.Error(),
			})
	}
}
