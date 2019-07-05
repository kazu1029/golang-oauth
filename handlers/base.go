package handlers

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/handlers/google"
)

func New() http.Handler {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("templates", false)))

	// r.GET("/auth/google/oauth2", func(c *gin.Context) {
	// 	config := google.GetConnect()
	// 	url := config.AuthCodeURL("")
	// 	c.Redirect(http.StatusMovedPermanently, url)
	// })

	 r.GET("/auth/google/oauth2", controllersGoogle.GoogleOauth2Handler)
	 r.GET("auth/gogle/oauth2_2", func(c *gin.Context) { &controllersGoogle.Oauth2Interface(c) })
	// r.GET("/auth/google/callback", &controllersGoogle.CallbackHandler{})

	return r
}
