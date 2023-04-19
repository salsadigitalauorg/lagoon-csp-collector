package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/salsadigitalauorg/lagoon-csp-collector/internal/handler"
	"github.com/salsadigitalauorg/lagoon-csp-collector/internal/util"
)

var (
	version string
	commit  string
)

func main() {
	port := flag.String("port", "3000", "Port to run the collector on")
	a := flag.String("api", "", "The endpoint to hydrate the CSP report")
	d := flag.String("test-domain", "", "A domain to validate in the health check")

	flag.Parse()

	project := util.Project{API: *a}

	http.HandleFunc("/v1", (&handler.CSPHandler{
		ReportOnly:           false,
		LogClientIP:          false,
		LogTruncatedClientIP: false,
		Version:              version,
		Commit:               commit,
		Project:              project,
	}).Serve)

	http.HandleFunc("/v1/healthz", (&handler.HealthcheckHandler{
		Version:   version,
		Commit:    commit,
		APIConfig: *a,
		Domain:    *d,
	}).Serve)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}
