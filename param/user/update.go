package userparam

type UserUpdateData struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	ProfilePhotoURL string `json:"profile_photo_url"`
}
type UpdateRequest struct {
	AuthenticatedUserID string
	UpdateData          UserUpdateData `json:"update_data"`
}

type UpdateResponse struct {
	Updated bool `json:"updated"`
}
