package facebook

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io/ioutil"
	"net/http"
)

var (
	scopes = []string{"public_profile", "email"}
)

const (
	url = "https://graph.facebook.com/me?access_token="
)

type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
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
		Endpoint:     facebook.Endpoint,
	}
}
