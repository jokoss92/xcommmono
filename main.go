package main

import (
	"log"
	"time"
	"xcomm/route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Printf("main: starting HTTP server")
	srv := route.StartHTTPServer()

	log.Printf("main: serving for 2 minutes")
	time.Sleep(120 * time.Second)
	log.Printf("main: stopping HTTP server")

	// now close the server gracefully ("shutdown")
	// timeout could be given instead of nil as a https://golang.org/pkg/context/
	if err := srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

	log.Printf("main: done. exiting")
}
