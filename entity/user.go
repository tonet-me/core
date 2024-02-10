package entity

type User struct {
	ID              string
	FirstName       string
	LastName        string
	Email           string
	PhoneNumber     string
	ProfilePhotoURL string
	Status          UserStatus
}

type UserStatus int

const (
	Active UserStatus = iota + 1
	DeActive
	Suspend
)

func (u UserStatus) IsValid() bool {
	return u >= Active && int(u) <= len(UserTypeStrings)
}

var UserTypeStrings = map[UserStatus]string{
	Active:   "active",
	DeActive: "deActive",
	Suspend:  "suspend",
}

func (u UserStatus) String() string {
	return UserTypeStrings[u]
}
