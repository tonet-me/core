package entity

type SocialMedia struct {
	Type  SocialMediasType `bson:"type" json:"type"`
	Value string           `bson:"value" json:"value"`
}

type SocialMediasType int

const (
	SocialMediasTypeFacebook SocialMediasType = iota + 1
	SocialMediasTypeYoutube
	SocialMediasTypeWhatsapp
	SocialMediasTypeInstagram
	SocialMediasTypeWechat
	SocialMediasTypeTiktok
	SocialMediasTypeSnapchat
	SocialMediasTypeTwitter
	SocialMediasTypeLinkedin
	SocialMediasTypeReddit
	SocialMediasTypePinterest
	SocialMediasTypeTelegram
	SocialMediasTypeDiscord
	SocialMediasTypeTinder
	SocialMediasTypeLine
	SocialMediasTypeClubhouse
	SocialMediasTypeTwitch
)

func (s SocialMediasType) IsValid() bool {
	return s >= SocialMediasTypeFacebook && int(s) <= len(SocialMediasTypeStrings)
}

var SocialMediasTypeStrings = map[SocialMediasType]string{
	SocialMediasTypeFacebook:  "facebook",
	SocialMediasTypeYoutube:   "youtube",
	SocialMediasTypeWhatsapp:  "Whatsapp",
	SocialMediasTypeInstagram: "instagram",
	SocialMediasTypeWechat:    "wechat",
	SocialMediasTypeTiktok:    "tiktok",
	SocialMediasTypeSnapchat:  "snapchat",
	SocialMediasTypeTwitter:   "twitter",
	SocialMediasTypeLinkedin:  "linkedin",
	SocialMediasTypeReddit:    "reddit",
	SocialMediasTypePinterest: "pinterest",
	SocialMediasTypeTelegram:  "telegram",
	SocialMediasTypeDiscord:   "discord",
	SocialMediasTypeTinder:    "tinder",
	SocialMediasTypeLine:      "line",
	SocialMediasTypeClubhouse: "clubhouse",
	SocialMediasTypeTwitch:    "twitch",
}

func (s SocialMediasType) String() string {
	return SocialMediasTypeStrings[s]
}
