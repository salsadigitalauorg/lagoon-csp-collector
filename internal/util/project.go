package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Project struct {
	API    string
	Name   string
	Domain string
}

/*
GetName from the configured API service.

Makes an API request to the configured service, this will pass the received domain
from the CSP violation and will determine if a Lagoon project can be inferred from
the hostname.

@TODO: Current expectation is the service matching domain to project will return a
slice of project names. To make this more robust, this handler should be pluggable.
*/
func (p *Project) GetName() (string, error) {
	client := &http.Client{}

	if p.API == "" {
		return "", errors.New("Missing API configuration")
	}

	req, err := http.NewRequest("GET", p.API+"?domain="+p.Domain, nil)

	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var r []string
	err = json.Unmarshal(body, &r)

	if err != nil {
		return "", err
	}

	if len(r) > 0 {
		p.Name = r[0]
		return p.Name, nil
	}

	return "", errors.New("Unable to get project from domain")
}
