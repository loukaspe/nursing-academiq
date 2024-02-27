package quizSessionByStudent

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/loukaspe/nursing-academiq/internal/core/services"
//	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
//	log "github.com/sirupsen/logrus"
//	"net/http"
//)
//
//type GetQuizSessionsByStudentHandler struct {
//	QuizSessionByStudentService *services.QuizSessionByStudentService
//	logger                      *log.Logger
//}
//
//func NewGetQuizSessionsByStudentHandler(
//	service *services.QuizSessionByStudentService,
//	logger *log.Logger,
//) *GetQuizSessionByStudentHandler {
//	return &GetQuizSessionByStudentHandler{
//		QuizSessionByStudentService: service,
//		logger:                      logger,
//	}
//}
//
//type GetQuizSessionsByStudentResponse struct {
//	ErrorMessage          string `json:"errorMessage,omitempty"`
//	QuizSessionsByStudent []struct {
//		ID          uint32 `json:"id"`
//		Title       string `json:"title"`
//		Description string `json:"description"`
//	} `json:"QuizSessionsByStudent,omitempty"`
//}
//
//func (handler *GetQuizSessionByStudentHandler) GetQuizSessionsByStudentController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	//var err error
//	response := &GetQuizSessionsByStudentResponse{}
//
//	QuizSessionsByStudent, err := handler.QuizSessionByStudentService.GetQuizSessionsByStudent(context.TODO())
//	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": dataNotFoundErrorWrapper.Unwrap().Error(),
//		}).Debug("Error in getting QuizSessionsByStudent")
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
//	for _, QuizSessionByStudent := range QuizSessionsByStudent {
//		response.QuizSessionsByStudent = append(response.QuizSessionsByStudent, struct {
//			ID          uint32 `json:"id"`
//			Title       string `json:"title"`
//			Description string `json:"description"`
//		}{
//			ID:          QuizSessionByStudent.ID,
//			Title:       QuizSessionByStudent.Title,
//			Description: QuizSessionByStudent.Description,
//		})
//	}
//
//	json.NewEncoder(w).Encode(response)
//	w.WriteHeader(http.StatusOK)
//	return
//}
