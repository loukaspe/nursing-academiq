package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func NewServer(
	db *gorm.DB,
	router *mux.Router,
) *Server {
	return &Server{
		DB:     db,
		Router: router,
	}
}

func (s *Server) Run(addr string) {
	s.initializeRoutes()

	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
