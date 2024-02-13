package entity

type User struct {
	ID              string
	FirstName       string     `bson:"first_name,omitempty"`
	LastName        string     `bson:"last_name,omitempty"`
	Email           string     `bson:"email"`
	PhoneNumber     string     `bson:"phone_number"`
	ProfilePhotoURL string     `bson:"profile_photo_url,omitempty"`
	Status          UserStatus `bson:"status"`
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
