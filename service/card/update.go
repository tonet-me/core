package cardservice

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) Update(ctx context.Context, req cardparam.UpdateRequest) (*cardparam.UpdateResponse, error) {
	const op = richerror.OP("cardservice.Update")
	fmt.Println(req)
	card, gErr := s.repo.GetCardByID(ctx, req.CardID)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr),
		)
	}
	if card.UserID != req.AuthenticatedUserID {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindForbidden),
			richerror.WithMessage(errmsg.ErrorMsgUserNotAllowed),
		)
	}
	//created to fills zero value to pointer fields in request
	var optionalCardField entity.Card
	s.checkOptionalOnUpdate(req, card, &optionalCardField)

	cardDataUpdate := entity.Card{
		UserID:       req.AuthenticatedUserID,
		Name:         req.UpdateData.Name,
		Title:        req.UpdateData.Title,
		About:        optionalCardField.About,
		PhotoURL:     optionalCardField.PhotoURL,
		PhoneNumbers: optionalCardField.PhoneNumbers,
		Emails:       optionalCardField.Emails,
		SocialMedias: optionalCardField.SocialMedias,
		Links:        optionalCardField.Links,
		Status:       req.UpdateData.Status,
	}

	updated, uErr := s.repo.UpdateCard(ctx, card.ID, cardDataUpdate)
	if uErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(uErr),
		)
	}

	return &cardparam.UpdateResponse{Updated: updated}, nil
}

func (s Service) checkOptionalOnUpdate(req cardparam.UpdateRequest, oldCard entity.Card, optionalCardField *entity.Card) {
	if req.UpdateData.About != nil {
		optionalCardField.About = *req.UpdateData.About
	} else {
		optionalCardField.About = oldCard.About

	}

	if req.UpdateData.PhotoURL != nil {
		optionalCardField.PhotoURL = *req.UpdateData.PhotoURL
	} else {
		optionalCardField.PhotoURL = oldCard.PhotoURL

	}

	if req.UpdateData.PhoneNumbers != nil {
		optionalCardField.PhoneNumbers = *req.UpdateData.PhoneNumbers
	} else {
		optionalCardField.PhoneNumbers = oldCard.PhoneNumbers

	}

	if req.UpdateData.Emails != nil {
		optionalCardField.Emails = *req.UpdateData.Emails
	} else {
		optionalCardField.Emails = oldCard.Emails

	}

	if req.UpdateData.SocialMedias != nil {
		optionalCardField.SocialMedias = *req.UpdateData.SocialMedias
	} else {
		optionalCardField.SocialMedias = oldCard.SocialMedias

	}

	if req.UpdateData.Links != nil {
		optionalCardField.Links = *req.UpdateData.Links
	} else {
		optionalCardField.Links = oldCard.Links

	}

}
