package route

import (
	"log"
	"net/http"
)

func StartHTTPServer() *http.Server {
	srv := &http.Server{Addr: ":8898"}

	Routing()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	// returning reference so caller can call Shutdown()
	return srv
}
