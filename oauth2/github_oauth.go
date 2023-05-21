package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func githubLogin() {
	// https://www.sohamkamani.com/golang/oauth/
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// We will be using httpclient to make external HTTTP requests later in our code
	httpclient := http.Client{}

	const (
		clientID     = "09f284058a21a54ac468"
		clientSecret = "ef23cf7f2371e32436ba6247d0191eb31592e8c1"
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
