package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Project struct {
	API    string
	Name   string
	Domain string
}

type Response struct {
	Name string `json:"field"`
}

func (p *Project) GetName() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", p.API, nil)

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
		fmt.Println(err)
		return "", nil
	}

	var r Response
	_ = json.Unmarshal(body, r)
	p.Name = r.Name

	return r.Name, nil
}
