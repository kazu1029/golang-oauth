package facebook

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

type FacebookConfig struct {
	ClientID     string `envconfig:"FACEBOOK_CLIENT_ID"`
	ClientSecret string `envconfig:"FACEBOOK_CLIENT_SECRET"`
}

const (
	authorizeEndpoint = "https://www.facebook.com/dialog/oauth"
	tokenEndpoint     = "https://graph.facebook.com/oauth/access_token"
)

func GetConnect() *oauth2.Config {
	var facebookConfig FacebookConfig
	envconfig.Process("FACEBOOK", &facebookConfig)

	config := &oauth2.Config{
		ClientID:     facebookConfig.ClientID,
		ClientSecret: facebookConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"email"},
		RedirectURL: "http://localhost:8080/auth/facebook/callback",
	}

	return config
}
