package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/salsadigitalauorg/lagoon-csp-collector/internal/util"
)

type HealthStatus struct {
	Reason    string `json:"reason"`
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	APIStatus string `json:"api_status"`
}

type HealthcheckHandler struct {
	Version   string
	Commit    string
	APIConfig string
	Domain    string
}

func (cfg *HealthcheckHandler) Serve(w http.ResponseWriter, r *http.Request) {

	var project util.Project

	url, _ := url.Parse(cfg.Domain)
	project.API = cfg.APIConfig
	project.Domain = url.Scheme + "://" + url.Host
	p, err := project.GetName()

	status := HealthStatus{
		Version: cfg.Version,
		Commit:  cfg.Commit,
	}

	if err != nil {
		status.Reason = fmt.Sprint(err)
		h, _ := json.Marshal(status)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(h)
		return
	}

	status.Reason = "OK"
	status.APIStatus = p

	h, _ := json.Marshal(status)
	w.WriteHeader(http.StatusOK)
	w.Write(h)
}
