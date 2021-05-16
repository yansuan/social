package social

import (
	"context"
	"github.com/yansuan/social/facebook"
	"github.com/yansuan/social/google"
	"golang.org/x/oauth2"
)

type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verifiedEmail"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
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

	result = &UserInfo{}

	//get user info
	if this.SocialType == Google {
		userInfo, err1 := google.GetUserInfo(this.Token.AccessToken)
		if err1 != nil {
			err = err1
			return
		}

		result.Id = userInfo.Id
		result.Email = userInfo.Email
		result.VerifiedEmail = userInfo.VerifiedEmail
		result.Gender = userInfo.Gender
		result.Name = userInfo.Name
		result.FirstName = userInfo.GivenName
		result.LastName = userInfo.FamilyName
		result.Picture = userInfo.Picture
		return
	}

	if this.SocialType == Facebook {
		userInfo, err1 := facebook.GetUserInfo(this.Token.AccessToken)
		if err1 != nil {
			err = err1
			return
		}

		result.Id = userInfo.Id
		result.Email = userInfo.Email
		result.VerifiedEmail = false
		result.Gender = userInfo.Gender
		result.Name = userInfo.Name
		result.FirstName = userInfo.FirstName
		result.LastName = userInfo.LastName
		if userInfo.Picture != nil {
			result.Picture = userInfo.Picture.Data.URL
		}
		return
	}

	return
}
