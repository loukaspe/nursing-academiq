package quiz

import (
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type SearchQuizHandler struct {
	QuizService *services.QuizService
	logger      *log.Logger
}

func NewSearchQuizHandler(
	service *services.QuizService,
	logger *log.Logger,
) *SearchQuizHandler {
	return &SearchQuizHandler{
		QuizService: service,
		logger:      logger,
	}
}

type SearchQuizRequest struct {
	Title      string `json:"title"`
	CourseName string `json:"courseName,omitempty"`
}

type SearchQuizResponse struct {
	ErrorMessage string         `json:"errorMessage,omitempty"`
	Quizzes      []QuizResponse `json:"quizzes,omitempty"`
}

func (handler *SearchQuizHandler) SearchQuizController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &SearchQuizResponse{}

	request := &SearchQuizRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in searching quiz")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed search quiz request"
		json.NewEncoder(w).Encode(response)
		return
	}

	quizzes, err := handler.QuizService.SearchQuiz(r.Context(), request.Title, request.CourseName)
	if _, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		w.WriteHeader(http.StatusNoContent)
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
