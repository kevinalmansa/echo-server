package webserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/kevinalmansa/echo-server/internal/endpoints"
)

//NewRestServer is a factory method returning a new WebServer object
func NewRestServer(addr string) WebServer {
	mux := http.NewServeMux()

	server := &RestServer{
		mux:    mux,
		server: &http.Server{Addr: addr, Handler: mux},
	}

	endpointEcho := endpoints.NewEndpointEcho()
	server.mux.Handle("/", LogCall(endpointEcho.Routes()))

	endpointHealthz := endpoints.NewEndpointHealthz()
	server.mux.Handle("/healthz", LogCall(endpointHealthz.Routes()))
	return server
}

//RestServer is a REST implimentation of WebServer
type RestServer struct {
	mux    *http.ServeMux
	server *http.Server
}

//Mux is a getter for the http.ServeMux used by RestServer
func (s *RestServer) Mux() *http.ServeMux {
	return s.mux
}

//Start starts the webserver. It can be used to initialize anything required
// before started and start listening on the configured interface & port.
func (s *RestServer) Start() {
	log.Println("Starting webserver at ", s.server.Addr)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Webserver Error: ", err.Error())
		}
	}()
}

//Shutdown gracefully shutsdown the webserver. Will stop recieving new requests and
//terminate anything else needed
func (s *RestServer) Shutdown() {
	log.Println("Shutdown request recieved")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		// extra handling here
		log.Println("Gracefully shutting down...")
		cancel()
	}()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Println("Error shutting down webserver: ", err.Error())
	}
}
