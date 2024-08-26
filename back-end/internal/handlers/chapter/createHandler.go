package chapter

import (
	"context"
	"encoding/json"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type ChapterRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    uint   `json:"courseID"`
}

type CreateChapterResponse struct {
	CreatedChapterID uint   `json:"insertedID"`
	ErrorMessage     string `json:"errorMessage,omitempty"`
}

type CreateChapterHandler struct {
	ChapterService *services.ChapterService
	logger         *log.Logger
}

func NewCreateChapterHandler(
	service *services.ChapterService,
	logger *log.Logger,
) *CreateChapterHandler {
	return &CreateChapterHandler{
		ChapterService: service,
		logger:         logger,
	}
}

func (handler *CreateChapterHandler) CreateChapterController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &CreateChapterResponse{}
	request := &ChapterRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating chapter")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed chapter data"
		json.NewEncoder(w).Encode(response)
		return
	}

	domainChapter := &domain.Chapter{
		Title:       request.Title,
		Description: request.Description,
	}

	uid, err := handler.ChapterService.CreateChapter(context.Background(), domainChapter, request.CourseID)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in creating chapter in service")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "error creating chapter"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.CreatedChapterID = uid

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusCreated)
	return
}
