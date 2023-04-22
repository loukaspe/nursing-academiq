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

// swagger:model HealthCheckResponse
type HealthCheckResponse struct {
	Message string `json:"message"`
}

// swagger:operation GET /health-check healthCheck
//
// # Check for the health of the app
//
// ---
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes:
//	- http
//	- https
//
// responses:
//   '200':
//     description: api is ok
//     schema:
//       "$ref": "#/definitions/HealthCheckResponse"
//   '500':
//     description: api got problem
//     schema:
//       "$ref": "#/definitions/HealthCheckResponse"

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
