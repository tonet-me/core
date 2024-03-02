package entity

import (
	"time"
)

type User struct {
	ID              string           `bson:"_id,omitempty" json:"id"`
	FirstName       string           `bson:"first_name,omitempty" json:"first_name"`
	LastName        string           `bson:"last_name,omitempty" json:"last_name"`
	Email           string           `bson:"email,omitempty" json:"email"`
	EmailVerified   bool             `bson:"email_verified,omitempty" json:"email_verified"`
	PhoneNumber     PhoneNumberValue `bson:"phone_number,omitempty,inline" json:"phone_number"`
	ProfilePhotoURL string           `bson:"profile_photo_url" json:"profile_photo_url"`
	Status          UserStatus       `bson:"status,omitempty" json:"status"`
	CreatedAt       time.Time        `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time        `bson:"updated_at" json:"updated_at"`
}

type UserStatus int

const (
	UserStatusActive UserStatus = iota + 1
	UserStatusDeActive
	UserStatusSuspend
)

func (u UserStatus) IsValid() bool {
	return u >= UserStatusActive && int(u) <= len(UserStatusStrings)
}

var UserStatusStrings = map[UserStatus]string{
	UserStatusActive:   "active",
	UserStatusDeActive: "deActive",
	UserStatusSuspend:  "suspend",
}

func (u UserStatus) String() string {
	return UserStatusStrings[u]
}
