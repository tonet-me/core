package cardservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) DeActive(ctx context.Context, req cardparam.DeActiveRequest) (*cardparam.DeActiveResponse, error) {
	const op = richerror.OP("cardservice.DeActive")

	card, gErr := s.repo.GetCardByID(ctx, req.CardID)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	if card.UserID != req.AuthenticatedUserID {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindForbidden),
			richerror.WithMessage(errmsg.ErrorMsgUserNotAllowed),
		)
	}

	success, dErr := s.repo.DeActiveCard(ctx, req.CardID)
	if dErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(dErr))
	}

	return &cardparam.DeActiveResponse{Success: success}, nil
}
