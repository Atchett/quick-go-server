package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", "/", "the directory of static file to host")
	flag.Parse()

	fs := http.FileServer(http.Dir(*directory))
	http.Handle("/", fs)

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

// /Users/johnspurgin/Documents/Git/Grid/Arthaus
