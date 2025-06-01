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
		os.Getenv("JWT_ACCESS_SECRET_KEY"),
		os.Getenv("JWT_REFRESH_SECRET_KEY"),
		os.Getenv("JWT_SIGNING_METHOD"),
	)
	jwtService := services.NewJwtService(jwtMechanism)
	apiKey := os.Getenv("API_KEY")
	authMiddleware := handlers.NewAuthenticationMw(jwtMechanism, apiKey)

	userRepository := repositories.NewUserRepository(s.DB)
	loginService := services.NewLoginService(userRepository)
	userService := services.NewUserService(userRepository)

	loginHandler := handlers.NewLoginHandler(jwtService, loginService, s.logger)
	refreshTokenHandler := handlers.NewRefreshTokenHandler(jwtService, userService, s.logger)

	protectedJWT := s.router.PathPrefix("/").Subrouter()
	protectedJWT.Use(authMiddleware.JWTAuthenticationMW)

	protectedApiKey := s.router.PathPrefix("/").Subrouter()
	protectedApiKey.Use(authMiddleware.APIKeyAuthenticationMW)

	protectedApiKey.HandleFunc("/login", loginHandler.LoginController).Methods(http.MethodPost)
	protectedApiKey.HandleFunc(
		"/login",
		optionsHandlerForCors,
	).Methods(http.MethodOptions)

	protectedApiKey.HandleFunc("/refresh-token", refreshTokenHandler.RefreshTokenController).Methods(http.MethodPost)
	protectedApiKey.HandleFunc(
		"/refresh-token",
		optionsHandlerForCors,
	).Methods(http.MethodOptions)

	// user
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
	getCourseChaptersHandler := course.NewGetCourseChaptersHandler(courseService, s.logger)
	getCoursesHandler := course.NewGetCoursesHandler(courseService, s.logger)
	getMostRecentCoursesHandler := course.NewGetMostRecentCoursesHandler(courseService, s.logger)
	getCourseByTutorIDHandler := course.NewGetCourseByTutorIDHandler(courseService, s.logger)
	createCourseHandler := course.NewCreateCourseHandler(courseService, s.logger)
	deleteCourseHandler := course.NewDeleteCourseHandler(courseService, s.logger)
	updateCourseHandler := course.NewUpdateCourseHandler(courseService, s.logger)

	protectedJWT.HandleFunc("/course", createCourseHandler.CreateCourseController).Methods("POST")
	protectedJWT.HandleFunc("/course", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}", getCourseHandler.GetCourseController).Methods("GET")
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/extended", getExtendedCourseHandler.GetExtendedCourseController).Methods("GET")
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/extended", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/chapters", getCourseChaptersHandler.GetCourseChaptersController).Methods("GET")
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/chapters", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/courses", getCoursesHandler.GetCoursesController).Methods("GET")
	protectedApiKey.HandleFunc("/courses", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/courses/recent", getMostRecentCoursesHandler.GetMostRecentCoursesController).Methods("GET")
	protectedApiKey.HandleFunc("/courses/recent", optionsHandlerForCors).Methods(http.MethodOptions)
	// Wanted to do a PATCH but did not work with CORS
	protectedJWT.HandleFunc("/course/{id:[0-9]+}", updateCourseHandler.UpdateCourseController).Methods("PUT")
	protectedJWT.HandleFunc("/course/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/course/{id:[0-9]+}", deleteCourseHandler.DeleteCourseController).Methods("DELETE")
	protectedJWT.HandleFunc("/course/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/courses", getCourseByTutorIDHandler.GetCourseByTutorIDController).Methods("GET")
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/courses", optionsHandlerForCors).Methods(http.MethodOptions)

	// quiz
	quizRepository := repositories.NewQuizRepository(s.DB)
	quizService := services.NewQuizService(quizRepository)

	getQuizHandler := quiz.NewGetQuizHandler(quizService, s.logger)
	getQuizzesHandler := quiz.NewGetQuizzesHandler(quizService, s.logger)
	getMostRecentQuizzesHandler := quiz.NewGetMostRecentQuizzesHandler(quizService, s.logger)
	getQuizByTutorIDHandler := quiz.NewGetQuizByTutorIDHandler(quizService, s.logger)
	getQuizByCourseIDHandler := quiz.NewGetQuizByCourseIDHandler(quizService, s.logger)
	createQuizHandler := quiz.NewCreateQuizHandler(quizService, s.logger)
	deleteQuizHandler := quiz.NewDeleteQuizHandler(quizService, s.logger)
	updateQuizHandler := quiz.NewUpdateQuizHandler(quizService, s.logger)
	updateQuizQuestionsHandler := quiz.NewUpdateQuizQuestionsHandler(quizService, s.logger)

	protectedJWT.HandleFunc("/quiz", createQuizHandler.CreateQuizController).Methods("POST")
	protectedJWT.HandleFunc("/quiz", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/quizzes", getQuizzesHandler.GetQuizzesController).Methods("GET")
	protectedApiKey.HandleFunc("/quizzes", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/quizzes/recent", getMostRecentQuizzesHandler.GetMostRecentQuizzesController).Methods("GET")
	protectedApiKey.HandleFunc("/quizzes/recent", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/quiz/{id:[0-9]+}", getQuizHandler.GetQuizController).Methods("GET")
	protectedApiKey.HandleFunc("/quiz/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", deleteQuizHandler.DeleteQuizController).Methods("DELETE")
	protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", updateQuizHandler.UpdateQuizController).Methods("PUT")
	protectedJWT.HandleFunc("/quiz/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/quiz/{id:[0-9]+}/questions", updateQuizQuestionsHandler.UpdateQuizQuestionsController).Methods("POST")
	protectedJWT.HandleFunc("/quiz/{id:[0-9]+}/questions", optionsHandlerForCors).Methods(http.MethodOptions)

	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/quizzes", getQuizByTutorIDHandler.GetQuizByTutorIDController).Methods("GET")
	protectedJWT.HandleFunc("/tutor/{id:[0-9]+}/quizzes", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/quizzes", getQuizByCourseIDHandler.GetQuizByCourseIDController).Methods("GET")
	protectedApiKey.HandleFunc("/course/{id:[0-9]+}/quizzes", optionsHandlerForCors).Methods(http.MethodOptions)

	// chapter
	chapterRepository := repositories.NewChapterRepository(s.DB)
	chapterService := services.NewChapterService(chapterRepository, quizRepository)

	getChapterHandler := chapter.NewGetChapterHandler(chapterService, s.logger)
	updateChapterHandler := chapter.NewUpdateChapterHandler(chapterService, s.logger)
	deleteChapterHandler := chapter.NewDeleteChapterHandler(chapterService, s.logger)
	createChapterHandler := chapter.NewCreateChapterHandler(chapterService, s.logger)

	protectedJWT.HandleFunc("/chapter", createChapterHandler.CreateChapterController).Methods("POST")
	protectedJWT.HandleFunc("/chapter", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/chapter/{id:[0-9]+}", getChapterHandler.GetChapterController).Methods("GET")
	protectedApiKey.HandleFunc("/chapter/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	// Wanted to do a PATCH but did not work with CORS
	protectedJWT.HandleFunc("/chapter/{id:[0-9]+}", updateChapterHandler.UpdateChapterController).Methods("PUT")
	protectedJWT.HandleFunc("/chapter/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/chapter/{id:[0-9]+}", deleteChapterHandler.DeleteChapterController).Methods("DELETE")
	protectedJWT.HandleFunc("/chapter/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	// question
	questionsRepository := repositories.NewQuestionRepository(s.DB)
	questionsService := services.NewQuestionService(questionsRepository)

	importQuestionHandler := question.NewImportQuestionHandler(questionsService, chapterService, s.logger)
	getQuestionHandler := question.NewGetQuestionHandler(questionsService, s.logger)
	getQuestionByCourseIDHandler := question.NewGetQuestionByCourseIDHandler(questionsService, s.logger)
	updateQuestionHandler := question.NewUpdateQuestionHandler(questionsService, s.logger)
	deleteQuestionHandler := question.NewDeleteQuestionHandler(questionsService, s.logger)
	bulkDeleteQuestionsHandler := question.NewBulkDeleteQuestionHandler(questionsService, s.logger)
	createQuestionHandler := question.NewCreateQuestionHandler(questionsService, s.logger)

	protectedJWT.HandleFunc("/questions", createQuestionHandler.CreateQuestionController).Methods("POST")
	protectedJWT.HandleFunc("/questions", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/questions/{id:[0-9]+}", getQuestionHandler.GetQuestionController).Methods("GET")
	protectedApiKey.HandleFunc("/questions/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	// Wanted to do a PATCH but did not work with CORS
	protectedJWT.HandleFunc("/questions/{id:[0-9]+}", updateQuestionHandler.UpdateQuestionController).Methods("PUT")
	protectedJWT.HandleFunc("/questions/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/questions/{id:[0-9]+}", deleteQuestionHandler.DeleteQuestionController).Methods("DELETE")
	protectedJWT.HandleFunc("/questions/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedJWT.HandleFunc("/questions/bulk", bulkDeleteQuestionsHandler.BulkDeleteQuestionController).Methods("POST")
	protectedJWT.HandleFunc("/questions/bulk", optionsHandlerForCors).Methods(http.MethodOptions)

	// TODO change to jwt since only tutors can
	protectedJWT.HandleFunc("/courses/{id:[0-9]+}/questions/import", importQuestionHandler.ImportQuestionController).Methods("POST")
	protectedJWT.HandleFunc("/courses/{id:[0-9]+}/questions/import", optionsHandlerForCors).Methods(http.MethodOptions)
	protectedApiKey.HandleFunc("/courses/{id:[0-9]+}/questions", getQuestionByCourseIDHandler.GetQuestionByCourseIDController).Methods("GET")
	protectedApiKey.HandleFunc("/courses/{id:[0-9]+}/questions", optionsHandlerForCors).Methods(http.MethodOptions)

	s.router.Use(mux.CORSMethodMiddleware(s.router))

	//// TODO fix allowed origns
	corsOrigins := gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000", "https://https://nursing-academiq.vercel.app/"})
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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

}
