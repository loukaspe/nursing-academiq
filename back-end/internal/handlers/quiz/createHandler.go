package quiz

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type QuizRequest struct {
	Title       string
	Description string
	Visibility  bool
	ShowSubset  bool
	SubsetSize  int
	//NumberOfSessions  int
	ScoreSum          float32
	MaxScore          int
	NumberOfQuestions int
	CourseID          uint
}

type CreateQuizResponse struct {
	CreatedQuizID uint   `json:"insertedID"`
	ErrorMessage  string `json:"errorMessage,omitempty"`
}

type CreateQuizHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewCreateQuizHandler(
	service *services.QuizService,
	logger *log.Logger,
) *CreateQuizHandler {
	return &CreateQuizHandler{
		QuizService: service,
		logger:      logger,
	}
}

func (handler *CreateQuizHandler) CreateQuizController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &CreateQuizResponse{}
	request := &QuizRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating quiz")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed quiz data"
		json.NewEncoder(w).Encode(response)
		return
	}

	domainQuiz := &domain.Quiz{
		Title:       request.Title,
		Description: request.Description,
		Visibility:  request.Visibility,
		ShowSubset:  request.ShowSubset,
		SubsetSize:  request.SubsetSize,
		ScoreSum:    request.ScoreSum,
		MaxScore:    request.MaxScore,
		Course:      &domain.Course{ID: uint32(request.CourseID)},
	}

	uid, err := handler.QuizService.CreateQuiz(context.Background(), domainQuiz)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating quiz in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating quiz"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.CreatedQuizID = uid

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	return
}
