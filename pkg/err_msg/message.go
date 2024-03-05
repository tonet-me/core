package errmsg

const (
	ErrorMsg                       = "err-msg"
	ErrorMsgNotFound               = "record not found"
	ErrorMsgSomethingWentWrong     = "something went wrong"
	ErrorMsgInvalidInput           = "invalid input"
	ErrorMsgInvalidStatus          = "invalid status"
	ErrorMsgPhoneNumberIsNotUnique = "phone number is not unique"
	ErrorMsgEmailIsNotUnique       = "email is not unique"
	ErrorMsgCardNameNotUnique      = "card name is not unique"
	ErrorMsgPhoneNumberIsNotValid  = "phone number is not valid"
	ErrorMsgUserNotAllowed         = "user not allowed"
	ErrorMsgUserNotFound           = "user not found"

	ErrorMsgNeedToken           = "token not found"
	ErrorMsgNeedRefreshToken    = "refresh_token not found"
	ErrorMsgInvalidRefreshToken = "refresh token invalid"
	ErrorMsgTypeOfOAuthInvalid  = "the oAuth provider type invalid"
	ErrorMsgExpiredToken        = "token is expired"

	ErrorMsgInvalidJson = "invalid json format"

	ErrorMsgCardStatusInvalid          = "the card status invalid"
	ErrorMsgCardSocialMediaTypeInvalid = "the social media type invalid"
	ErrorMsgNeedName                   = "name not found"
	ErrorMsgNeedTitle                  = "title not found"
)
