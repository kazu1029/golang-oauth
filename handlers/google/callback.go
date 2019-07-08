package controllersGoogle

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/lib/google"
	v2 "google.golang.org/api/oauth2/v2"
)

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func Callback(c *gin.Context) {
	request := CallbackRequest{}
	ctx := context.Background()
	if err := c.Bind(&request); err != nil {
		log.Println(err)
	}

	config := google.GetConnect()

	tok, err := config.Exchange(ctx, request.Code)
	if err != nil {
		log.Println(err)
	}

	if tok.Valid() == false {
		log.Println("Invalid Token")
	}

	service, _ := v2.New(config.Client(ctx, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()

	c.JSON(200, tokenInfo)
}
