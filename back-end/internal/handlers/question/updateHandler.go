package question

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type QuestionRequest struct {
	ChapterID              uint
	CourseID               uint
	Text                   string
	Explanation            string
	Source                 string
	MultipleCorrectAnswers bool
	NumberOfAnswers        int
	Answers                []Answer
}

type UpdateQuestionHandler struct {
	QuestionService *services.QuestionService
	logger          *log.Logger
}

func NewUpdateQuestionHandler(
	service *services.QuestionService,
	logger *log.Logger,
) *UpdateQuestionHandler {
	return &UpdateQuestionHandler{
		QuestionService: service,
		logger:          logger,
	}
}

type UpdateQuestionResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateQuestionHandler) UpdateQuestionController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateQuestionResponse{}

	request := &QuestionRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating question")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed question request"
		json.NewEncoder(w).Encode(response)
		return
	}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing question id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed question id"
		json.NewEncoder(w).Encode(response)
		return
	}

	domainAnswers := make([]domain.Answer, 0, len(request.Answers))
	for _, answer := range request.Answers {
		domainAnswers = append(domainAnswers, domain.Answer{
			Text:      answer.Text,
			IsCorrect: answer.IsCorrect,
		})
	}

	domainQuestion := &domain.Question{
		Text:                   request.Text,
		Explanation:            request.Explanation,
		Source:                 request.Source,
		MultipleCorrectAnswers: request.MultipleCorrectAnswers,
		NumberOfAnswers:        request.NumberOfAnswers,
		Answers:                domainAnswers,
		Course:                 &domain.Course{ID: uint32(request.CourseID)},
		Chapter:                &domain.Chapter{ID: uint32(request.ChapterID)},
	}

	err = handler.QuestionService.UpdateQuestion(r.Context(), uint32(uid), domainQuestion)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating question")

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
