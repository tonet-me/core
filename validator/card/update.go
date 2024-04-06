package cardvalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"regexp"
)

func (v Validator) UpdateRequest(req cardparam.UpdateRequest) (map[string]string, error) {
	const op = "cardvalidator.UpdateRequest"

	fieldErrors := make(map[string]string)

	if err := validation.ValidateStruct(&req.UpdateData,
		validation.Field(&req.UpdateData.Title,
			validation.Required.Error(errmsg.ErrorMsgNeedTitle),
			validation.Length(1, 20)),

		validation.Field(&req.UpdateData.Name,
			validation.Required.Error(errmsg.ErrorMsgNeedName),
			validation.Length(4, 25),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9]+(?:[_.][a-zA-Z0-9]+)*$"))),

		validation.Field(&req.UpdateData.PhotoURL,
			validation.Length(10, 220)),

		validation.Field(&req.UpdateData.Status,
			validation.NotNil, validation.By(v.doesStatusExist)),
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
	v.validateEmails(req.UpdateData.Emails, fieldErrors)
	v.validateLinks(req.UpdateData.Links, fieldErrors)
	v.validateSocialMedias(req.UpdateData.SocialMedias, fieldErrors)
	v.validatePhoneNumbers(req.UpdateData.PhoneNumbers, fieldErrors)

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
