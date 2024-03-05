package cardvalidator

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"regexp"
)

/*
Required vs. Not Nil
When validating input values, there are two different scenarios about checking if input values are provided or not.

In the first scenario, an input value is considered missing if it is not entered or it is entered as a zero value (e.g. an empty string, a zero integer). You can use the validation.Required rule in this case.

In the second scenario, an input value is considered missing only if it is not entered. A pointer field is usually used in this case so that you can detect if a value is entered or not by checking if the pointer is nil or not. You can use the validation.NotNil rule to ensure a value is entered (even if it is a zero value).
*/

func (v Validator) CreateRequest(req cardparam.CreateNewRequest) (map[string]string, error) {
	const op = "cardvalidator.CreateRequest"

	fieldErrors := make(map[string]string)

	if err := validation.ValidateStruct(&req.CreateData,
		validation.Field(&req.CreateData.Title,
			validation.Required.Error(errmsg.ErrorMsgNeedTitle),
			validation.Length(1, 20)),

		validation.Field(&req.CreateData.Name,
			validation.Required.Error(errmsg.ErrorMsgNeedName),
			validation.Length(4, 25),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9]+(?:[_.][a-zA-Z0-9]+)*$"))),

		validation.Field(&req.CreateData.PhotoURL,
			validation.Length(10, 220)),

		validation.Field(&req.CreateData.Status,
			validation.NotNil, validation.By(v.doesStatusExist)),
	); err != nil {
		fmt.Println("err1", err)
		vErr := validation.Errors{}
		if errors.As(err, &vErr) {
			for key, value := range vErr {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}
	}
	v.validateEmails(req.CreateData.Emails, fieldErrors)
	v.validateLinks(req.CreateData.Links, fieldErrors)
	v.validateSocialMedias(req.CreateData.SocialMedias, fieldErrors)
	v.validatePhoneNumbers(req.CreateData.PhoneNumbers, fieldErrors)

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
