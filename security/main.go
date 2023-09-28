package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

type Comment struct {
	Comment string
}

func main() {
	tmpl, err := template.New("pets.html").ParseFiles("pets.html")
	if err != nil {
		panic(err)
	}

	var comments []Comment // null comment store

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		rawData := string(body)
		r.Body.Close()
		if len(rawData) > 0 {
			value := strings.Split(rawData, "=")[1]
			data, _ := url.QueryUnescape(value)
			comments = append(comments, Comment{
				Comment: data,
			})
		}

		err = tmpl.Execute(w, comments)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
