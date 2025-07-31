package utils

import (
	"context"
	"os"

	"golang.org/x/oauth2"
	// ganti jika pakai provider lain
)

var OAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
	RedirectURL:  "http://localhost:8083/oauth2/callback",
	Scopes:       []string{"read", "write"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "http://localhost:8081/oauth/authorize",
		TokenURL: "http://localhost:8081/oauth/token",
	},
}

func GetAuthURL(state string) string {
	return OAuthConfig.AuthCodeURL(state)
}

func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return OAuthConfig.Exchange(ctx, code)
}
