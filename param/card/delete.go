package cardparam

type DeleteRequest struct {
	AuthenticatedUserID string
	CardID              string `json:"card_id"`
}
type DeleteResponse struct {
	Success bool `json:"success"`
}
