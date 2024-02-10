package userparam

import (
	"tonet-core/entity"
)

type LoginOrRegisterRequest struct {
	Token string
}
type LoginOrRegisterResponse struct {
	User   entity.User
	Tokens Tokens
}
