package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server is the core implementation of the HTTP server
type Server struct {
	debug      bool
	port       int
	router     *mux.Router
	containers map[string]Container
	nextPID    int64
}

func (s *Server) run() {
	//TODO: Any final initialization?

	s.router = mux.NewRouter().StrictSlash(false)

	// Generate routes for each API
	for _, a := range apis {
		s.createRoute(a)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router))
}

func (s *Server) createRoute(a API) {
	uri := fmt.Sprintf("/%s", a.name)

	s.router.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		//TODO: Add debug output

		var resp Response
		resp = NotImplemented(fmt.Errorf("%s does not support %s", r.URL, r.Method))

		//TODO: Handle method appropriately

		if err := resp.Render(w); err != nil {
			if err := InternalError(err).Render(w); err != nil {
				panic("Error while rendering internal error, bailing")
			}
		}
	})
}
