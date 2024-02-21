package entity

type OAuthUserInfo struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	ProfilePhotoURL string `json:"profile_photo_url"`
}
