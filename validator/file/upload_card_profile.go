package filevalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	fileparam "github.com/tonet-me/tonet-core/param/file"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (v Validator) UploadCardProfileRequest(req fileparam.CardProfileUploadRequest) (map[string]string, error) {
	const op = "cardvalidator.UploadCardProfileRequest"

	fieldErrors := make(map[string]string)

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.CardID,
			validation.Required.Error(errmsg.ErrorMsgNeedCardID)),
	); err != nil {
		vErr := validation.Errors{}
		if errors.As(err, &vErr) {
			for key, value := range vErr {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
	}

	if len(fieldErrors) > 0 {

		return fieldErrors, richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithMessage(errmsg.ErrorMsgInvalidInput),
			richerror.WithMeta(map[string]interface{}{"req": req}),
		)
	}

	return nil, nil
}
