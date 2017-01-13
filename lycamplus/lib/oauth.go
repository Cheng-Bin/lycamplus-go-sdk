package lib

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

// LycamPlusOAuth2 Object .
type LycamPlusOAuth2 struct {
	appKey    string
	appSecret string
}

// NewLycamPlusOAuth2 create a OAuth instance.
func NewLycamPlusOAuth2(appKey, appSecret string) *LycamPlusOAuth2 {
	return &LycamPlusOAuth2{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

// OAuth .
func (that *LycamPlusOAuth2) OAuth(username, password string) (*oauth2.Token, error) {
	ctx := context.Background()
	conf := that.getConfig()
	token, err := conf.PasswordCredentialsToken(ctx, username, password)

	if err != nil {
		return nil, err
	}

	return token, nil
}

// getConfig by structure field and constant .
func (that *LycamPlusOAuth2) getConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     that.appKey,
		ClientSecret: that.appSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: fmt.Sprintf("%s%s", DefaultOAuth2URL, DefaultTokenPath),
		},
	}
}
