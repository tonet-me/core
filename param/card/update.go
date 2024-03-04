package cardparam

import "github.com/tonet-me/tonet-core/entity"

type UpdateRequest struct {
	AuthenticatedUserID string
	CardID              string         `json:"card_id"`
	UpdateData          CardUpdateData `json:"update_data"`
}

type UpdateResponse struct {
	Updated bool `json:"updated"`
}

type CardUpdateData struct {
	Name         string                `json:"name"`
	Title        string                `json:"title"`
	About        *string               `json:"about"`
	PhotoURL     *string               `json:"photo_url"`
	PhoneNumbers *[]entity.PhoneNumber `json:"phone_numbers"`
	Emails       *[]entity.Email       `json:"emails"`
	SocialMedias *[]entity.SocialMedia `json:"social_medias"`
	Links        *[]entity.Link        `json:"links"`
	Status       entity.CardStatus     `json:"status"`
}
