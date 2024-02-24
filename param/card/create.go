package cardparam

import "github.com/tonet-me/tonet-core/entity"

type CreateNewRequest struct {
	AuthenticatedUserID string
	CreateData          CardCreateData `json:"create_data"`
}

type CreateNewResponse struct {
	Card entity.Card `json:"card"`
}

type CardCreateData struct {
	Name         string               `json:"name"`
	Title        string               `json:"title"`
	Photo        string               `json:"photo"`
	PhoneNumbers []entity.PhoneNumber `json:"phone_numbers"`
	Emails       []entity.Email       `json:"emails"`
	SocialMedias []entity.SocialMedia `json:"social_medias"`
	Links        []entity.Link        `json:"links"`
	Status       entity.CardStatus    `json:"status"`
}
