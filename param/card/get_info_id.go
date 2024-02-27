package cardparam

import "github.com/tonet-me/tonet-core/entity"

type GetInfoByIDRequest struct {
	AuthenticatedUserID string
	CardID              string `json:"card_id"`
}

type GetInfoByIDResponse struct {
	Card entity.Card `json:"card"`
}
