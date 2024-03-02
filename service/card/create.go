package cardservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) CreateNew(ctx context.Context, req cardparam.CreateNewRequest) (*cardparam.CreateNewResponse, error) {
	const op = richerror.OP("cardservice.CreateNew")

	//created to fills zero value to pointer fields in request
	var optionalCardField entity.Card

	existed, iErr := s.repo.IsCardExistByName(ctx, req.CreateData.Name)
	if iErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(iErr),
		)
	}
	if existed {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindStatusConflict),
			richerror.WithMessage(errmsg.ErrorMsgCardNameNotUnique),
		)
	}

	if !req.CreateData.Status.IsValid() {
		req.CreateData.Status = entity.CardStatusActive
	}

	newCard := entity.Card{
		UserID:       req.AuthenticatedUserID,
		Name:         req.CreateData.Name,
		Title:        req.CreateData.Title,
		PhotoURL:     optionalCardField.PhotoURL,
		PhoneNumbers: optionalCardField.PhoneNumbers,
		Emails:       optionalCardField.Emails,
		SocialMedias: optionalCardField.SocialMedias,
		Links:        optionalCardField.Links,
		Status:       req.CreateData.Status,
	}

	createdCard, cErr := s.repo.CreateNewCard(ctx, newCard)
	if cErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(cErr),
		)
	}

	return &cardparam.CreateNewResponse{Card: createdCard}, nil
}
