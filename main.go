package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/salsadigitalauorg/lagoon-csp-collector/internal/handler"
)

func main() {
	domainFile := flag.String("domains", "", "Path to the domain/project mapping list")
	var domains map[string]string

	data, _ := ioutil.ReadFile(*domainFile)
	json.Unmarshal(data, &domains)

	http.HandleFunc("/v1", (&handler.CSPHandler{
		ReportOnly:           false,
		LogClientIP:          false,
		LogTruncatedClientIP: false,
		DomainList:           domains,
	}).Serve)

	http.HandleFunc("/v1/healthz", (&handler.HealthcheckHandler{}).Serve)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
