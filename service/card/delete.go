package cardservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) Delete(ctx context.Context, req cardparam.DeleteRequest) (*cardparam.DeleteResponse, error) {
	const op = richerror.OP("cardservice.Delete")

	//TODO: check if not deleted
	card, gErr := s.repo.GetCardByID(ctx, req.CardID)
	if gErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	if card.UserID != req.AuthenticatedUserID || card.Status == entity.CardStatusDelete {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindForbidden),
			richerror.WithMessage(errmsg.ErrorMsgUserNotAllowed),
		)
	}

	success, dErr := s.repo.DeleteCard(ctx, req.CardID)
	if dErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(dErr))
	}

	return &cardparam.DeleteResponse{Success: success}, nil
}
