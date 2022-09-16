package api

import (
	"github.com/ProjectSegfault/segfautils/config"
)

var (
	clientID     = config.OAuthClientID()
	clientSecret = config.OAuthClientSecret()
	redirectURL  = config.OAuthRedirectURL()
	authURL      = config.OAuthURL()
)

func LoginOAuth() {
}
