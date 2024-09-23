package question

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type BulkDeleteQuestionsRequest struct {
	IDs []uint32
}

type BulkDeleteQuestionHandler struct {
	QuestionService *services.QuestionService
	logger          *log.Logger
}

func NewBulkDeleteQuestionHandler(
	service *services.QuestionService,
	logger *log.Logger,
) *BulkDeleteQuestionHandler {
	return &BulkDeleteQuestionHandler{
		QuestionService: service,
		logger:          logger,
	}
}

type BulkDeleteQuestionResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *BulkDeleteQuestionHandler) BulkDeleteQuestionController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &BulkDeleteQuestionResponse{}

	request := &BulkDeleteQuestionsRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in bulk deleting questions")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed questions request"
		json.NewEncoder(w).Encode(response)
		return
	}

	err = handler.QuestionService.BulkDeleteQuestions(context.Background(), request.IDs)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating solar panel data")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
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
