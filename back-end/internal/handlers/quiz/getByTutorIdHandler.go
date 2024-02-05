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

type GetQuizByTutorIDHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewGetQuizByTutorIDHandler(
	service *services.QuizService,
	logger *log.Logger,
) *GetQuizByTutorIDHandler {
	return &GetQuizByTutorIDHandler{
		QuizService: service,
		logger:      logger,
	}
}

type GetQuizByTutorIDResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Quizzes      []struct {
		Title             string
		Description       string
		Visibility        bool
		ShowSubset        bool
		SubsetSize        int
		NumberOfSessions  int
		ScoreSum          float32
		MaxScore          int
		NumberOfQuestions int
		CourseName        string
	} `json:"quizzes,omitempty"`
}

func (handler *GetQuizByTutorIDHandler) GetQuizByTutorIDController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetQuizByTutorIDResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing tutor id"
		json.NewEncoder(w).Encode(response)
		return
	}
	tutorId, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed tutor id"
		json.NewEncoder(w).Encode(response)
		return
	}

	quizs, err := handler.QuizService.GetQuizByTutorID(context.TODO(), uint32(tutorId))
	if dataNotFoundErrorWrapper, ok := err.(*apierrors.DataNotFoundErrorWrapper); ok {
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

	for _, quiz := range quizs {
		response.Quizzes = append(response.Quizzes, struct {
			Title             string
			Description       string
			Visibility        bool
			ShowSubset        bool
			SubsetSize        int
			NumberOfSessions  int
			ScoreSum          float32
			MaxScore          int
			NumberOfQuestions int
			CourseName        string
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
			CourseName:        quiz.Course.Title,
		})
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
