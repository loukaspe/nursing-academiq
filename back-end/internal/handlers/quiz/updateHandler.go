package quiz

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UpdateQuizRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Visibility  bool   `json:"visibility"`
	ShowSubset  bool   `json:"showSubset"`
	SubsetSize  int    `json:"subsetSize"`
}

type UpdateQuizHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewUpdateQuizHandler(
	service *services.QuizService,
	logger *log.Logger,
) *UpdateQuizHandler {
	return &UpdateQuizHandler{
		QuizService: service,
		logger:      logger,
	}
}

type UpdateQuizResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateQuizHandler) UpdateQuizController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateQuizResponse{}

	request := &UpdateQuizRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating quiz")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed quiz request"
		json.NewEncoder(w).Encode(response)
		return
	}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing quiz id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed quiz id"
		json.NewEncoder(w).Encode(response)
		return
	}

	domainQuiz := &domain.Quiz{
		Title:       request.Title,
		Description: request.Description,
		Visibility:  request.Visibility,
		ShowSubset:  request.ShowSubset,
		SubsetSize:  request.SubsetSize,
	}

	err = handler.QuizService.UpdateQuiz(context.Background(), uint32(uid), domainQuiz)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating quiz")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

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
