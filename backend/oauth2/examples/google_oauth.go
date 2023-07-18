package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Person struct {
	Name string
}

func googleLogin() {
	var oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/oauth/redirect",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	http.HandleFunc("/oauth/google", func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(oauthConf.Endpoint.AuthURL)
		if err != nil {
			log.Println(err)
		}
		params := url.Values{}
		params.Add("client_id", oauthConf.ClientID)
		params.Add("scope", strings.Join(oauthConf.Scopes, " "))
		params.Add("redirect_uri", oauthConf.RedirectURL)
		params.Add("response_type", "code")
		// params.Add("state", "")
		u.RawQuery = params.Encode()

		http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		if code == "" {
			w.Write([]byte("code not found"))
			return
		}
		// convert oauth code => token
		token, err := oauthConf.Exchange(oauth2.NoContext, code)

		// get user info
		resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		defer resp.Body.Close()

		response, _ := ioutil.ReadAll(resp.Body)
		w.Write(response)
		return
	})

	http.ListenAndServe(":8080", nil)
}
