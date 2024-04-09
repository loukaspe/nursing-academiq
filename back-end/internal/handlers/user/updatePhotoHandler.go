package user

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type SetUserPhotoHandler struct {
	UserService     *services.UserService
	logger          *log.Logger
	photosDirectory string
}

func NewSetUserPhotoHandler(
	service *services.UserService,
	logger *log.Logger,
	photosDirectory string,
) *SetUserPhotoHandler {
	return &SetUserPhotoHandler{
		UserService:     service,
		logger:          logger,
		photosDirectory: photosDirectory,
	}
}

type SetUserPhotoResponse struct {
	ErrorMessage string `json:"errorMessage,omitempty"`
}

func (handler *SetUserPhotoHandler) SetUserPhotoController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := &SetUserPhotoResponse{}

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "missing user id"
		json.NewEncoder(w).Encode(response)
		return
	}
	uid, err := strconv.Atoi(id)
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = "malformed user id"
		json.NewEncoder(w).Encode(response)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, fileHandler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = handler.saveFileToDisk(file, fileHandler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.UserService.SetUserPhoto(context.Background(), uint32(uid), fileHandler.Filename)
	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
		handler.logger.WithFields(log.Fields{
			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
		}).Debug("Error in updating user photo")

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

func (handler *SetUserPhotoHandler) saveFileToDisk(file multipart.File, filename string) error {
	filePath := handler.photosDirectory + filename
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(f, file); err != nil {
		return err
	}

	return nil
}
