package handlers

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/handlers/google"
	"github.com/kazu1029/golang-oauth/handlers/facebook"
)

func New() http.Handler {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("templates", false)))

	r.GET("/auth/google/login", controllersGoogle.Oauth2Handler)
	r.GET("/auth/google/callback", controllersGoogle.Callback)
	r.GET("/auth/facebook/login", controllersFacebook.Oauth2Handler)
	r.GET("/auth/facebook/callback", controllersFacebook.Callback)

	return r
}
