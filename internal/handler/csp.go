package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
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

func ProjectFromDocumentURI(d string) (string, error) {
	url, err := url.Parse(d)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(url.Hostname(), "www."), nil
}

func (csp *CSPHandler) Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var report CSPReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// DocumentURI in the CSP payload will container a FQDN of
	// the violation we can keep an in-memory map of all knwon
	// subscribers to the CSP violation service and map them back
	// to Lagoon projects to manage the index.
	// p, err := ProjectFromDocumentURI(report.Body.DocumentURI)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("Unable to parse document-uri %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
