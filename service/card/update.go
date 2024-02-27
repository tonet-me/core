package cardservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) Update(ctx context.Context, req cardparam.UpdateRequest) (*cardparam.UpdateResponse, error) {
	const op = richerror.OP("cardservice.Update")

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

	cardDataUpdate := entity.Card{
		Name:         req.UpdateData.Name,
		Title:        req.UpdateData.Title,
		PhotoURL:     req.UpdateData.PhotoURL,
		PhoneNumbers: req.UpdateData.PhoneNumbers,
		Emails:       req.UpdateData.Emails,
		SocialMedias: req.UpdateData.SocialMedias,
		Links:        req.UpdateData.Links,
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
