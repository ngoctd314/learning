package main

import "net/http"

func main() {
	http.HandleFunc("/api/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", nil)
}
