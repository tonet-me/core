package uservalidator

import (
	"errors"
	userparam "github.com/tonet-me/tonet-core/param/user"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) RefreshTokenRequest(req userparam.GetRefreshTokenRequest) (map[string]string, error) {
	const op = "uservalidator.LoginRegisterRequest"

	fieldErrors := make(map[string]string)

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.RefreshToken,
			validation.Required.Error(errmsg.ErrorMsgNeedRefreshToken)),
	); err != nil {

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
