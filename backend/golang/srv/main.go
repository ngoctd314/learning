package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "close")
		w.Write([]byte("pong"))
	})
	var f *os.File
	f.Sync()
	http.ListenAndServe(":8080", nil)
}
