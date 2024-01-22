package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type HealthCheckHandler struct {
	db *gorm.DB
}

func NewHealthCheckHandler(db *gorm.DB) *HealthCheckHandler {
	return &HealthCheckHandler{
		db: db,
	}
}

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func (handler *HealthCheckHandler) HealthCheckController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthCheckResponse{}

	if postgreSQL, err := handler.db.DB(); err != nil {
		if err = postgreSQL.Ping(); err == nil {
			response.Message = "db not connected"
			json.NewEncoder(w).Encode(response)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	response.Message = "OK"
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}
