package entity

type SocialMedia struct {
	Type  SocialMediasType `bson:"type"`
	Value string           `bson:"value"`
}

type SocialMediasType int

const (
	SocialMediasTypeInstagram SocialMediasType = iota + 1
	SocialMediasTypeYoutube
)

func (s SocialMediasType) IsValid() bool {
	return s >= SocialMediasTypeInstagram && int(s) <= len(SocialMediasTypeStrings)
}

var SocialMediasTypeStrings = map[SocialMediasType]string{
	SocialMediasTypeInstagram: "instagram",
	SocialMediasTypeYoutube:   "youtube",
}

func (s SocialMediasType) String() string {
	return SocialMediasTypeStrings[s]
}
