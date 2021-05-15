package social

import (
	"context"
	"fmt"
	"github.com/yansuan/social/facebook"
	"github.com/yansuan/social/google"
	"golang.org/x/oauth2"
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

type Client struct {
	SocialType Type
	Config     *oauth2.Config
	Token      *oauth2.Token
}

func (this *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return this.Config.AuthCodeURL(state, opts...)
}

func (this *Client) exchange(code string) (err error) {
	this.Token, err = this.Config.Exchange(context.TODO(), code)
	return
}

func (this *Client) GetUserInfo(code string) (result *UserInfo, err error) {
	err = this.exchange(code)
	if err != nil {
		return
	}

	//get user info
	if this.SocialType == Google {
		userinfo, err1 := google.GetUserInfo(this.Token.AccessToken)
		if err1 != nil {
			err = err1
			return
		}

		fmt.Printf("%+v\n", userinfo)
	}

	if this.SocialType == Facebook {
		userinfo, err1 := facebook.GetUserInfo(this.Token.AccessToken)
		if err1 != nil {
			err = err1
			return
		}
		fmt.Printf("%+v\n", userinfo)
	}

	return
}
