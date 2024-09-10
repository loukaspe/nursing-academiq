package question

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type QuestionRequest struct {
	Question Question `json:""`
}

type CreateQuestionResponse struct {
	CreatedQuestionID uint   `json:"insertedID"`
	ErrorMessage      string `json:"errorMessage,omitempty"`
}

type CreateQuestionHandler struct {
	QuestionService *services.QuestionService
	logger          *log.Logger
}

func NewCreateQuestionHandler(
	service *services.QuestionService,
	logger *log.Logger,
) *CreateQuestionHandler {
	return &CreateQuestionHandler{
		QuestionService: service,
		logger:          logger,
	}
}

func (handler *CreateQuestionHandler) CreateQuestionController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &CreateQuestionResponse{}
	request := &QuestionRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating question")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed question data"
		json.NewEncoder(w).Encode(response)
		return
	}

	questionRequest := request.Question

	domainAnswers := make([]domain.Answer, 0, len(questionRequest.Answers))
	for _, answer := range questionRequest.Answers {
		domainAnswers = append(domainAnswers, domain.Answer{
			Text:      answer.Text,
			IsCorrect: answer.IsCorrect,
		})
	}

	domainQuestion := &domain.Question{
		Text:                   questionRequest.Text,
		Explanation:            questionRequest.Explanation,
		Source:                 questionRequest.Source,
		MultipleCorrectAnswers: questionRequest.MultipleCorrectAnswers,
		NumberOfAnswers:        questionRequest.NumberOfAnswers,
		Answers:                domainAnswers,
		Course:                 &domain.Course{ID: uint32(questionRequest.CourseID)},
	}

	uid, err := handler.QuestionService.CreateQuestion(context.Background(), domainQuestion)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating question in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating question"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.CreatedQuestionID = uid

	w.WriteHeader(http.StatusCreated)
	return
}
