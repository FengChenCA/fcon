package main

import (
	"net/http"
)

// API is the basic structure for every API call.
type API struct {
	name   string
	get    func(s *Server, r *http.Request) Response
	put    func(s *Server, r *http.Request) Response
	post   func(s *Server, r *http.Request) Response
	delete func(s *Server, r *http.Request) Response
}

var containersAPI = API{
	name: "containers",
	//TODO: Add containers endpoints
}

var containerAPI = API{
	name: "containers/{name}",
	//TODO: Add container endpoints
}

var apis = []API{
	containersAPI,
	containerAPI,
}
