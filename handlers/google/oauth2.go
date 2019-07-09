package controllersGoogle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/handlers/auth"
	"github.com/kazu1029/golang-oauth/lib/google"
)

func Oauth2Handler(c *gin.Context) {
	config := google.GetConnect()
	oauthState := auth.GenerateStateOauthCookie(c)

	url := config.AuthCodeURL(oauthState)
	c.Redirect(http.StatusMovedPermanently, url)
}
