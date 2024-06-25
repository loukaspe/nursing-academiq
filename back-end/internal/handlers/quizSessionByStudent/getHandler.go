package quizSessionByStudent

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
//type GetQuizSessionByStudentHandler struct {
//	QuizSessionByStudentService *services.QuizSessionByStudentService
//	logger                      *log.Logger
//}
//
//func NewGetQuizSessionByStudentHandler(
//	service *services.QuizSessionByStudentService,
//	logger *log.Logger,
//) *GetQuizSessionByStudentHandler {
//	return &GetQuizSessionByStudentHandler{
//		QuizSessionByStudentService: service,
//		logger:                      logger,
//	}
//}
//
//type GetQuizSessionByStudentResponse struct {
//	ErrorMessage         string                       `json:"errorMessage,omitempty"`
//	QuizSessionByStudent *QuizSessionByStudentRequest `json:"QuizSessionByStudent,omitempty"`
//}
//
//func (handler *GetQuizSessionByStudentHandler) GetQuizSessionByStudentController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	//var err error
//	response := &GetQuizSessionByStudentResponse{}
//
//	id := mux.Vars(r)["id"]
//	if id == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "missing QuizSessionByStudent id"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//	uid, err := strconv.Atoi(id)
//	if id == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "malformed QuizSessionByStudent id"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	QuizSessionByStudent, err := handler.QuizSessionByStudentService.GetQuizSessionByStudent(context.Background(), uint32(uid))
//	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
//		}).Debug("Error in getting user data")
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
//	response.QuizSessionByStudent = &QuizSessionByStudentRequest{
//		QuizSessionByStudent: struct {
//			Title       string `json:"title"`
//			Description string `json:"description"`
//		}{
//			Title:       QuizSessionByStudent.Title,
//			Description: QuizSessionByStudent.Description,
//		},
//	}
//
//	json.NewEncoder(w).Encode(response)
//	w.WriteHeader(http.StatusOK)
//	return
//}
