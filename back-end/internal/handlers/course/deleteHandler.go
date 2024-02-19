package course

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"github.com/loukaspe/nursing-academiq/internal/core/services"
//	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
//	log "github.com/sirupsen/logrus"
//	"net/http"
//	"strconv"
//)
//
//type DeleteCourseHandler struct {
//	CourseService *services.CourseService
//	logger        *log.Logger
//}
//
//func NewDeleteCourseHandler(
//	service *services.CourseService,
//	logger *log.Logger,
//) *DeleteCourseHandler {
//	return &DeleteCourseHandler{
//		CourseService: service,
//		logger:        logger,
//	}
//}
//
//type DeleteCourseResponse struct {
//	ErrorMessage string `json:"errorMessage,omitempty"`
//}
//
//func (handler *DeleteCourseHandler) DeleteCourseController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	//var err error
//	response := &DeleteCourseResponse{}
//
//	id := mux.Vars(r)["id"]
//	if id == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "missing course id"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//	uid, err := strconv.Atoi(id)
//	if id == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "malformed course id"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	err = handler.CourseService.DeleteCourse(context.Background(), uint32(uid))
//	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
//		}).Debug("Error in updating solar panel data")
//
//		w.WriteHeader(dataNotFoundErrorWrapper.ReturnedStatusCode)
//
//		return
//	}
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = err.Error()
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	json.NewEncoder(w).Encode(response)
//	w.WriteHeader(http.StatusOK)
//	return
//}
