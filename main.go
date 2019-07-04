package main

import (
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
  "log"
	"github.com/gin-gonic/gin"
)

const (
	authorizeEndpoint = "https://www.facebook.com/dialog/oauth"
	tokenEndpoint     = "https://graph.facebook.com/oauth/access_token"
)

var facebookOAuthConfig *oauth2.Config

type FacebookEnv struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func init() {
	var facebookEnv FacebookEnv
	envconfig.Process("", &facebookEnv)
	facebookOAuthConfig := &oauth2.Config{
		ClientID:     facebookEnv.ClientID,
		ClientSecret: facebookEnv.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"email"},
		RedirectURL: facebookEnv.RedirectURL,
	}
}

func main() {
	tok, err := facebookOAuthConfig.Exchange(oauth2.NoContext, request.Code)
	if err != nil {
		log.Fatal(err)
	}

	if tok.Valid() == false {
		log.Fatal(error.New("this token is invalid"))
	}

	client := facebookOAuthConfig.Client(oauth2.NoContext, tok)
	session := &fb.Session{
		Version: "v2.8",
		HttpClient: client,
	}

	res, err := session.Get("/me?fields=id,name,email", nil)
	if err != nil {
		log.Fatal(err)
	}

	c.Data["ID"] = res["id"]
	c.Data["Name"] = res["name"]
	c.Data["Email"] = res["email"]

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(200, "index.html", gin.H{})
	})
}

