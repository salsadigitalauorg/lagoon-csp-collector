package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/salsadigitalauorg/csp-opensearch-collector/internal/handler"
)

func main() {
	var domains map[string]string
	data, err := ioutil.ReadFile("domains.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	json.Unmarshal(data, &domains)

	if err != nil {
		log.Fatal(err)
		return
	}

	http.HandleFunc("/v1/healthz", (&handler.HealthcheckHandler{
		Client: client,
	}).Serve)

	http.HandleFunc("/v1/csp", (&handler.CSPHandler{
		ReportOnly:           false,
		LogClientIP:          false,
		LogTruncatedClientIP: false,
		DomainList:           domains,
	}).Serve)

	http.HandleFunc("/v1/csp/report-only", (&handler.CSPHandler{
		ReportOnly:           true,
		LogClientIP:          false,
		LogTruncatedClientIP: false,
		DomainList:           domains,
	}).Serve)

	log.Println("Server started listining on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
