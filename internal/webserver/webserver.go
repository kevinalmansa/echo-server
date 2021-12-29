package webserver

import "net/http"

//WebServer is an interface for the webservice used.
type WebServer interface {
	Mux() *http.ServeMux
	Start()
	Shutdown()
}
