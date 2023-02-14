package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type CSPReport struct {
	Body CSPReportBody `json:"csp-report"`
}

type CSPReportBody struct {
	DocumentURI        string      `json:"document-uri"`
	Referrer           string      `json:"referrer"`
	ViolatedDirective  string      `json:"violated-directive"`
	EffectiveDirective string      `json:"effective-directive"`
	OriginalPolicy     string      `json:"original-policy"`
	Disposition        string      `json:"disposition"`
	BlockedURI         string      `json:"blocked-uri"`
	SourceFile         string      `json:"source-file"`
	ScriptSample       string      `json:"script-sample"`
	StatusCode         interface{} `json:"status-code"`
	LineNumber         interface{} `json:"line-number"`
}

// @TODO: Accept configured opensearch client
type CSPHandler struct {
	ReportOnly           bool
	LogClientIP          bool
	LogTruncatedClientIP bool

	DomainList map[string]string
}

type CSPResponse struct {
	LagoonProject      string      `json:"lagoon_project"`
	Host               string      `json:"host"`
	OriginalURI        string      `json:"original_uri"`
	Referrer           string      `json:"referrer"`
	ViolatedDirective  string      `json:"violated_directive"`
	EffectiveDirective string      `json:"effective_directive"`
	Policy             string      `json:"policy"`
	BlockedURI         string      `json:"blocked_uri"`
	Source             string      `json:"source"`
	Status             interface{} `json:"status"`
	LineNumber         interface{} `json:"line_number"`
}

type ErrorReponse struct {
	Reason  string `json:"reason"`
	Details string `json:"details"`
}

func (csp *CSPHandler) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var report CSPReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorReponse{
			Reason:  "Invalid domain provided",
			Details: fmt.Sprintf("%s", err),
		})
		return
	}

	url, _ := url.Parse(report.Body.DocumentURI)
	host := strings.TrimPrefix(url.Hostname(), "www.")

	lagoonProject, ok := csp.DomainList[host]
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorReponse{
			Details: "Invalid domain provided",
		})
		return
	}

	json.NewEncoder(os.Stdout).Encode(CSPResponse{
		LagoonProject:      lagoonProject,
		Host:               host,
		OriginalURI:        report.Body.DocumentURI,
		Referrer:           report.Body.Referrer,
		ViolatedDirective:  report.Body.ViolatedDirective,
		EffectiveDirective: report.Body.EffectiveDirective,
		Policy:             report.Body.OriginalPolicy,
		BlockedURI:         report.Body.BlockedURI,
		Source:             report.Body.SourceFile,
		Status:             report.Body.StatusCode,
		LineNumber:         report.Body.LineNumber,
	})
	w.WriteHeader(http.StatusOK)
}
