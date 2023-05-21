package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func githubLogin() {
	// https://www.sohamkamani.com/golang/oauth/
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// We will be using httpclient to make external HTTTP requests later in our code
	httpclient := http.Client{}

	var (
		clientID     = os.Getenv("GITHUB_CLIENT_ID")
		clientSecret = os.Getenv("GITHUB_SECRET")
	)

	// Adding a Redirect Route
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		// First, we need to get the value of the code query param
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		code := r.FormValue("code")

		// call the github oauth endpoint to get our access token
		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// We set this header since we want the response
		// as JSON
		req.Header.Set("accept", "application/json")

		// Send out the HTTP request
		res, err := httpclient.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		// Parse the request body
		var t oauthAccessResponse
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// redirect to welcome page
		w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
		w.WriteHeader(http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}

type oauthAccessResponse struct {
	AccessToken string `json:"access_token"`
}
