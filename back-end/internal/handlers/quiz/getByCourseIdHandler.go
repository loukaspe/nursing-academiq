package quiz

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetQuizByCourseIDHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewGetQuizByCourseIDHandler(
	service *services.QuizService,
	logger *log.Logger,
) *GetQuizByCourseIDHandler {
	return &GetQuizByCourseIDHandler{
		QuizService: service,
		logger:      logger,
	}
}

type GetQuizByCourseIDResponse struct {
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

func (handler *GetQuizByCourseIDHandler) GetQuizByCourseIDController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &GetQuizByCourseIDResponse{}

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

	quizs, err := handler.QuizService.GetQuizByCourseID(r.Context(), uint32(courseID))
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
