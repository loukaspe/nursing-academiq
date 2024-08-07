package server

import (
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	"github.com/loukaspe/nursing-academiq/internal/handlers"
	"github.com/loukaspe/nursing-academiq/internal/handlers/chapter"
	"github.com/loukaspe/nursing-academiq/internal/handlers/course"
	"github.com/loukaspe/nursing-academiq/internal/handlers/question"
	"github.com/loukaspe/nursing-academiq/internal/handlers/quiz"
	"github.com/loukaspe/nursing-academiq/internal/handlers/tutor"
	"github.com/loukaspe/nursing-academiq/internal/handlers/user"
	"github.com/loukaspe/nursing-academiq/internal/repositories"
	"github.com/loukaspe/nursing-academiq/pkg/auth"
	"net/http"
	"os"
)

func (s *Server) initializeRoutes() {
	fs := http.FileServer(http.Dir("./uploads"))
	s.router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	// health check
	healthCheckHandler := handlers.NewHealthCheckHandler(s.DB)
	s.router.HandleFunc("/health-check", healthCheckHandler.HealthCheckController).Methods("GET")

	// auth
	jwtMechanism := auth.NewAuthMechanism(
		os.Getenv("JWT_SECRET_KEY"),
		os.Getenv("JWT_SIGNING_METHOD"),
	)
	jwtService := services.NewJwtService(jwtMechanism)
	apiKey := os.Getenv("API_KEY")
	authMiddleware := handlers.NewAuthenticationMw(jwtMechanism, apiKey)

	userRepository := repositories.NewUserRepository(s.DB)
	loginService := services.NewLoginService(userRepository)

	jwtHandler := handlers.NewJwtClaimsHandler(jwtService, loginService, s.logger)

	protectedJWT := s.router.PathPrefix("/").Subrouter()
	protectedJWT.Use(authMiddleware.JWTAuthenticationMW)

	protectedApiKey := s.router.PathPrefix("/").Subrouter()
	protectedApiKey.Use(authMiddleware.APIKeyAuthenticationMW)

	s.router.HandleFunc("/login", jwtHandler.JwtTokenController).Methods(http.MethodPost)
	s.router.HandleFunc(
		"/login",
		optionsHandlerForCors,
	).Methods(http.MethodOptions)

	// user
	userService := services.NewUserService(userRepository)

	changeUserPasswordHandler := user.NewChangeUserPasswordHandler(userService, s.logger)

	protectedJWT.HandleFunc("/user/{id:[0-9]+}/change_password", changeUserPasswordHandler.ChangeUserPasswordController).Methods("POST")
	protectedJWT.HandleFunc("/user/{id:[0-9]+}/change_password", optionsHandlerForCors).Methods(http.MethodOptions)

	// tutor
	tutorRepository := repositories.NewTutorRepository(s.DB)
	tutorService := services.NewTutorService(tutorRepository)

	getTutorHandler := tutor.NewGetTutorHandler(tutorService, s.logger)
	createTutorHandler := tutor.NewCreateTutorHandler(tutorService, s.logger)
	deleteTutorHandler := tutor.NewDeleteTutorHandler(tutorService, s.logger)
	updateTutorHandler := tutor.NewUpdateTutorHandler(tutorService, s.logger)

	protectedJWT.HandleFunc("/tutor", createTutorHandler.CreateTutorController).Methods("POST")
	protectedJWT.HandleFunc("/tutor", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}", getTutorHandler.GetTutorController).Methods("GET")
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}", deleteTutorHandler.DeleteTutorController).Methods("DELETE")
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}", updateTutorHandler.UpdateTutorController).Methods("PUT")
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	// course
	courseRepository := repositories.NewCourseRepository(s.DB)
	courseService := services.NewCourseService(courseRepository)

	getCourseHandler := course.NewGetCourseHandler(courseService, s.logger)
	getExtendedCourseHandler := course.NewGetExtendedCourseHandler(courseService, s.logger)
	getCoursesHandler := course.NewGetCoursesHandler(courseService, s.logger)
	getCourseByTutorIDHandler := course.NewGetCourseByTutorIDHandler(courseService, s.logger)
	//createCourseHandler := course.NewCreateCourseHandler(courseService, s.logger)
	//deleteCourseHandler := course.NewDeleteCourseHandler(courseService, s.logger)
	//updateCourseHandler := course.NewUpdateCourseHandler(courseService, s.logger)

	//protectedJWT.HandleFunc("/course", createCourseHandler.CreateCourseController).Methods("POST")
	//protectedJWT.HandleFunc("/course", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}", getCourseHandler.GetCourseController).Methods("GET")
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/extended", getExtendedCourseHandler.GetExtendedCourseController).Methods("GET")
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/extended", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/courses", getCoursesHandler.GetCoursesController).Methods("GET")
	protectedApiKey.HandleFunc("/courses", optionsHandlerForCors).Methods(http.MethodOptions)
	//protectedJWT.HandleFunc("/course/{id:[0-9]+}", deleteCourseHandler.DeleteCourseController).Methods("DELETE")
	//protectedJWT.HandleFunc("/course/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	//protectedJWT.HandleFunc("/course/{id:[0-9]+}", updateCourseHandler.UpdateCourseController).Methods("PUT")
	//protectedJWT.HandleFunc("/course/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/courses", getCourseByTutorIDHandler.GetCourseByTutorIDController).Methods("GET")
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/courses", optionsHandlerForCors).Methods(http.MethodOptions)

	// quiz
	quizRepository := repositories.NewQuizRepository(s.DB)
	quizService := services.NewQuizService(quizRepository)

	getQuizHandler := quiz.NewGetQuizHandler(quizService, s.logger)
	getQuizByTutorIDHandler := quiz.NewGetQuizByTutorIDHandler(quizService, s.logger)
	getQuizByCourseIDHandler := quiz.NewGetQuizByCourseIDHandler(quizService, s.logger)
	//createQuizHandler := quiz.NewCreateQuizHandler(quizService, s.logger)
	//deleteQuizHandler := quiz.NewDeleteQuizHandler(quizService, s.logger)
	//updateQuizHandler := quiz.NewUpdateQuizHandler(quizService, s.logger)

	//protectedJWT.HandleFunc("/quiz", createQuizHandler.CreateQuizController).Methods("POST")
	//protectedJWT.HandleFunc("/quiz", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/quiz/{id:[0-9]+}", getQuizHandler.GetQuizController).Methods("GET")
	protectedApiKey.HandleFunc("/quiz/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	//protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", deleteQuizHandler.DeleteQuizController).Methods("DELETE")
	//protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	//protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", updateQuizHandler.UpdateQuizController).Methods("PUT")
	//protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/quizzes", getQuizByTutorIDHandler.GetQuizByTutorIDController).Methods("GET")
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/quizzes", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/quizzes", getQuizByCourseIDHandler.GetQuizByCourseIDController).Methods("GET")
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/quizzes", optionsHandlerForCors).Methods(http.MethodOptions)

	// chapter
	chapterRepository := repositories.NewChapterRepository(s.DB)
	chapterService := services.NewChapterService(chapterRepository, quizRepository)

	getChapterHandler := chapter.NewGetChapterHandler(chapterService, s.logger)

	protectedApiKey.HandleFunc("/chapter/{id:[0-9]+}", getChapterHandler.GetChapterController).Methods("GET")
	protectedApiKey.HandleFunc("/chapter/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	// question
	questionsRepository := repositories.NewQuestionRepository(s.DB)
	questionsService := services.NewQuestionService(questionsRepository)

	importQuestionHandler := question.NewImportQuestionHandler(questionsService, s.logger)

	protectedApiKey.HandleFunc("/courses/{id:[0-9]+}/questions/import", importQuestionHandler.ImportQuestionController).Methods("POST")
	protectedApiKey.HandleFunc("/courses/{id:[0-9]+}/questions/import", optionsHandlerForCors).Methods(http.MethodOptions)

	s.router.Use(mux.CORSMethodMiddleware(s.router))

	//// TODO fix allowed origns
	corsOrigins := gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"})
	corsMethods := gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	corsCredentials := gorillaHandlers.AllowCredentials()
	corsHandler := gorillaHandlers.CORS(corsOrigins, corsMethods, corsHeaders, corsCredentials)
	s.router.Use(corsHandler)
}

func optionsHandlerForCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Max-Age", "86400")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

}
