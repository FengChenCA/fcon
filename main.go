package main

import (
	"flag"
	"fmt"
)

func main() {
	version := flag.Bool("version", false, "Print the version and exit")
	//TODO: Add `debug` and `port` flags

	flag.Parse()

	if *version {
		fmt.Printf("FCON 1.0")
		return
	}

	//TODO: Initialize Server, assign to `s`
	s.run()
}
