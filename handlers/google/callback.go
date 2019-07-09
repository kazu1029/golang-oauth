package controllersGoogle

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kazu1029/golang-oauth/lib/google"
	// v2 "google.golang.org/api/oauth2/v2"
)

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func Callback(c *gin.Context) {
	oauthState, _ := c.Cookie("oauthstate")
	request := CallbackRequest{}
	if err := c.Bind(&request); err != nil {
		log.Println(err)
	}
	fmt.Printf("oauthState is %v\n", oauthState)

	if request.State != oauthState {
		log.Println("invalid oauth google state")
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	data, err := GetUserDataFromGoogle(request.Code)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	c.JSON(200, data)
	// request := CallbackRequest{}
	// ctx := context.Background()
	// if err := c.Bind(&request); err != nil {
	// 	log.Println(err)
	// }

	// config := google.GetConnect()

	// tok, err := config.Exchange(ctx, request.Code)
	// if err != nil {
	// 	log.Println(err)
	// }

	// if tok.Valid() == false {
	// 	log.Println("Invalid Token")
	// }

	// service, _ := v2.New(config.Client(ctx, tok))
	// tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(ctx).Do()

	// c.JSON(200, tokenInfo)
}

func GetUserDataFromGoogle(code string) ([]byte, error) {
	config := google.GetConnect()
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wronge: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
