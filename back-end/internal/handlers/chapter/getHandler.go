package chapter

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetChapterHandler struct {
	ChapterService *services.ChapterService
	logger         *log.Logger
}

type Quiz struct {
	ID                uint32
	Title             string
	NumberOfQuestions int
	CourseName        string
}

type Course struct {
	ID    uint32
	Title string
}

type Chapter struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Course      Course `json:"course"`
	Quizzes     []Quiz `json:"quizzes,omitempty"`
}

func NewGetChapterHandler(
	service *services.ChapterService,
	logger *log.Logger,
) *GetChapterHandler {
	return &GetChapterHandler{
		ChapterService: service,
		logger:         logger,
	}
}

type GetChapterResponse struct {
	ErrorMessage string   `json:"errorMessage,omitempty"`
	Chapter      *Chapter `json:"chapter,omitempty"`
}

func (handler *GetChapterHandler) GetChapterController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetChapterResponse{
		Chapter: &Chapter{},
	}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing chapter id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed chapter id"
		json.NewEncoder(w).Encode(response)
		return
	}

	chapter, err := handler.ChapterService.GetChapter(r.Context(), uint32(uid))
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

	for _, quiz := range chapter.Quizzes {
		response.Chapter.Quizzes = append(response.Chapter.Quizzes, Quiz{
			ID:                quiz.ID,
			Title:             quiz.Title,
			CourseName:        quiz.Course.Title,
			NumberOfQuestions: len(quiz.Questions),
		})
	}

	response.Chapter.Title = chapter.Title
	response.Chapter.Description = chapter.Description
	response.Chapter.Course = Course{
		ID:    chapter.Course.ID,
		Title: chapter.Course.Title,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
