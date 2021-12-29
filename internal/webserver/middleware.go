package webserver

import (
	"log"
	"net/http"
)

//statusRecorder is a customization of the http responsewriter implimentation
// tro enable middleware to access response codes. By usiong composition we only
//need to preload one method, WriteHeader
type statusRecorder struct {
	http.ResponseWriter
	status int
}

//this is the only customization to the default http ResponseWriter, we record
//the status code
func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

//LogCall is an example middleware function to log all requests
func LogCall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//initialize to 200 in case WriteHeader isn't called explicitly
		wreq := statusRecorder{w, 200}

		log.Printf("RECIEVED %s request FOR %s FROM %s\n",
			r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(&wreq, r)
		log.Printf("RETURNED HTTP Code %d FOR %s request FOR %s FROM %s\n",
			wreq.status, r.Method, r.URL.Path, r.RemoteAddr)
	})
}

//StdHeaders is an example middleware function to apply basic HTTP Headers
func StdHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application//json; charset=UTF-8")
		next.ServeHTTP(w, r)
	})
}
