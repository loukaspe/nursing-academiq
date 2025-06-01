package course

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type GetCourseChaptersHandler struct {
	CourseService *services.CourseService
	logger        *log.Logger
}

func NewGetCourseChaptersHandler(
	service *services.CourseService,
	logger *log.Logger,
) *GetCourseChaptersHandler {
	return &GetCourseChaptersHandler{
		CourseService: service,
		logger:        logger,
	}
}

type GetCourseChaptersResponse struct {
	ErrorMessage string    `json:"errorMessage,omitempty"`
	Title        string    `json:"title"`
	Chapters     []Chapter `json:"chapters,omitempty"`
}

func (handler *GetCourseChaptersHandler) GetCourseChaptersController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//var err error
	response := &GetCourseChaptersResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing course id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed course id"
		json.NewEncoder(w).Encode(response)
		return
	}

	courseWithChapters, err := handler.CourseService.GetCourseChapters(r.Context(), uint32(uid))
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in getting extended course data")

		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.ErrorMessage = err.Error()
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, chapter := range courseWithChapters.Chapters {
		response.Chapters = append(response.Chapters, struct {
			ID          uint32
			Title       string
			Description string
		}{
			ID:          chapter.ID,
			Title:       chapter.Title,
			Description: chapter.Description,
		})
	}

	response.Title = courseWithChapters.Title

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}
