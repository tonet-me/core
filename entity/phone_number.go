package entity

type PhoneNumberValue struct {
	Number      string `bson:"number" json:"number"`
	Prefix      string `bson:"prefix" json:"prefix"`
	CountryCode string `bson:"country_code" json:"country_code"`
}
