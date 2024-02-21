package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

type GoogleClaim struct {
	Email      string `json:"email"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
}

type GoogleConfig struct {
	ClientID     string          `koanf:"client_id"`
	ClientSecret string          `koanf:"client_secret"`
	Scopes       []string        `koanf:"scopes"`
	Endpoint     oauth2.Endpoint `koanf:"end_point"`
}

type Google struct {
	oauthConfig *oauth2.Config
}

func NewGoogle(cfg GoogleConfig) Google {
	return Google{
		oauthConfig: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Endpoint:     cfg.Endpoint,
			Scopes:       cfg.Scopes,
		},
	}
}

func (g Google) ValidationToken(ctx context.Context, token string) (*entity.OAuthUserInfo, error) {
	const op = richerror.OP("oauth.google.ValidationAndGetInfoFromToken")

	payload, vErr := idtoken.Validate(context.Background(), token, "")
	if vErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindForbidden),
			richerror.WithMessage(fmt.Sprintf("code vlidation error: %v", vErr)),
		)
	}

	userInfoData, mErr := json.Marshal(payload.Claims)
	if mErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithMessage(fmt.Sprintf("can't marshal user info from token's payload: %v", mErr)),
		)
	}
	oAuthUserInfo, uErr := mapToGoogleClaim(userInfoData)
	if uErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithInnerError(uErr),
		)
	}

	return oAuthUserInfo, nil
}
