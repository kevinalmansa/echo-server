package endpoints

import (
	"io/ioutil"
	"net/http"

	"github.com/kevinalmansa/echo-server/internal/models"
)

//EndpointEcho definition
type EndpointEcho struct{}

//NewEndpointEcho is a factory to create this struct
func NewEndpointEcho() *EndpointEcho {
	return &EndpointEcho{}
}

//Routes sets up routing for this endpoint
func (e *EndpointEcho) Routes() http.Handler {
	//no subroutes, so return self
	return e
}

//ServeHTTP to impliment http.Handler interface. this is what is called
//when the endpoint is requested by http clients.
func (e *EndpointEcho) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	// if r.Method != "GET" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }

	request := &models.Request{}
	var responseBody []byte
	var err error

	request.Host = r.Host
	request.Path = r.URL.Path
	request.Method = r.Method
	request.Header = r.Header
	request.Body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ua := r.Header.Get("User-Agent")
	if ua == "application/json" {
		responseBody, err = request.JsonOutput()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
	} else {
		responseBody, err = request.TextOutput()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
	}

	//Response
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
