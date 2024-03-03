package cardvalidator

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/tonet-me/tonet-core/entity"
	"regexp"
)

func (v Validator) validateLinks(links *[]entity.Link, fieldErrors map[string]string) {
	if links != nil {
		//for value
		for _, link := range *links {
			if err := validation.ValidateStruct(&link,
				validation.Field(&link.Value, validation.Required, is.URL),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["link."+key] = value.Error()
						}
					}
				}
			}
			//for title
			if err := validation.ValidateStruct(&link,
				validation.Field(&link.Title, validation.Required),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["link."+key] = value.Error()
						}
					}
				}
			}
		}
	}
}

func (v Validator) validateEmails(emails *[]entity.Email, fieldErrors map[string]string) {
	if emails != nil {
		//for value
		for _, email := range *emails {
			if err := validation.ValidateStruct(&email,
				validation.Field(&email.Value, validation.Required, is.Email),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["email."+key] = value.Error()
						}
					}
				}
			}

			//for title
			if err := validation.ValidateStruct(&email,
				validation.Field(&email.Title, validation.Required),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["email."+key] = value.Error()
						}
					}
				}
			}
		}
	}
}

func (v Validator) validateSocialMedias(socialMedias *[]entity.SocialMedia, fieldErrors map[string]string) {
	if socialMedias != nil {

		//for value
		for _, socialMedia := range *socialMedias {
			if err := validation.ValidateStruct(&socialMedia,
				validation.Field(&socialMedia.Value, validation.Required),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["social_media."+key] = value.Error()
						}
					}
				}
			}

			//for title
			if err := validation.ValidateStruct(&socialMedia,
				validation.Field(&socialMedia.Type, validation.Required, validation.By(v.doesSocialMediaTypeExist)),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["social_media."+key] = value.Error()
						}
					}
				}
			}
		}
	}
}

func (v Validator) validatePhoneNumbers(phoneNumbers *[]entity.PhoneNumber, fieldErrors map[string]string) {
	if phoneNumbers != nil {

		//for value
		for _, phoneNumber := range *phoneNumbers {
			if err := validation.ValidateStruct(&phoneNumber.Value,
				validation.Field(&phoneNumber.Value.Number, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]`))),
				validation.Field(&phoneNumber.Value.CountryCode, validation.Required),
				validation.Field(&phoneNumber.Value.Prefix, validation.Required),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["phone_number."+key] = value.Error()
						}
					}
				}
			}

			//for title
			if err := validation.ValidateStruct(&phoneNumber,
				validation.Field(&phoneNumber.Title, validation.Required),
			); err != nil {
				vErr := validation.Errors{}
				if errors.As(err, &vErr) {
					for key, value := range vErr {
						if value != nil {
							fieldErrors["phone_number."+key] = value.Error()
						}
					}
				}
			}
		}
	}
}
