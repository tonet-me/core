package entity

type OAuthType uint

const (
	OAuthTypeGoogle OAuthType = iota + 1
	OAuthTypeApple
)

var mapOAuthTypeString map[OAuthType]string = map[OAuthType]string{
	OAuthTypeGoogle: "google",
	OAuthTypeApple:  "apple",
}

func (o OAuthType) IsValid() bool {
	return o >= OAuthTypeGoogle && int(o) <= len(mapOAuthTypeString)
}
func (o OAuthType) GetTypeFromValue(value string) OAuthType {
	for index, valueFromMap := range mapOAuthTypeString {
		if value == valueFromMap {
			return index
		}
	}
	return 0
}
