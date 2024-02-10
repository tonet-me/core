package userparam

import "github.com/tonet-me/tonet-core/entity"

type UserUpdateData struct {
	FirstName       string
	LastName        string
	PhoneNumber     string
	ProfilePhotoURL string
}
type UpdateRequest struct {
	AuthenticatedUserID string
	UpdateData          UserUpdateData
}

type UpdateResponse struct {
	User entity.User
}
