package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Some very expensive database call
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}
