package cardvalidator

import (
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
)

type Validator struct {
}

func New() Validator {
	return Validator{}
}

func (v Validator) doesStatusExist(value interface{}) error {
	statusType, ok := value.(entity.CardStatus)
	if !ok {
		return fmt.Errorf(errmsg.ErrorMsgCardStatusInvalid)
	}
	validType := statusType.IsValid()
	if !validType {
		return fmt.Errorf(errmsg.ErrorMsgCardStatusInvalid)
	}

	return nil
}

func (v Validator) doesSocialMediaTypeExist(value interface{}) error {
	socialMediaType, ok := value.(entity.SocialMediasType)
	if !ok {
		return fmt.Errorf(errmsg.ErrorMsgCardSocialMediaTypeInvalid)
	}
	validType := socialMediaType.IsValid()
	if !validType {
		return fmt.Errorf(errmsg.ErrorMsgCardSocialMediaTypeInvalid)
	}

	return nil
}
