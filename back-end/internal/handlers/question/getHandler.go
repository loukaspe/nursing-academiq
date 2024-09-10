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

type GetQuestionHandler struct {
	QuestionService *services.QuestionService
	logger          *log.Logger
}

func NewGetQuestionHandler(
	service *services.QuestionService,
	logger *log.Logger,
) *GetQuestionHandler {
	return &GetQuestionHandler{
		QuestionService: service,
		logger:          logger,
	}
}

type Answer struct {
	Text      string
	IsCorrect bool
}

type Question struct {
	Text                   string
	Explanation            string
	Source                 string
	MultipleCorrectAnswers bool
	NumberOfAnswers        int
	Answers                []Answer
	ChapterID              uint
	CourseID               uint
}

type Course struct {
	ID    uint32
	Title string
}

type GetQuestionResponse struct {
	ErrorMessage string    `json:"errorMessage,omitempty"`
	Question     *Question `json:"question,omitempty"`
}

func (handler *GetQuestionHandler) GetQuestionController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetQuestionResponse{}

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

	domainQuestion, err := handler.QuestionService.GetQuestion(context.Background(), uint32(uid))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting questions data")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	domainAnswers := make([]Answer, 0, domainQuestion.NumberOfAnswers)
	for _, modelAnswer := range domainQuestion.Answers {
		answer := &Answer{
			Text:      modelAnswer.Text,
			IsCorrect: modelAnswer.IsCorrect,
		}

		domainAnswers = append(domainAnswers, *answer)
	}

	response.Question = &Question{
		Text:                   domainQuestion.Text,
		Explanation:            domainQuestion.Explanation,
		Source:                 domainQuestion.Source,
		MultipleCorrectAnswers: domainQuestion.MultipleCorrectAnswers,
		NumberOfAnswers:        domainQuestion.NumberOfAnswers,
		Answers:                domainAnswers,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
