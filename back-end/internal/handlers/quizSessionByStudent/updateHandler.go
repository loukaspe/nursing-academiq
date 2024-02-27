package quizSessionByStudent

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
//type UpdateQuizSessionByStudentHandler struct {
//	QuizSessionByStudentService *services.QuizSessionByStudentService
//	logger        *log.Logger
//}
//
//func NewUpdateQuizSessionByStudentHandler(
//	service *services.QuizSessionByStudentService,
//	logger *log.Logger,
//) *UpdateQuizSessionByStudentHandler {
//	return &UpdateQuizSessionByStudentHandler{
//		QuizSessionByStudentService: service,
//		logger:        logger,
//	}
//}
//
//type UpdateQuizSessionByStudentResponse struct {
//	ErrorMessage string `json:"errorMessage,omitempty"`
//}
//
//func (handler *UpdateQuizSessionByStudentHandler) UpdateQuizSessionByStudentController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	response := &UpdateQuizSessionByStudentResponse{}
//
//	request := &QuizSessionByStudentRequest{}
//
//	err := json.NewDecoder(r.Body).Decode(request)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in updating QuizSessionByStudent")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "malformed QuizSessionByStudent request"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	QuizSessionByStudentRequest := request.QuizSessionByStudent
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
//	birthDate, err := time.Parse("17-03-2023", QuizSessionByStudentRequest.BirthDate)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating QuizSessionByStudent birth date")
//
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "malformed QuizSessionByStudent data: birth date"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	domainUser := &domain.User{
//		Username:    QuizSessionByStudentRequest.Username,
//		Password:    QuizSessionByStudentRequest.Password,
//		FirstName:   QuizSessionByStudentRequest.FirstName,
//		LastName:    QuizSessionByStudentRequest.LastName,
//		Email:       QuizSessionByStudentRequest.Email,
//		BirthDate:   birthDate,
//		PhoneNumber: QuizSessionByStudentRequest.PhoneNumber,
//		Photo:       QuizSessionByStudentRequest.Photo,
//	}
//
//	domainQuizSessionByStudent := &domain.QuizSessionByStudent{
//		User:               *domainUser,
//		RegistrationNumber: QuizSessionByStudentRequest.RegistrationNumber,
//	}
//
//	if dataNotFoundErrorWrapper, ok := err.(apierrors.DataNotFoundErrorWrapper); ok {
//	err = handler.QuizSessionByStudentService.UpdateQuizSessionByStudent(context.Background(), uint32(uid), domainQuizSessionByStudent)
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
