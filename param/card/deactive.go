package cardparam

type DeActiveRequest struct {
	AuthenticatedUserID string
	CardID              string `json:"card_id"`
}
type DeActiveResponse struct {
	Success bool `json:"success"`
}
