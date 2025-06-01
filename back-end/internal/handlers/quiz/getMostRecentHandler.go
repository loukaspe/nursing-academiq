package quiz

import (
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetMostRecentQuizzesHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewGetMostRecentQuizzesHandler(
	service *services.QuizService,
	logger *log.Logger,
) *GetMostRecentQuizzesHandler {
	return &GetMostRecentQuizzesHandler{
		QuizService: service,
		logger:      logger,
	}
}

type GetMostRecentQuizzesResponse struct {
	ErrorMessage string         `json:"errorMessage,omitempty"`
	Quizzes      []QuizResponse `json:"quizzes,omitempty"`
}

func (handler *GetMostRecentQuizzesHandler) GetMostRecentQuizzesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &GetMostRecentQuizzesResponse{}

	var limit int
	var err error

	limitParam := r.URL.Query().Get("limit")
	if limitParam == "" {

	} else {
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.ErrorMessage = "malformed limit"
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	quizzes, err := handler.QuizService.GetMostRecentQuizzes(r.Context(), limit)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting most recent quizzes")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, quiz := range quizzes {
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
				ID:                     domainQuestion.ID,
				Text:                   domainQuestion.Text,
				Explanation:            domainQuestion.Explanation,
				Source:                 domainQuestion.Source,
				MultipleCorrectAnswers: domainQuestion.MultipleCorrectAnswers,
				NumberOfAnswers:        domainQuestion.NumberOfAnswers,
				Answers:                answers,
			}

			questions = append(questions, *question)
		}

		quizResponse := QuizResponse{
			ID:          quiz.ID,
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

		response.Quizzes = append(response.Quizzes, quizResponse)
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
