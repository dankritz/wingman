package main

import (
	"net/http"
	"log"
)

func quote(w http.ResponseWriter, r *http.Request) {

}

func HTTPServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/quote", quote)

	log.Fatal(http.ListenAndServe(":9091", mux))
}