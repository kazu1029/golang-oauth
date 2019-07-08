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

	r.GET("/auth/google/login", controllersGoogle.Oauth2Handler)
	r.GET("/auth/google/callback", controllersGoogle.Callback)

	return r
}
