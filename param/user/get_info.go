package userparam

import "tonet-core/entity"

type GetInfoRequest struct {
	AuthenticatedUserID string
}

type GetInfoResponse struct {
	User entity.User
}
