package handlers

import (
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

func (handler *HealthCheckHandler) HealthCheckController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if postgreSQL, err := handler.db.DB(); err != nil {
		if err = postgreSQL.Ping(); err == nil {
			w.Write([]byte(`{message:"db not connected"}`))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.Write([]byte(`OK`))
	w.WriteHeader(http.StatusOK)
}
