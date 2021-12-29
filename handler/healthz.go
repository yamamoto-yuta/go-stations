package handler

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := &model.HealthzResponse{
		Message: "OK",
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(response); err != nil {
		log.Println(err)
	}
}
