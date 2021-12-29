package webserver

import "net/http"

//Endpoint interface to abstract away endpoint logic, and permit a simple
//interface for webserver to call
type Endpoint interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Routes() http.Handler
}
