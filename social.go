package social

import (
	"github.com/yansuan/social/facebook"
	"github.com/yansuan/social/g"
	"github.com/yansuan/social/google"
	"golang.org/x/oauth2"
	"log"
)

type Type int

const (
	// Facebook type
	Facebook Type = iota

	// Google Type
	Google
)

func (t Type) String() string {
	switch t {
	case Facebook:
		return "Facebook"
	case Google:
		return "Google"
	default:
		log.Fatal("unknown social media")
	}

	return ""
}

func New(socialType Type, clientId, clientSecret, redirectUrl string) (c *Client) {
	c = &Client{}
	config := &oauth2.Config{}
	if socialType == Google {
		config = google.NewConfig(clientId, clientSecret, redirectUrl)
	} else if socialType == Facebook {
		config = facebook.NewConfig(clientId, clientSecret, redirectUrl)
	}

	c.Config = config
	c.SocialType = socialType
	return
}

func Debug(isDebug bool) {
	g.Debug = isDebug
}
