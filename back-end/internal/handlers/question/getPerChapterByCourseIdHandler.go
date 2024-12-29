package question

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

type GetQuestionByCourseIDHandler struct {
	QuestionService *services.QuestionService
	logger          *log.Logger
}

func NewGetQuestionByCourseIDHandler(
	service *services.QuestionService,
	logger *log.Logger,
) *GetQuestionByCourseIDHandler {
	return &GetQuestionByCourseIDHandler{
		QuestionService: service,
		logger:          logger,
	}
}

type Chapter struct {
	ID        uint32     `json:"id"`
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}

type GetQuestionByCourseIDResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
	Course       Course `json:"course,omitempty"`
}

func (handler *GetQuestionByCourseIDHandler) GetQuestionByCourseIDController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetQuestionByCourseIDResponse{}

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

	domainCourse, err := handler.QuestionService.GetChapterAndQuestionsByCourseID(context.Background(), uint32(courseID))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting domainQuestions data")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	course := Course{
		ID:                domainCourse.ID,
		Title:             domainCourse.Title,
		NumberOfQuestions: domainCourse.NumberOfQuestions,
	}

	for _, domainChapter := range domainCourse.Chapters {
		chapter := Chapter{
			ID:    domainChapter.ID,
			Title: domainChapter.Title,
		}
		questions := make([]Question, 0, len(domainChapter.Questions))
		for _, domainQuestion := range domainChapter.Questions {
			question := Question{
				ID:                     domainQuestion.ID,
				Text:                   domainQuestion.Text,
				Explanation:            domainQuestion.Explanation,
				Source:                 domainQuestion.Source,
				MultipleCorrectAnswers: domainQuestion.MultipleCorrectAnswers,
				NumberOfAnswers:        domainQuestion.NumberOfAnswers,
				ChapterID:              uint(domainChapter.ID),
				CourseID:               uint(domainCourse.ID),
				Chapter:                Chapter{Title: domainChapter.Title},
			}
			questions = append(questions, question)
		}
		chapter.Questions = questions
		course.Chapters = append(course.Chapters, chapter)
	}

	response.Course = course

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
