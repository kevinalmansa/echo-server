package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Path   string      `json:"path"`
	Host   string      `json:"host"`
	Method string      `json:"method"`
	Header http.Header `json:"headers"`
	Body   []byte      `json:"body"`
}

func (r Request) JsonOutput() ([]byte, error) {
	response, err := json.Marshal(r)
	return response, err
}

func (r Request) TextOutput() ([]byte, error) {
	response := ""

	response += fmt.Sprintf("HOST: %s\n", r.Host)
	response += fmt.Sprintf("URL: %s\n", r.Path)
	response += fmt.Sprintf("METHOD: %s\n", r.Method)
	response += "HEADERS:\n"
	for name, headers := range r.Header {
		for _, h := range headers {
			response += fmt.Sprintf("%v: %v\n", name, h)
		}
	}
	response += "\nBODY:\n"
	response += string(r.Body)
	return []byte(response), nil
}
