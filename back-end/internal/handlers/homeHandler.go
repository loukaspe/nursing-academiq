package handlers

import (
	"net/http"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (handler *HomeHandler) HomeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(`{message:"aaaaaaaa"}`))
	w.WriteHeader(http.StatusOK)
}
