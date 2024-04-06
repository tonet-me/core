package cardparam

import "github.com/tonet-me/tonet-core/entity"

type GetAllUserCardsRequest struct {
	AuthenticatedUserID string
}

type GetAllUserCardsResponse struct {
	Cards []entity.Card `json:"cards"`
}
