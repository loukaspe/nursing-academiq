package question

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ImportQuestionHandler struct {
	QuestionService *services.QuestionService
	logger          *log.Logger
}

func NewImportQuestionHandler(
	service *services.QuestionService,
	logger *log.Logger,
) *ImportQuestionHandler {
	return &ImportQuestionHandler{
		QuestionService: service,
		logger:          logger,
	}
}

func (handler *ImportQuestionHandler) ImportQuestionController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		http.Error(w, "missing course id", http.StatusBadRequest)
		return
	}
	courseID, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "malformed course id", http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // max 10MB file
	if err != nil {
		handler.logger.Error("Unable to parse form", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		handler.logger.Error("Error retrieving file", err)
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	domainQuestions := make([]domain.Question, 0)
	reader := csv.NewReader(file)
	questionCounter := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Skip header
		if questionCounter == 0 {
			questionCounter++
			continue
		}
		questionCounter++

		if err != nil {
			handler.logger.Error("Error reading csv file", err)
			http.Error(w, "Error reading csv file", http.StatusInternalServerError)
			return
		}

		domainQuestion := domain.Question{}
		domainQuestion.Text = record[0]
		domainQuestion.Chapter = &domain.Chapter{}
		domainQuestion.Chapter.Title = strings.TrimSpace(record[1])
		domainQuestion.Explanation = record[3]
		domainQuestion.Source = record[4]

		domainQuestion.NumberOfAnswers, err = strconv.Atoi(record[2])
		if err != nil {
			handler.logger.Error("Error parsing number of answers", err, questionCounter+1)
			http.Error(w, fmt.Sprintf("Error reading csv file in line %d", questionCounter+1), http.StatusInternalServerError)
			return
		}

		domainAnswers := make([]domain.Answer, 0, domainQuestion.NumberOfAnswers)
		foundCorrectAnswer := false
		for i := 0; i < domainQuestion.NumberOfAnswers; i++ {
			domainAnswer := domain.Answer{
				Text:      record[5+i*2],
				IsCorrect: record[6+i*2] == "1",
			}
			if domainAnswer.IsCorrect && foundCorrectAnswer {
				domainQuestion.MultipleCorrectAnswers = true
			} else if domainAnswer.IsCorrect {
				foundCorrectAnswer = true
			}

			domainAnswers = append(domainAnswers, domainAnswer)
		}
		domainQuestion.Answers = domainAnswers
		domainQuestions = append(domainQuestions, domainQuestion)
	}

	err = handler.QuestionService.ImportForCourse(context.Background(), domainQuestions, uint(courseID))
	if err != nil {
		handler.logger.Error("Error importing questions", err)
		http.Error(w, "Error importing questions", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}
