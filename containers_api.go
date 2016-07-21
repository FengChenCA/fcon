package main

import "net/http"

func containersGet(s *Server, r *http.Request) Response {
	return &jsonResponse{http.StatusOK, s.containers}
}

func containersPost(s *Server, r *http.Request) Response {

	//TODO: Create new container
}
