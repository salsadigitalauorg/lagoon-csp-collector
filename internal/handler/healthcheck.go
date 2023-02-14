package handler

import (
	"net/http"
)

type HealthStatus struct {
	Reason  string `json:"reason"`
	Details string `json:"details"`
}

type HealthcheckHandler struct{}

func (cfg *HealthcheckHandler) Serve(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
