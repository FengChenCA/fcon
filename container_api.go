package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func containerGet(s *Server, r *http.Request) Response {
	name := mux.Vars(r)["name"]

	//TODO: Query container
}

func containerPut(s *Server, r *http.Request) Response {
	name := mux.Vars(r)["name"]

	//TODO: Update container
}

func containerDelete(s *Server, r *http.Request) Response {
	name := mux.Vars(r)["name"]

	//TODO: Delete container
}
