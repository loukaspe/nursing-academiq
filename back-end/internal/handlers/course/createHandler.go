package course

type CourseRequest struct {
	Course struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:""`
}

type CreateCourseResponse struct {
	CreatedCourseUid uint   `json:"insertedUid"`
	ErrorMessage     string `json:"errorMessage,omitempty"`
}

//type CreateCourseHandler struct {
//	CourseService *services.CourseService
//	logger        *log.Logger
//}
//
//func NewCreateCourseHandler(
//	service *services.CourseService,
//	logger *log.Logger,
//) *CreateCourseHandler {
//	return &CreateCourseHandler{
//		CourseService: service,
//		logger:        logger,
//	}
//}

//func (handler *CreateCourseHandler) CreateCourseController(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//
//	response := &CreateCourseResponse{}
//	request := &CourseRequest{}
//
//	err := json.NewDecoder(r.Body).Decode(request)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating course")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "malformed course data"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	courseRequest := request.Course
//
//	birthDate, err := time.Parse("01-02-2006", courseRequest.BirthDate)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating course birth date")
//
//		w.WriteHeader(http.StatusBadRequest)
//		response.ErrorMessage = "malformed course data: birth date"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	domainUser := &domain.User{
//		Username:    courseRequest.Username,
//		Password:    courseRequest.Password,
//		FirstName:   courseRequest.FirstName,
//		LastName:    courseRequest.LastName,
//		Email:       courseRequest.Email,
//		BirthDate:   birthDate,
//		PhoneNumber: courseRequest.PhoneNumber,
//		Photo:       courseRequest.Photo,
//	}
//
//	domainCourse := &domain.Course{
//		User:               *domainUser,
//		RegistrationNumber: courseRequest.RegistrationNumber,
//	}
//
//	uid, err := handler.CourseService.CreateCourse(context.TODO(), domainCourse)
//	if err != nil {
//		handler.logger.WithFields(log.Fields{
//			"errorMessage": err.Error(),
//		}).Error("Error in creating course in service")
//
//		w.WriteHeader(http.StatusInternalServerError)
//		response.ErrorMessage = "error creating course"
//		json.NewEncoder(w).Encode(response)
//		return
//	}
//
//	response.CreatedCourseUid = uid
//
//	w.WriteHeader(http.StatusCreated)
//	return
//}
