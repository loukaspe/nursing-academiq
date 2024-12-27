package quiz

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

type UpdateQuizQuestionsRequest struct {
	QuestionsIDs []uint32
}

type UpdateQuizQuestionsHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewUpdateQuizQuestionsHandler(
	service *services.QuizService,
	logger *log.Logger,
) *UpdateQuizQuestionsHandler {
	return &UpdateQuizQuestionsHandler{
		QuizService: service,
		logger:      logger,
	}
}

type UpdateQuizQuestionsResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateQuizQuestionsHandler) UpdateQuizQuestionsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateQuizQuestionsResponse{}

	request := &UpdateQuizQuestionsRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating quiz questions")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed quiz questions request"
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

	err = handler.QuizService.UpdateQuizQuestions(context.Background(), uint32(uid), request.QuestionsIDs)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating quiz questions")

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
