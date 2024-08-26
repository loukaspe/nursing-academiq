package chapter

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

type DeleteChapterHandler struct {
	ChapterService *services.ChapterService
	logger         *log.Logger
}

func NewDeleteChapterHandler(
	service *services.ChapterService,
	logger *log.Logger,
) *DeleteChapterHandler {
	return &DeleteChapterHandler{
		ChapterService: service,
		logger:         logger,
	}
}

type DeleteChapterResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *DeleteChapterHandler) DeleteChapterController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &DeleteChapterResponse{}

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

	err = handler.ChapterService.DeleteChapter(context.Background(), uint32(uid))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in deleting chapter")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
