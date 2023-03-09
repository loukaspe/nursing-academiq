package server

import (
	"github.com/loukaspe/nursing-academiq/internal/core/services"
	"github.com/loukaspe/nursing-academiq/internal/handlers"
	"github.com/loukaspe/nursing-academiq/internal/repositories"
)

func (s *Server) initializeRoutes() {

	// health check
	healthCheckHandler := handlers.NewHealthCheckHandler(s.DB)
	s.Router.HandleFunc("/health-check", healthCheckHandler.HealthCheckController).Methods("GET")

	// user
	userRepository := repositories.NewUserRepository(s.DB)
	userService := services.NewUserService(userRepository)

	getUserHandler := handlers.NewGetUserHandler(userService)
	createUserHandler := handlers.NewCreateUserHandler(userService)
	deleteUserHandler := handlers.NewDeleteUserHandler(userService)
	updateUserHandler := handlers.NewUpdateUserHandler(userService)

	s.Router.HandleFunc("/user", createUserHandler.CreateUserController).Methods("POST")
	s.Router.HandleFunc("/user/{id:[0-9]+}", getUserHandler.GetUserController).Methods("GET")
	s.Router.HandleFunc("/user/{id:[0-9]+}", deleteUserHandler.DeleteUserController).Methods("DELETE")
	s.Router.HandleFunc("/user/{id:[0-9]+}", updateUserHandler.UpdateUserController).Methods("PUT")
}
