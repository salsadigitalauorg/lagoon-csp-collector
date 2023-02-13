package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/opensearch-project/opensearch-go"
)

type HealthStatus struct {
	Reason  string `json:"reason"`
	Details string `json:"details"`
}

type HealthcheckHandler struct {
	Client *opensearch.Client
}

func (cfg *HealthcheckHandler) Serve(w http.ResponseWriter, r *http.Request) {
	_, err := cfg.Client.Info()
	w.Header().Set("content-type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(HealthStatus{
			Reason:  "Unable to connect to opensearch",
			Details: fmt.Sprintf("%v", err),
		})
		return
	}

	// @TODO: Additional logic for healthcheck.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthStatus{
		Reason: "",
	})
}
