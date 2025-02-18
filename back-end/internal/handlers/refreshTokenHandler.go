package handlers

import (
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type RefreshTokenHandler struct {
	jwtService  *services.JwtService
	userService *services.UserService
	logger      *log.Logger
}

func NewRefreshTokenHandler(
	jwtService *services.JwtService,
	userService *services.UserService,
	logger *log.Logger,
) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		jwtService:  jwtService,
		userService: userService,
		logger:      logger,
	}
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *RefreshTokenHandler) RefreshTokenController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		http.Error(w, "Refresh token not found", http.StatusUnauthorized)
		return
	}

	refreshToken := cookie.Value
	if refreshToken == "" {
		http.Error(w, "Refresh token empty", http.StatusUnauthorized)
		return
	}

	claims, err := handler.jwtService.RefreshClaimsFromJwtToken(refreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		http.Error(w, "Expired refresh token", http.StatusUnauthorized)
		return
	}

	username := claims["sub"].(string)
	if username == "" {
		http.Error(w, "Invalid refresh token username", http.StatusUnauthorized)
		return
	}

	user, err := handler.userService.GetUserByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, "Failed to get user to authorize token", http.StatusInternalServerError)
		return
	}

	jwtSubject := &domain.JwtSubject{
		User:   user,
		UserID: user.ID,
	}

	newAccessToken, err := handler.jwtService.CreateAccessJwtToken(jwtSubject)
	if err != nil {
		http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
		return
	}

	response := RefreshTokenResponse{
		AccessToken: newAccessToken,
	}
	handler.JsonResponse(w, http.StatusOK, &response)

	return
}

func (handler *RefreshTokenHandler) JsonResponse(
	w http.ResponseWriter,
	statusCode int,
	response *RefreshTokenResponse,
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
