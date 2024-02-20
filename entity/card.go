package entity

type Card struct {
	ID           string        `json:"id"`
	UserID       string        `bson:"user_id" json:"user_id"`
	Name         string        `bson:"name2" json:"name"`
	Title        string        `bson:"title" json:"title"`
	Photo        string        `bson:"photo" json:"photo"`
	PhoneNumbers []PhoneNumber `bson:"phone_numbers" json:"phoneNumbers"`
	Emails       []Email       `bson:"emails" json:"emails"`
	SocialMedias []SocialMedia `bson:"social_medias" json:"socialMedias"`
	Links        []Link        `bson:"links" json:"links"`
	Status       CardStatus    `bson:"status" json:"status"`
}

type PhoneNumber struct {
	Title string `bson:"title2" json:"title"`
	Value string `bson:"value" json:"value"`
}

type Email struct {
	Title string `bson:"title" json:"title"`
	Value string `bson:"value" json:"value"`
}

type Link struct {
	Title string `bson:"title" json:"title"`
	Value string `bson:"value" json:"value"`
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
