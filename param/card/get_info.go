package cardparam

import "github.com/tonet-me/tonet-core/entity"

type GetInfoRequest struct {
	AuthenticatedUserID string
	CardID              string `json:"card_id"`
}

type GetInfoResponse struct {
	Card entity.Card `json:"card"`
}
