package userparam

import (
	"github.com/tonet-me/tonet-core/entity"
)

type LoginOrRegisterRequest struct {
	Token      string           `json:"token"`
	ProviderID entity.OAuthType `json:"provider_id"`
}
type LoginOrRegisterResponse struct {
	User    entity.User `json:"user"`
	Tokens  Tokens      `json:"tokens"`
	NewUser bool        `json:"new_user"`
}
