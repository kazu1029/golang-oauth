package controllersFacebook

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	fb "github.com/huandu/facebook"
	"github.com/kazu1029/golang-oauth/lib/facebook"
	"golang.org/x/oauth2"
)

type CallbackRequest struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type UserData struct {
	AccessToken string
	Email       string
	UserId      string
	Name        string
}

func Callback(c *gin.Context) {
	oauthState, _ := c.Cookie("oauthstate")
	request := CallbackRequest{}
	if err := c.Bind(&request); err != nil {
		log.Println(err)
	}
	fmt.Printf("oauthState is %+v\n", oauthState)

	data, err := GetUserDataFromFacebook(c, request.Code)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	c.JSON(200, data)
}

func GetUserDataFromFacebook(c *gin.Context, code string) (UserData, error) {
	config := facebook.GetConnect()

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println(err)
	}

	if tok.Valid() == false {
		log.Println("Invalid Token")
	}

	client := config.Client(oauth2.NoContext, tok)
	session := &fb.Session{
		Version:    "v2.8",
		HttpClient: client,
	}

	res, err := session.Get("/me?fields=id,name,email,picture", nil)
	if err != nil {
		log.Fatal(err)
	}
	data := UserData{
		AccessToken: tok.AccessToken,
		Email:       res["email"].(string),
		UserId:      res["id"].(string),
		Name:        res["name"].(string),
	}

	return data, nil
}
