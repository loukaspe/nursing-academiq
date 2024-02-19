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

type GetQuizHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewGetQuizHandler(
	service *services.QuizService,
	logger *log.Logger,
) *GetQuizHandler {
	return &GetQuizHandler{
		QuizService: service,
		logger:      logger,
	}
}

type GetQuizResponse struct {
	ErrorMessage string       `json:"errorMessage,omitempty"`
	Quiz         *QuizRequest `json:"quiz,omitempty"`
}

func (handler *GetQuizHandler) GetQuizController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetQuizResponse{}

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

	quiz, err := handler.QuizService.GetQuiz(context.TODO(), uint32(uid))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting user data")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Quiz = &QuizRequest{
		Quiz: struct {
			Title             string
			Description       string
			Visibility        bool
			ShowSubset        bool
			SubsetSize        int
			NumberOfSessions  int
			ScoreSum          float32
			MaxScore          int
			NumberOfQuestions int
		}{
			Title:             quiz.Title,
			Description:       quiz.Description,
			Visibility:        quiz.Visibility,
			ShowSubset:        quiz.ShowSubset,
			SubsetSize:        quiz.SubsetSize,
			NumberOfSessions:  quiz.NumberOfSessions,
			ScoreSum:          quiz.ScoreSum,
			MaxScore:          quiz.MaxScore,
			NumberOfQuestions: quiz.NumberOfQuestions,
		},
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
