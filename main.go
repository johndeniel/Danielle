package main

import (
	"log"
	"net/http"

	"github.com/johndeniel/danielle/server"
)

func main() {
	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/article/", server.ReadHandler)

	log.Fatal(http.ListenAndServe(":7860", nil))
}
