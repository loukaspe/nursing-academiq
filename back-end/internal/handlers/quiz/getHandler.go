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
}

type Course struct {
	ID    uint32
	Title string
}

type QuizResponse struct {
	Title             string
	Description       string
	Visibility        bool
	ShowSubset        bool
	SubsetSize        int
	ScoreSum          float32
	MaxScore          int
	NumberOfQuestions int
	Questions         []Question
	Course            Course
}

type GetQuizResponse struct {
	ErrorMessage string        `json:"errorMessage,omitempty"`
	Quiz         *QuizResponse `json:"quiz,omitempty"`
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

	quiz, err := handler.QuizService.GetQuiz(context.Background(), uint32(uid))
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

	questions := make([]Question, 0, len(quiz.Questions))
	for _, domainQuestion := range quiz.Questions {

		answers := make([]Answer, 0, domainQuestion.NumberOfAnswers)
		for _, modelAnswer := range domainQuestion.Answers {
			answer := &Answer{
				Text:      modelAnswer.Text,
				IsCorrect: modelAnswer.IsCorrect,
			}

			answers = append(answers, *answer)
		}

		question := &Question{
			Text:                   domainQuestion.Text,
			Explanation:            domainQuestion.Explanation,
			Source:                 domainQuestion.Source,
			MultipleCorrectAnswers: domainQuestion.MultipleCorrectAnswers,
			NumberOfAnswers:        domainQuestion.NumberOfAnswers,
			Answers:                answers,
		}

		questions = append(questions, *question)
	}

	response.Quiz = &QuizResponse{
		Title:       quiz.Title,
		Description: quiz.Description,
		Visibility:  quiz.Visibility,
		ShowSubset:  quiz.ShowSubset,
		SubsetSize:  quiz.SubsetSize,
		//ScoreSum:          ,
		MaxScore:          quiz.MaxScore,
		NumberOfQuestions: quiz.NumberOfQuestions,
		Questions:         questions,
		Course: Course{
			ID:    quiz.Course.ID,
			Title: quiz.Course.Title,
		},
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
