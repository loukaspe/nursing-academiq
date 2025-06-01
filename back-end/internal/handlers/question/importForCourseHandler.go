package question

import (
	"encoding/csv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// Text, Chapter, NumberOfAnswers, Explanation, Source, Answer1, IsCorrect1, Answer2, IsCorrect2, Answer3, IsCorrect3, Answer4, IsCorrect4, Answer5, IsCorrect5, Answer6, IsCorrect6, Answer7, IsCorrect7, Answer8, IsCorrect8, Link
const CSVColumns = 22

var CSVHeaders = []string{"Εκφώνηση", "Κατηγορία", "Πλήθος Απαντήσεων", "Επεξήγηση Λύσης", "Πηγή", "Απάντηση 1", "Ορθότητα 1", "Απάντηση 2", "Ορθότητα 2", "Απάντηση 3", "Ορθότητα 3", "Απάντηση 4", "Ορθότητα 4", "Απάντηση 5", "Ορθότητα 5", "Απάντηση 6", "Ορθότητα 6", "Απάντηση 7", "Ορθότητα 7", "Απάντηση 8", "Ορθότητα 8", "Εικόνα (Σύνδεσμος/Αρχείο)"}

type ImportQuestionHandler struct {
	questionService *services.QuestionService
	chapterService  *services.ChapterService
	logger          *log.Logger
}

func NewImportQuestionHandler(
	questionService *services.QuestionService,
	chapterService *services.ChapterService,
	logger *log.Logger,
) *ImportQuestionHandler {
	return &ImportQuestionHandler{
		questionService: questionService,
		chapterService:  chapterService,
		logger:          logger,
	}
}

type ImportQuestionRequest struct {
	CreateNewChapters bool `json:"create_new_chapters"`
}

func (handler *ImportQuestionHandler) ImportQuestionController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var problematicRecords [][]string

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

	requestBody := r.FormValue("jsonData")
	var request ImportQuestionRequest
	err = json.Unmarshal([]byte(requestBody), &request)
	if err != nil {
		http.Error(w, "Could not parse JSON data", http.StatusBadRequest)
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

		if len(record) < CSVColumns {
			problematicRecords = append(problematicRecords, record)
			continue
		}

		if record[0] == "" || record[1] == "" || record[2] == "" || record[3] == "" || record[4] == "" || record[5] == "" || record[6] == "" {
			problematicRecords = append(problematicRecords, record)
			continue
		}

		domainQuestion := domain.Question{}

		domainQuestion.Text = record[0]

		domainQuestion.Chapter = &domain.Chapter{}
		domainQuestion.Chapter.Title = strings.TrimSpace(record[1])

		if request.CreateNewChapters == false {
			_, err = handler.chapterService.GetChapterByTitle(r.Context(), domainQuestion.Chapter.Title)
			if _, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
				problematicRecords = append(problematicRecords, record)
				continue
			}
		}

		domainQuestion.NumberOfAnswers, err = strconv.Atoi(record[2])
		if err != nil {
			problematicRecords = append(problematicRecords, record)
			continue
		}

		domainQuestion.Explanation = record[3]
		domainQuestion.Source = record[4]

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

	err = handler.questionService.ImportForCourse(r.Context(), domainQuestions, uint(courseID))
	if err != nil {
		handler.logger.Error("Error importing questions", err)
		http.Error(w, "Error importing questions", http.StatusInternalServerError)
		return
	}

	// If there are problematic rows, return them as a CSV
	if len(problematicRecords) > 0 {
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", "attachment; filename=\"problematic_questions.csv\"")
		w.WriteHeader(http.StatusOK)
		csvWriter := csv.NewWriter(w)
		defer csvWriter.Flush()

		// Write header
		csvWriter.Write(CSVHeaders)

		// Write problematic rows
		for _, row := range problematicRecords {
			csvWriter.Write(row)
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
	return
}
