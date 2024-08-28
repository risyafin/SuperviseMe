package domain

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	OAuthGoogleConf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	OAuthStateString = ""
)

func InitGoogleConfig() {
	OAuthGoogleConf.ClientID = os.Getenv("CLIENT_ID")
	OAuthGoogleConf.ClientSecret = os.Getenv("CLIENT_SECRET")
	OAuthGoogleConf.RedirectURL = os.Getenv("REDIRECT_URL")
	OAuthStateString = os.Getenv("STATE_STRING")
}
