package controllersGoogle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/lib/google"
)

type Oauth2Handler struct {
	Ctx *gin.Context
}

type Oauth2Interface interface {
	Get()
}

func (c *Oauth2Handler) Get() {
	config := google.GetConnect()
	url := config.AuthCodeURL("")
	c.Ctx.Redirect(http.StatusMovedPermanently, url)
}

func GoogleOauth2Handler(c Oauth2Handler) {
	config := google.GetConnect()
	url := config.AuthCodeURL("")
	c.Ctx.Redirect(http.StatusMovedPermanently, url)
}
