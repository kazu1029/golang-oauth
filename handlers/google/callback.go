package controllersGoogle

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/lib/google"
	v2 "google.golang.org/api/oauth2/v2"
)

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type UserData struct {
	AccessToken string
	Email       string
	UserId      string
}

func Callback(c *gin.Context) {
	oauthState, _ := c.Cookie("oauthstate")
	request := CallbackRequest{}
	if err := c.Bind(&request); err != nil {
		log.Println(err)
	}
	fmt.Printf("oauthState is %v\n", oauthState)
	// if request.State != oauthState {
	// 	log.Println("invalid oauth google state")
	// 	c.Redirect(http.StatusMovedPermanently, "/")
	// 	return
	// }

	data, err := GetUserDataFromGoogle(request.Code, c)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	c.JSON(200, data)
}

func GetUserDataFromGoogle(code string, c *gin.Context) (UserData, error) {
	ctx := context.Background()
	config := google.GetConnect()

	tok, err := config.Exchange(ctx, code)
	if err != nil {
		log.Println(err)
	}

	if tok.Valid() == false {
		log.Println("Invalid Token")
	}

	service, _ := v2.New(config.Client(ctx, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()

	data := UserData{
		AccessToken: tok.AccessToken,
		Email:       tokenInfo.Email,
		UserId:      tokenInfo.UserId,
	}

	return data, nil
}
