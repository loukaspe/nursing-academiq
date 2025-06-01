package question

import (
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

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

	uid, err := handler.QuestionService.CreateQuestion(r.Context(), domainQuestion)
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
