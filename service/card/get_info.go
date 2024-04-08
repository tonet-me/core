package cardservice

import (
	"context"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) GetInfoByID(ctx context.Context, req cardparam.GetInfoByIDRequest) (*cardparam.GetInfoByIDResponse, error) {
	const op = richerror.OP("cardservice.GetInfoByID")

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

	return &cardparam.GetInfoByIDResponse{Card: card}, nil
}

// IMPORTANT: Just use inner request (This method don't check authenticated user)
func (s Service) GetInfoByName(ctx context.Context, req cardparam.GetInfoByNameRequest) (*cardparam.GetInfoByNameResponse, error) {
	const op = richerror.OP("cardservice.GetInfoByName")

	card, gErr := s.repo.GetCardByName(ctx, req.Name)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	return &cardparam.GetInfoByNameResponse{Card: card}, nil
}
