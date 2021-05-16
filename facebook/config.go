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
	url = "https://graph.facebook.com/me?fields=id,name,first_name,last_name,picture,email&access_token="
)

type UserInfo struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Birthday  string   `json:"birthday"`
	Gender    string   `json:"gender"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Picture   *Picture `json:"picture,omitempty"`
}

// Picture structure
type Picture struct {
	Data struct {
		Height       int    `json:"height"`
		IsSilhouette bool   `json:"is_silhouette"`
		URL          string `json:"url"`
		Width        int    `json:"width"`
	} `json:"data"`
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
