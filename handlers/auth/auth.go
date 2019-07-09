package auth

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateStateOauthCookie(c *gin.Context) string {
	var expiration = int(24 * 60 * time.Millisecond)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, expiration, "/", "google.com", true, false)

	return state
}
