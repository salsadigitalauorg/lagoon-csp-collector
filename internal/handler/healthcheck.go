package handler

import (
	"encoding/json"
	"net/http"
)

type HealthStatus struct {
	Reason  string `json:"reason"`
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

type HealthcheckHandler struct {
	Version string
	Commit  string
}

func (cfg *HealthcheckHandler) Serve(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	h, _ := json.Marshal(HealthStatus{
		Reason:  "OK",
		Version: cfg.Version,
		Commit:  cfg.Commit,
	})

	w.Write(h)
}
