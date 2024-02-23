package uservalidator

import (
	"errors"
	userparam "github.com/tonet-me/tonet-core/param/user"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateLoginRegisterRequest(req userparam.LoginOrRegisterRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRegisterRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Token,
			validation.Required.Error(errmsg.ErrorMsgNeedToken)),
		validation.Field(&req.ProviderName,
			validation.Required, validation.By(v.doesTypeOfOAuthProviderExist)),
	); err != nil {
		fieldErrors := make(map[string]string)

		vErr := validation.Errors{}
		if errors.As(err, &vErr) {
			for key, value := range vErr {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithMessage(errmsg.ErrorMsgInvalidInput),
			richerror.WithMeta(map[string]interface{}{"req": req}),
			richerror.WithInnerError(err),
		)
	}

	//nolint
	return nil, nil
}
