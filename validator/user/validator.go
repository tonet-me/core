package uservalidator

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

func (v Validator) doesTypeOfOAuthProviderExist(value interface{}) error {
	providerType, ok := value.(entity.OAuthType)
	if !ok {
		return fmt.Errorf(errmsg.ErrorMsgTypeOfOAuthInvalid)
	}
	validType := providerType.IsValid()
	if !validType {
		return fmt.Errorf(errmsg.ErrorMsgTypeOfOAuthInvalid)
	}

	return nil
}
