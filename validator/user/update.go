package uservalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	userparam "github.com/tonet-me/tonet-core/param/user"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (v Validator) UpdateRequest(req userparam.UpdateRequest) (map[string]string, error) {
	const op = "uservalidator.UpdateRequest"

	fieldErrors := make(map[string]string)

	if err := validation.ValidateStruct(&req.UpdateData,
		validation.Field(&req.UpdateData.FirstName,
			validation.Length(2, 30),
		),
		validation.Field(&req.UpdateData.LastName,
			validation.Length(2, 30),
		),
		//validation.Field(&req.UpdateData.PhoneNumber,
		//	validation.Match(regexp.MustCompile(`^[+][0-9]*$`))),
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

	return nil, nil
}
