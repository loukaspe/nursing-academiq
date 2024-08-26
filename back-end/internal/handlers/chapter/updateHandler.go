package chapter

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/domain"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UpdateChapterRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateChapterHandler struct {
	ChapterService *services.ChapterService
	logger         *log.Logger
}

func NewUpdateChapterHandler(
	service *services.ChapterService,
	logger *log.Logger,
) *UpdateChapterHandler {
	return &UpdateChapterHandler{
		ChapterService: service,
		logger:         logger,
	}
}

type UpdateChapterResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *UpdateChapterHandler) UpdateChapterController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &UpdateChapterResponse{}

	request := &UpdateChapterRequest{}

	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		handler.logger.WithFields(log.Fields{
			"errorMessage": err.Error(),
		}).Error("Error in updating chapter")

		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = "malformed chapter request"
		json.NewEncoder(w).Encode(response)
		return
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

	domainChapter := &domain.Chapter{
		Title:       *request.Title,
		Description: *request.Description,
	}

	err = handler.ChapterService.UpdateChapter(context.Background(), uint32(uid), domainChapter)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating chapter")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
