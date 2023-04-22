package server

import (
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	"github.com/loukaspe/nursing-academiq/internal/handlers"
	"github.com/loukaspe/nursing-academiq/internal/handlers/student"
	"github.com/loukaspe/nursing-academiq/internal/handlers/tutor"
	"github.com/loukaspe/nursing-academiq/internal/repositories"
)

func (s *Server) initializeRoutes() {
	// health check
	healthCheckHandler := handlers.NewHealthCheckHandler(s.DB)
	s.router.HandleFunc("/health-check", healthCheckHandler.HealthCheckController).Methods("GET")

	// student
	studentRepository := repositories.NewStudentRepository(s.DB)
	studentService := services.NewStudentService(studentRepository)

	getStudentHandler := student.NewGetStudentHandler(studentService, s.logger)
	createStudentHandler := student.NewCreateStudentHandler(studentService, s.logger)
	deleteStudentHandler := student.NewDeleteStudentHandler(studentService, s.logger)
	updateStudentHandler := student.NewUpdateStudentHandler(studentService, s.logger)

	s.router.HandleFunc("/student", createStudentHandler.CreateStudentController).Methods("POST")
	s.router.HandleFunc("/student/{id:[0-9]+}", getStudentHandler.GetStudentController).Methods("GET")
	s.router.HandleFunc("/student/{id:[0-9]+}", deleteStudentHandler.DeleteStudentController).Methods("DELETE")
	s.router.HandleFunc("/student/{id:[0-9]+}", updateStudentHandler.UpdateStudentController).Methods("PUT")

	// tutor
	tutorRepository := repositories.NewTutorRepository(s.DB)
	tutorService := services.NewTutorService(tutorRepository)

	getTutorHandler := tutor.NewGetTutorHandler(tutorService, s.logger)
	createTutorHandler := tutor.NewCreateTutorHandler(tutorService, s.logger)
	deleteTutorHandler := tutor.NewDeleteTutorHandler(tutorService, s.logger)
	updateTutorHandler := tutor.NewUpdateTutorHandler(tutorService, s.logger)

	s.router.HandleFunc("/tutor", createTutorHandler.CreateTutorController).Methods("POST")
	s.router.HandleFunc("/tutor/{id:[0-9]+}", getTutorHandler.GetTutorController).Methods("GET")
	s.router.HandleFunc("/tutor/{id:[0-9]+}", deleteTutorHandler.DeleteTutorController).Methods("DELETE")
	s.router.HandleFunc("/tutor/{id:[0-9]+}", updateTutorHandler.UpdateTutorController).Methods("PUT")
}
