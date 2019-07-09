package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateStateOauthCookie(c *gin.Context) string {
	var expiration = int(24 * 60 * time.Millisecond)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, expiration, "/", "google.com", true, false)
	cookie, _ := c.Request.Cookie("oauthstate")
	fmt.Printf("cookie is %v\n", cookie)

	return state
}
