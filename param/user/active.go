package userparam

type ActiveRequest struct {
	AuthenticatedUserID string
}
type ActiveResponse struct {
	Success bool `json:"success"`
}
