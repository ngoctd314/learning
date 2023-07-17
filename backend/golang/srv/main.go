package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * 100)
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}
