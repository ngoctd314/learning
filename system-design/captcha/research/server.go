package main

import (
	"net/http"

	"github.com/dchest/captcha"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
	// 	numCaptcha := numcaptcha{}
	// 	id, err := numCaptcha.Generate()
	// 	if err != nil {
	// 		w.Write([]byte(err.Error()))
	// 		return
	// 	}
	// 	fmt.Println("store", store)
	// 	w.Write([]byte(id))
	// })
	// http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.URL.Query().Get("id")
	// 	solution := r.URL.Query().Get("solution")
	// 	if numCaptcha.Verify(id, solution) {
	// 		w.Write([]byte("verify success"))
	// 		return
	// 	}
	// 	w.Write([]byte("invalid"))
	// })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := captcha.NewImage("", captcha.RandomDigits(6), captcha.StdWidth, captcha.StdHeight).WriteTo(w)
		if err != nil {
			w.Write([]byte("invalid"))
			return
		}
	})

	http.ListenAndServe(":8080", nil)

}
