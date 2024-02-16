package entity

type Card struct {
	ID           string
	UserID       string        `json:"user_id"`
	Name         string        `json:"name"`
	Title        string        `bson:"title"`
	Photo        string        `bson:"photo"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
	Emails       []Email       `bson:"emails"`
	SocialMedias []SocialMedia `bson:"social_medias"`
	Links        []Link        `bson:"links"`
	Status       CardStatus    `bson:"status"`
}

type PhoneNumber struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type Email struct {
	Title string `bson:"title"`
	Value string `bson:"value"`
}

type Link struct {
	Title string `bson:"title"`
	Value string `bson:"value"`
}

type CardStatus int

const (
	CardStatusActive CardStatus = iota + 1
	CardStatusDeActive
)

func (c CardStatus) IsValid() bool {
	return c >= CardStatusActive && int(c) <= len(CardStatusStrings)
}

var CardStatusStrings = map[CardStatus]string{
	CardStatusActive:   "active",
	CardStatusDeActive: "deActive",
}

func (c CardStatus) String() string {
	return CardStatusStrings[c]
}
