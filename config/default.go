package config

import (
	oauth "github.com/tonet-me/tonet-core/adapter/oauth"
	"golang.org/x/oauth2/google"
)

func Default() Config {
	cfx := Config{
		OAuth: oauth.Config{
			Google: oauth.GoogleConfig{
				Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email"},
				Endpoint: google.Endpoint,
			},
		},
	}

	return cfx
}
