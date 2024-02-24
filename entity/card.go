package entity

type Card struct {
	ID           string        `bson:"_id,omitempty" json:"id"`
	UserID       string        `bson:"user_id" json:"user_id"`
	Name         string        `bson:"name,omitempty" json:"name"`
	Title        string        `bson:"title,omitempty" json:"title"`
	Photo        string        `bson:"photo,omitempty" json:"photo"`
	PhoneNumbers []PhoneNumber `bson:"phone_numbers,inline" json:"phoneNumbers"`
	Emails       []Email       `bson:"emails,inline" json:"emails"`
	SocialMedias []SocialMedia `bson:"social_medias,inline" json:"socialMedias"`
	Links        []Link        `bson:"links,inline" json:"links"`
	Status       CardStatus    `bson:"status,omitempty" json:"status"`
}

type PhoneNumber struct {
	Title string `bson:"title" json:"title"`
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
