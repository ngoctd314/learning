package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dchest/captcha"
)

func main() {
	store := captcha.NewMemoryStore(1000, time.Minute)
	numCaptcha := newNumCaptcha(store)

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		id, err := numCaptcha.Generate()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println("store", store)
		w.Write([]byte(id))
	})
	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		solution := r.URL.Query().Get("solution")
		if numCaptcha.Verify(id, solution) {
			w.Write([]byte("verify success"))
			return
		}
		w.Write([]byte("invalid"))
	})

	http.HandleFunc("/captcha", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		_, err := captcha.NewImage(id, store.Get(id, false), captcha.StdWidth, captcha.StdHeight).WriteTo(w)
		if err != nil {
			w.Write([]byte("invalid"))
			return
		}
	})

	http.ListenAndServe(":8080", nil)

}
