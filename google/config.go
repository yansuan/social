package google

import (
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
)

var (
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"}
)

const (
	url = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
)

type UserInfo struct {
	//"id": "114512230444013345330",
	//"email": "wangshubo1989@126.com",
	//"verified_email": true,
	//"name": "王书博",
	//"given_name": "书博",
	//"family_name": "王",
	//"picture": "https://lh3.googleusercontent.com/-XdUIqdMkCWA/AAAAAAAAAAI/AAAAAAAAAAA/4252rscbv5M/photo.jpg",
	//"locale": "zh-CN"
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
		Endpoint:     google.Endpoint,
	}
}
