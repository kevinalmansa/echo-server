package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/kevinalmansa/echo-server/internal/webserver"
)

var Version = "development"
var Build = "dev"

type args struct {
	addr string
}

func argParse() *args {
	args := &args{}

	flag.StringVar(&args.addr, "addr", ":8080", "Address for webserver to listen on")
	flag.Parse()
	return args
}

func main() {
	args := argParse()

	log.Printf("%s Version %s Build %s\n",
		filepath.Base(os.Args[0]), Version, Build)
	server := webserver.NewRestServer(args.addr)

	//Catch OS Signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// start webserver
	server.Start()

	//wait for SIGINT or SIGTERM
	<-quit
	server.Shutdown()
}
