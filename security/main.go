package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Query struct {
	Question string
}

// password = ” or 1 = 1--'
func main() {
	tmpl, err := template.New("reflected_xss.html").ParseFiles("reflected_xss.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if len(strings.TrimSpace(query)) == 0 {
			query = "Unknown"
		}
		err = tmpl.Execute(w, Query{query})
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/csrf", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("header", r.Header)
		fmt.Println("cookie", r.Cookies())
		w.Write([]byte("ok"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
