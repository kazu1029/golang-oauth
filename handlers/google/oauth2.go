package controllersGoogle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/lib/google"
)

type Oauth2Handler struct {
	Context *gin.Context
}

type Oauth2Interface interface {
	Get()
}

func (c *gin.Context) Get() {
	config := google.GetConnect()
	url := config.AuthCodeURL("")
	c.Redirect(http.StatusMovedPermanently, url)
}

func GoogleOauth2Handler(c *gin.Context) {
	config := google.GetConnect()
	url := config.AuthCodeURL("")
	c.Redirect(http.StatusMovedPermanently, url)
}
