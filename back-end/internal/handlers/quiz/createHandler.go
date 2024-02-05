package quiz

type QuizRequest struct {
	Quiz struct {
		Title             string
		Description       string
		Visibility        bool
		ShowSubset        bool
		SubsetSize        int
		NumberOfSessions  int
		ScoreSum          float32
		MaxScore          int
		NumberOfQuestions int
	} `json:""`
}

type CreateQuizResponse struct {
	CreatedQuizUid uint   `json:"insertedUid"`
	ErrorMessage   string `json:"errorMessage,omitempty"`
}

//type CreateQuizHandler struct {
//	QuizService *services.QuizService
//	logger        *log.Logger
//}
//
//func NewCreateQuizHandler(
//	service *services.QuizService,
//	logger *log.Logger,
//) *CreateQuizHandler {
//	return &CreateQuizHandler{
//		QuizService: service,
//		logger:        logger,
//	}
//}

//func (handler *CreateQuizHandler) CreateQuizController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	response := &CreateQuizResponse{}
//	request := &QuizRequest{}
//
//	err := json.NewDecoder(r.Body).Decode(request)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating quiz")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "malformed quiz data"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	quizRequest := request.Quiz
//
//	birthDate, err := time.Parse("01-02-2006", quizRequest.BirthDate)
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
//	uid, err := handler.QuizService.CreateQuiz(context.TODO(), domainQuiz)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating quiz in service")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "error creating quiz"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	response.CreatedQuizUid = uid
//
//	w.WriteHeader(http.StatusCreated)
//	return
//}
