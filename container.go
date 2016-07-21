package main

// Legal states for a Container
const (
	STOPPED string = "STOPPED"
	RUNNING string = "RUNNING"
)

// Container is a basic struct for holding information about container.
type Container struct {
	Name  string `json:"name"`
	State string `json:"state"`
	PID   int64  `json:"pid"`
}
