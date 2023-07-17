package main

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

// https://www.codershood.info/2020/04/16/facebook-login-in-golang-tutorial/

func getFacebookOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("FACEBOOK_APP_ID"),
		ClientSecret: os.Getenv("FACEBOOK_SECRET"),
		Endpoint:     facebook.Endpoint,
		RedirectURL:  "http://localhost:8080/oauth/redirect",
		Scopes:       []string{"email"},
	}
}
