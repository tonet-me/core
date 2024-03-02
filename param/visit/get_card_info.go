package visitparam

import "github.com/tonet-me/tonet-core/entity"

type GetCardInfoByNameRequest struct {
	Name string `json:"card_name"`
}

type GetCardInfoByNameResponse struct {
	Card entity.Card `json:"card"`
}
