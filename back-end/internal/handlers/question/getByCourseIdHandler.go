package question

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

type GetQuestionByCourseIDHandler struct {
	QuestionService *services.QuestionService
	logger          *log.Logger
}

func NewGetQuestionByCourseIDHandler(
	service *services.QuestionService,
	logger *log.Logger,
) *GetQuestionByCourseIDHandler {
	return &GetQuestionByCourseIDHandler{
		QuestionService: service,
		logger:          logger,
	}
}

type GetQuestionByCourseIDResponse struct {
	ErrorMessage string     `json:"errorMessage,omitempty"`
	Questions    []Question `json:"questions,omitempty"`
}

func (handler *GetQuestionByCourseIDHandler) GetQuestionByCourseIDController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetQuestionByCourseIDResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing Course id"
		json.NewEncoder(w).Encode(response)
		return
	}
	courseID, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed Course id"
		json.NewEncoder(w).Encode(response)
		return
	}

	domainQuestions, err := handler.QuestionService.GetQuestionsByCourseID(context.Background(), uint32(courseID))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting domainQuestions data")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, question := range domainQuestions {
		domainAnswers := make([]Answer, 0, question.NumberOfAnswers)
		for _, modelAnswer := range question.Answers {
			answer := &Answer{
				Text:      modelAnswer.Text,
				IsCorrect: modelAnswer.IsCorrect,
			}

			domainAnswers = append(domainAnswers, *answer)
		}

		response.Questions = append(response.Questions, Question{
			Text:                   question.Text,
			Explanation:            question.Explanation,
			Source:                 question.Source,
			MultipleCorrectAnswers: question.MultipleCorrectAnswers,
			NumberOfAnswers:        question.NumberOfAnswers,
			Answers:                domainAnswers,
		})
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
