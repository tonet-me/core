package userparam

import "github.com/tonet-me/tonet-core/entity"

type GetInfoRequest struct {
	AuthenticatedUserID string
}

type GetInfoResponse struct {
	User entity.User `json:"user"`
}
