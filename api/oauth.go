package oauth

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"x/oauth2"
	"github.com/ProjectSegfault/segfautils/config"
	"github.com/goccy/go-json"
)

var (
	clientID     = config.OAuthClientID()
	clientSecret = config.OAuthClientSecret()
	redirectURL  = config.RedirectURL()
	authURL		 = config.AuthURL()
)

func LoginOAuth() {
// Create a new redirect route route
http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	code := r.FormValue("code")
		reqURL := fmt.Sprintf("%s/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", authURL, clientID, clientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}	
		req.Header.Set("accept", "application/json")
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()
	var t OAuthAccessResponse
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	w.Header().Set("Location", "/announcements/?access_token="+t.AccessToken)
	w.WriteHeader(http.StatusFound)	
	}
	
	type OAuthAccessResponse struct {
		AccessToken string `json:"access_token"`
	}
