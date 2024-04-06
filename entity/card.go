package entity

import (
	"time"
)

type Card struct {
	ID           string        `bson:"_id,omitempty" json:"id"`
	UserID       string        `bson:"user_id,omitempty" json:"user_id"`
	Name         string        `bson:"name" json:"name"`
	Title        string        `bson:"title" json:"title"`
	About        string        `bson:"about" json:"about"`
	PhotoURL     string        `bson:"photo_url" json:"photo_url"`
	PhoneNumbers []PhoneNumber `bson:"phone_numbers" json:"phone_numbers"`
	Emails       []Email       `bson:"emails" json:"emails"`
	SocialMedias []SocialMedia `bson:"social_medias" json:"social_medias"`
	Links        []Link        `bson:"links" json:"links"`
	Status       CardStatus    `bson:"status" json:"status"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at" json:"updated_at"`
}

type PhoneNumber struct {
	Title string           `bson:"title" json:"title"`
	Value PhoneNumberValue `bson:"value,inline" json:"value"`
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
	CardStatusDelete
)

func (c CardStatus) IsValid() bool {
	return c >= CardStatusActive && int(c) <= len(CardStatusStrings)
}

var CardStatusStrings = map[CardStatus]string{
	CardStatusActive:   "active",
	CardStatusDeActive: "deActive",
	CardStatusDelete:   "delete",
}

func (c CardStatus) String() string {
	return CardStatusStrings[c]
}
