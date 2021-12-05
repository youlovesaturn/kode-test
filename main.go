package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := NewServer()
	mux.HandleFunc("/note/", server.noteHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
