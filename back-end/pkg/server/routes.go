package server

import (
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	"github.com/loukaspe/nursing-academiq/internal/handlers"
	"github.com/loukaspe/nursing-academiq/internal/handlers/student"
	"github.com/loukaspe/nursing-academiq/internal/handlers/tutor"
	"github.com/loukaspe/nursing-academiq/internal/repositories"
	"github.com/loukaspe/nursing-academiq/pkg/auth"
	"net/http"
	"os"
)

func (s *Server) initializeRoutes() {
	// health check
	healthCheckHandler := handlers.NewHealthCheckHandler(s.DB)
	s.router.HandleFunc("/health-check", healthCheckHandler.HealthCheckController).Methods("GET")

	// auth
	jwtMechanism := auth.NewAuthMechanism(
		os.Getenv("JWT_SECRET_KEY"),
		os.Getenv("JWT_SIGNING_METHOD"),
	)
	jwtService := services.NewJwtService(jwtMechanism)
	jwtMiddleware := handlers.NewAuthenticationMw(jwtMechanism)

	userRepository := repositories.NewUserRepository(s.DB)
	loginService := services.NewLoginService(userRepository)

	jwtHandler := handlers.NewJwtClaimsHandler(jwtService, loginService, s.logger)

	s.router.HandleFunc("/login", jwtHandler.JwtTokenController).Methods(http.MethodPost)
	s.router.HandleFunc(
		"/login",
		optionsHandlerForCors,
	).Methods(http.MethodOptions)

	protected := s.router.PathPrefix("/").Subrouter()
	protected.Use(jwtMiddleware.AuthenticationMW)

	// student
	studentRepository := repositories.NewStudentRepository(s.DB)
	studentService := services.NewStudentService(studentRepository)

	getStudentHandler := student.NewGetStudentHandler(studentService, s.logger)
	createStudentHandler := student.NewCreateStudentHandler(studentService, s.logger)
	deleteStudentHandler := student.NewDeleteStudentHandler(studentService, s.logger)
	updateStudentHandler := student.NewUpdateStudentHandler(studentService, s.logger)

	protected.HandleFunc("/student", createStudentHandler.CreateStudentController).Methods("POST")
	protected.HandleFunc("/student", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/student/{id:[0-9]+}", getStudentHandler.GetStudentController).Methods("GET")
	protected.HandleFunc("/student/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/student/{id:[0-9]+}", deleteStudentHandler.DeleteStudentController).Methods("DELETE")
	protected.HandleFunc("/student/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/student/{id:[0-9]+}", updateStudentHandler.UpdateStudentController).Methods("PUT")
	protected.HandleFunc("/student/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	// tutor
	tutorRepository := repositories.NewTutorRepository(s.DB)
	tutorService := services.NewTutorService(tutorRepository)

	getTutorHandler := tutor.NewGetTutorHandler(tutorService, s.logger)
	createTutorHandler := tutor.NewCreateTutorHandler(tutorService, s.logger)
	deleteTutorHandler := tutor.NewDeleteTutorHandler(tutorService, s.logger)
	updateTutorHandler := tutor.NewUpdateTutorHandler(tutorService, s.logger)

	protected.HandleFunc("/tutor", createTutorHandler.CreateTutorController).Methods("POST")
	protected.HandleFunc("/tutor", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/tutor/{id:[0-9]+}", getTutorHandler.GetTutorController).Methods("GET")
	protected.HandleFunc("/tutor/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/tutor/{id:[0-9]+}", deleteTutorHandler.DeleteTutorController).Methods("DELETE")
	protected.HandleFunc("/tutor/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/tutor/{id:[0-9]+}", updateTutorHandler.UpdateTutorController).Methods("PUT")
	protected.HandleFunc("/tutor/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)

	// course
	courseRepository := repositories.NewCourseRepository(s.DB)
	courseService := services.NewCourseService(courseRepository)

	getCourseHandler := course.NewGetCourseHandler(courseService, s.logger)
	getCourseByStudentIDHandler := course.NewGetCourseByStudentIDHandler(courseService, s.logger)
	getCourseByTutorIDHandler := course.NewGetCourseByTutorIDHandler(courseService, s.logger)
	protected.HandleFunc("/course/{id:[0-9]+}", getCourseHandler.GetCourseController).Methods("GET")
	protected.HandleFunc("/course/{id:[0-9]+}", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/student/{id:[0-9]+}/courses", getCourseByStudentIDHandler.GetCourseByStudentIDController).Methods("GET")
	protected.HandleFunc("/student/{id:[0-9]+}/courses", optionsHandlerForCors).Methods(http.MethodOptions)
	protected.HandleFunc("/tutor/{id:[0-9]+}/courses", getCourseByTutorIDHandler.GetCourseByTutorIDController).Methods("GET")
	protected.HandleFunc("/tutor/{id:[0-9]+}/courses", optionsHandlerForCors).Methods(http.MethodOptions)

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
