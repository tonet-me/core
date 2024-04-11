package cardvalidator

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

/*
Required vs. Not Nil
When validating input values, there are two different scenarios about checking if input values are provided or not.

In the first scenario, an input value is considered missing if it is not entered or it is entered as a zero value (e.g. an empty string, a zero integer). You can use the validation.Required rule in this case.

In the second scenario, an input value is considered missing only if it is not entered. A pointer field is usually used in this case so that you can detect if a value is entered or not by checking if the pointer is nil or not. You can use the validation.NotNil rule to ensure a value is entered (even if it is a zero value).
*/

func (v Validator) CheckIfCardDeleted(req cardparam.DeleteRequest) (map[string]string, error) {
	const op = "cardvalidator.ActiveRequest"

	card, gErr := v.repo.GetCardByID(context.Background(), req.CardID)
	if gErr != nil {
		return nil, richerror.New(richerror.WithInnerError(gErr))
	}

	if card.Status == entity.CardStatusDelete {
		return nil, richerror.New(
			richerror.WithKind(richerror.ErrKindNotFound),
			richerror.WithMessage(errmsg.ErrorMsgNotFound))
	}

	return nil, nil
}
