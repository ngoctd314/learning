package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("~ping")
		// time.Sleep(time.Millisecond * 500)
		// w.WriteHeader(http.StatusNoContent)
		for i := 0; i < 10; i++ {
			// time.Sleep(time.Millisecond * 100)
			w.Write([]byte("pong"))
		}
	})
	var f *os.File
	f.Sync()
	http.ListenAndServe(":8080", nil)
}
