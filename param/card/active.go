package cardparam

type ActiveRequest struct {
	AuthenticatedUserID string
	CardID              string `json:"card_id"`
}
type ActiveResponse struct {
	Success bool `json:"success"`
}
