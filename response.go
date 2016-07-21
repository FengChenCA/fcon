package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response interface {
	Render(w http.ResponseWriter) error
	String() string
}

// JSON response
type jsonResponse struct {
	code int
	body interface{}
}

func (r *jsonResponse) String() string {
	return http.StatusText(r.code)
}

func (r *jsonResponse) Render(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)

	return json.NewEncoder(w).Encode(r.body)
}

// Error response
type errorResponse struct {
	code int
	msg  string
}

func (r *errorResponse) String() string {
	return fmt.Sprintf("%s: %s", http.StatusText(r.code), r.msg)
}

func (r *errorResponse) Render(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(r.code)
	fmt.Fprintln(w, r.String())

	return nil
}

/* Standard error responses */
func NotImplemented(err error) Response {
	return &errorResponse{http.StatusNotImplemented, err.Error()}
}
func NotFound(err error) Response {
	return &errorResponse{http.StatusNotFound, err.Error()}
}
func Forbidden(err error) Response {
	return &errorResponse{http.StatusForbidden, err.Error()}
}
func BadRequest(err error) Response {
	return &errorResponse{http.StatusBadRequest, err.Error()}
}
func InternalError(err error) Response {
	return &errorResponse{http.StatusInternalServerError, err.Error()}
}
func Conflict(err error) Response {
	return &errorResponse{http.StatusConflict, err.Error()}
}
