package cardparam

import "github.com/tonet-me/tonet-core/entity"

type CreateNewRequest struct {
	AuthenticatedUserID string
	CreateData          CardCreateData
}

type CreateNewResponse struct {
	Card entity.Card
}

type CardCreateData struct {
	Name         string               `bson:"name"`
	Title        string               `bson:"title"`
	Photo        string               `bson:"photo"`
	PhoneNumbers []entity.PhoneNumber `bson:"phone_numbers"`
	Emails       []entity.Email       `bson:"emails"`
	SocialMedias []entity.SocialMedia `bson:"social_medias"`
	Links        []entity.Link        `bson:"links"`
	Status       entity.CardStatus    `bson:"status"`
}
