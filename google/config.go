package google

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"}
)

const (
	url = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

/*
{
  "id": "12345678",
  "email": "hchlin@gmail.com",
  "verified_email": true,
  "name": "hu chenglin",
  "given_name": "hu",
  "family_name": "chenglin",
  "picture": "https://lh3.googleusercontent.com/a/AATXAJyLI2vOYsq4bGKu_jp_7UKcvdSIyVFZmZ3sTPIv=s96-c",
  "locale": "zh-CN"
}
 */
type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Gender        string `json:"gender"`
	Locale        string `json:"locale"`
}

func GetUserInfo(accessToken string) (result *UserInfo, err error) {
	response, err := http.Get(url + accessToken)
	defer response.Body.Close()
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		err = err1
		return
	}
	log.Println(string(body))
	result = &UserInfo{}
	err = json.Unmarshal(body, result)
	return
}

func NewConfig(clientId, clientSecret, redirectUrl string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
	}
}
