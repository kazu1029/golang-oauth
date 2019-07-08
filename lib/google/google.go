package google

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

type GoogleConfig struct {
	ClientID     string `envconfig:"GOOGLE_CLIENT_ID"`
	ClientSecret string `envconfig:"GOOGLE_CLIENT_SECRET"`
}

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

func GetConnect() *oauth2.Config {
	var googleConfig GoogleConfig
	envconfig.Process("GOOGLE", &googleConfig)

	config := &oauth2.Config{
		ClientID:     googleConfig.ClientID,
		ClientSecret: googleConfig.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: "http://localhost:8080/auth/google/callback",
	}

	return config
}
