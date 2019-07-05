package controllersGoogle

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/lib/google"
	v2 "google.golang.org/api/oauth2/v2"
)

type CallbackHandler struct {
	gin.Context
}

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

func (c *CallbackHandler) Get() {
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

	fmt.Printf("email is %+v\n", tokenInfo.Email)
	fmt.Printf("UserId is %+v\n", tokenInfo.UserId)
}
