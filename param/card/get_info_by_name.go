package cardparam

import "github.com/tonet-me/tonet-core/entity"

type GetInfoByNameRequest struct {
	Name string `json:"card_name"`
}

type GetInfoByNameResponse struct {
	Card entity.Card `json:"card"`
}
