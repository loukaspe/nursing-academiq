package quizSessionByStudent

type QuizSessionByStudentRequest struct {
	QuizSessionByStudent struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:""`
}

type CreateQuizSessionByStudentResponse struct {
	CreatedQuizSessionByStudentUid uint   `json:"insertedUid"`
	ErrorMessage                   string `json:"errorMessage,omitempty"`
}

//type CreateQuizSessionByStudentHandler struct {
//	QuizSessionByStudentService *services.QuizSessionByStudentService
//	logger        *log.Logger
//}
//
//func NewCreateQuizSessionByStudentHandler(
//	service *services.QuizSessionByStudentService,
//	logger *log.Logger,
//) *CreateQuizSessionByStudentHandler {
//	return &CreateQuizSessionByStudentHandler{
//		QuizSessionByStudentService: service,
//		logger:        logger,
//	}
//}

//func (handler *CreateQuizSessionByStudentHandler) CreateQuizSessionByStudentController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	response := &CreateQuizSessionByStudentResponse{}
//	request := &QuizSessionByStudentRequest{}
//
//	err := json.NewDecoder(r.Body).Decode(request)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating QuizSessionByStudent")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "malformed QuizSessionByStudent data"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	QuizSessionByStudentRequest := request.QuizSessionByStudent
//
//	birthDate, err := time.Parse("01-02-2006", QuizSessionByStudentRequest.BirthDate)
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
//	uid, err := handler.QuizSessionByStudentService.CreateQuizSessionByStudent(context.Background(), domainQuizSessionByStudent)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating QuizSessionByStudent in service")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "error creating QuizSessionByStudent"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	response.CreatedQuizSessionByStudentUid = uid
//
//	w.WriteHeader(http.StatusCreated)
//	return
//}
