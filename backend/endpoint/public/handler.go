package publicEndpoint

import (
	"bookmark-backend/common/config"
	"context"
	"github.com/bsthun/gut"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"net/url"
)

type Handler struct {
	config       *config.Config
	OidcProvider *oidc.Provider
	OidcVerifier *oidc.IDTokenVerifier
	Oauth2Config *oauth2.Config
}

func Handle(config *config.Config) *Handler {
	h := &Handler{
		config:       config,
		OidcProvider: nil,
		OidcVerifier: nil,
		Oauth2Config: nil,
	}

	redirectUrl, err := url.JoinPath(*config.FrontendUrl, "/entry/callback")
	if err != nil {
		gut.Fatal("unable to join url path", err)
	}

	h.OidcProvider, err = oidc.NewProvider(context.Background(), *config.AuthEndpoint)
	if err != nil {
		gut.Fatal("unable to create oidc provider", err)
	}

	h.Oauth2Config = &oauth2.Config{
		ClientID:     *config.AuthClientId,
		ClientSecret: *config.AuthClientSecret,
		RedirectURL:  redirectUrl,
		Endpoint:     h.OidcProvider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return h
}
