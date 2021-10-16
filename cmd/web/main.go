package main

import (
	"GO-CHAT/internal/handlers"
	"log"
	"net/http"
)

func main() {

	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting on :8080")

	_ = http.ListenAndServe(":8080", mux)

}
