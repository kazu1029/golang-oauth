package google

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
)

type GoogleConfig struct {
	ClientID     string
	ClientSecret string
}

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

func GetConnect() *oauth2.Config {
	var googleConfig GoogleConfig
	envconfig.Process("google", &googleConfig)

	fmt.Printf("googleConfig is %+v\n", googleConfig)
	fmt.Println(os.Getenv("GOOGLE_CLIENT_ID"))
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: "http://localhost:8080/auth/google/callback",
	}

	return config
}
