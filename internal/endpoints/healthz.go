package endpoints

import (
	"net/http"
)

//EndpointHealthz definition
type EndpointHealthz struct{}

//EndpointHealthz is a factory to create this struct
func NewEndpointHealthz() *EndpointHealthz {
	return &EndpointHealthz{}
}

//Routes sets up routing for this endpoint
func (e *EndpointHealthz) Routes() http.Handler {
	//no subroutes, so return self
	return e
}

//ServeHTTP to impliment http.Handler interface. this is what is called
//when the endpoint is requested by http clients.
func (e *EndpointHealthz) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	// if r.Method != "GET" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }

	//Response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
