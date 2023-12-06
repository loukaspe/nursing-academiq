package handlers

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type JwtClaimsHandler struct {
	jwtService   *services.JwtService
	loginService *services.LoginService
	logger       *log.Logger
}

func NewJwtClaimsHandler(
	jwtService *services.JwtService,
	loginService *services.LoginService,
	logger *log.Logger,
) *JwtClaimsHandler {
	return &JwtClaimsHandler{
		jwtService:   jwtService,
		loginService: loginService,
		logger:       logger,
	}
}

// request for generating jwt token
//
// swagger:parameters jwtToken
type JwtRequest struct {
	// in:body
	// Required: true
	Username string `json:"username"`
	// in:body
	// Required: true
	Password string `json:"password"`
}

// Response with jwtToken
// swagger:model JwtResponse
type JwtResponse struct {
	// jwt token
	//
	// Required: false
	Token string `json:"token"`
	// possible error message
	//
	// Required: false
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// swagger:operation POST /login jwtToken
//
// # Generates JWT token for authentication and authorization
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
//	responses:
//		"200":
//			description: OK
//			schema:
//				$ref: "#/definitions/JwtResponse"
//		"500":
//			description: Error
//			schema:
//				$ref: "#/definitions/JwtResponse"
func (handler *JwtClaimsHandler) JwtTokenController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request JwtRequest
	var response JwtResponse

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handler.logger.Error("Error in creating jwt token - request decode",
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

	loginUserResponse, err := handler.loginService.Login(context.Background(), request.Username, request.Password)
	if loginErrorWrapper, ok := err.(*apierrors.LoginError); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": loginErrorWrapper.Unwrap().Error(),
		}).Debug("Error in login")

		w.WriteHeader(loginErrorWrapper.ReturnedStatusCode)

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

	result, err := handler.jwtService.CreateJwtTokenService(loginUserResponse)
	if err != nil {
		response.ErrorMessage = "error during creation of the token"

		handler.JsonResponse(w, http.StatusInternalServerError, &response)

		return
	}

	response.Token = result
	handler.JsonResponse(w, http.StatusOK, &response)

	return
}

func (handler *JwtClaimsHandler) JsonResponse(
	w http.ResponseWriter,
	statusCode int,
	response *JwtResponse,
) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error in adding user favourite asset - json response"

		handler.logger.Error("Error in adding user favourite asset - json response",
			map[string]interface{}{
				"errorMessage": err.Error(),
			})
	}
}
