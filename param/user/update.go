package userparam

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
	Updated bool
}
