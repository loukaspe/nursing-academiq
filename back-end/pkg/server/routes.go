package server

import "github.com/loukaspe/nursing-academiq/internal/handlers"

func (s *Server) initializeRoutes() {
	homeHandler := handlers.NewHomeHandler()

	s.Router.HandleFunc("/", homeHandler.HomeController).Methods("GET")
}
