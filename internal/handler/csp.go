package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/salsadigitalauorg/lagoon-csp-collector/internal/util"
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
	Version              string
	Commit               string
	Project              util.Project
}

type CSPResponse struct {
	LagoonProject      string      `json:"lagoon_project"`
	Host               string      `json:"host"`
	OriginalURI        string      `json:"original_uri"`
	Disposition        string      `json:"disposition"`
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
		return
	}

	url, _ := url.Parse(report.Body.DocumentURI)
	csp.Project.Domain = url.Scheme + "://" + url.Host
	p, _ := csp.Project.GetName()

	json.NewEncoder(os.Stdout).Encode(CSPResponse{
		LagoonProject:      p,
		Host:               url.Host,
		OriginalURI:        report.Body.DocumentURI,
		Referrer:           report.Body.Referrer,
		ViolatedDirective:  report.Body.ViolatedDirective,
		EffectiveDirective: report.Body.EffectiveDirective,
		Policy:             report.Body.OriginalPolicy,
		BlockedURI:         report.Body.BlockedURI,
		Source:             report.Body.SourceFile,
		Status:             report.Body.StatusCode,
		LineNumber:         report.Body.LineNumber,
		Disposition:        report.Body.Disposition,
	})
	w.WriteHeader(http.StatusOK)
}
