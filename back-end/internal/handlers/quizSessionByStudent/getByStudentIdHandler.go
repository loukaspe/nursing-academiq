package quizSessionByStudent

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

type GetQuizSessionByStudentIDHandler struct {
	QuizSessionByStudentService *services.QuizSessionByStudentService
	logger                      *log.Logger
}

func NewGetQuizSessionByStudentIDHandler(
	service *services.QuizSessionByStudentService,
	logger *log.Logger,
) *GetQuizSessionByStudentIDHandler {
	return &GetQuizSessionByStudentIDHandler{
		QuizSessionByStudentService: service,
		logger:                      logger,
	}
}

type GetQuizSessionByStudentIDResponse struct {
	ErrorMessage          string `json:"errorMessage,omitempty"`
	QuizSessionsByStudent []struct {
		QuizName string  `json:"quizName"`
		Date     string  `json:"date"`
		Duration int     `json:"duration"`
		Score    float32 `json:"score"`
		MaxScore int     `json:"maxScore"`
	} `json:"quizSessions,omitempty"`
}

func (handler *GetQuizSessionByStudentIDHandler) GetQuizSessionByStudentIDController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetQuizSessionByStudentIDResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing student id"
		json.NewEncoder(w).Encode(response)
		return
	}
	studentId, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed student id"
		json.NewEncoder(w).Encode(response)
		return
	}

	QuizSessionsByStudent, err := handler.QuizSessionByStudentService.GetQuizSessionByStudentID(context.Background(), uint32(studentId))
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

	for _, quizSessionByStudent := range QuizSessionsByStudent {
		response.QuizSessionsByStudent = append(response.QuizSessionsByStudent, struct {
			QuizName string  `json:"quizName"`
			Date     string  `json:"date"`
			Duration int     `json:"duration"`
			Score    float32 `json:"score"`
			MaxScore int     `json:"maxScore"`
		}{
			QuizName: quizSessionByStudent.Quiz.Title,
			Date:     quizSessionByStudent.Date.Format("02/01/2006"),
			Duration: quizSessionByStudent.DurationInSeconds,
			Score:    quizSessionByStudent.Score,
			MaxScore: quizSessionByStudent.MaxScore,
		})
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
