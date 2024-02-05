package quiz

//
//import (
//	"context"
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"github.com/loukaspe/nursing-academiq/internal/core/domain"
//	"github.com/loukaspe/nursing-academiq/internal/core/services"
//	apierrors "github.com/loukaspe/nursing-academiq/pkg/errors"
//	log "github.com/sirupsen/logrus"
//	"net/http"
//	"strconv"
//	"time"
//)
//
//type UpdateQuizHandler struct {
//	QuizService *services.QuizService
//	logger        *log.Logger
//}
//
//func NewUpdateQuizHandler(
//	service *services.QuizService,
//	logger *log.Logger,
//) *UpdateQuizHandler {
//	return &UpdateQuizHandler{
//		QuizService: service,
//		logger:        logger,
//	}
//}
//
//type UpdateQuizResponse struct {
//	ErrorMessage string `json:"errorMessage,omitempty"`
//}
//
//func (handler *UpdateQuizHandler) UpdateQuizController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	response := &UpdateQuizResponse{}
//
//	request := &QuizRequest{}
//
//	err := json.NewDecoder(r.Body).Decode(request)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in updating quiz")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "malformed quiz request"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	quizRequest := request.Quiz
//
//	id := mux.Vars(r)["id"]
//	if id == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "missing quiz id"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//	uid, err := strconv.Atoi(id)
//	if id == "" {
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "malformed quiz id"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	birthDate, err := time.Parse("17-03-2023", quizRequest.BirthDate)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating quiz birth date")
//
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "malformed quiz data: birth date"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	domainUser := &domain.User{
//		Username:    quizRequest.Username,
//		Password:    quizRequest.Password,
//		FirstName:   quizRequest.FirstName,
//		LastName:    quizRequest.LastName,
//		Email:       quizRequest.Email,
//		BirthDate:   birthDate,
//		PhoneNumber: quizRequest.PhoneNumber,
//		Photo:       quizRequest.Photo,
//	}
//
//	domainQuiz := &domain.Quiz{
//		User:               *domainUser,
//		RegistrationNumber: quizRequest.RegistrationNumber,
//	}
//
//	err = handler.QuizService.UpdateQuiz(context.TODO(), uint32(uid), domainQuiz)
//	if dataNotFoundErrorWrapper, ok := err.(*apierrors.DataNotFoundErrorWrapper); ok {
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
//	w.WriteHeader(http.StatusOK)
//	return
//}
