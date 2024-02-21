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

	newCard := entity.Card{
		UserID:       req.AuthenticatedUserID,
		Name:         req.CreateData.Name,
		Title:        req.CreateData.Title,
		Photo:        req.CreateData.Photo,
		PhoneNumbers: req.CreateData.PhoneNumbers,
		Emails:       req.CreateData.Emails,
		SocialMedias: req.CreateData.SocialMedias,
		Links:        req.CreateData.Links,
		Status:       req.CreateData.Status,
	}

	existed, iErr := s.repo.IsCardExistByName(ctx, req.CreateData.Name)
	if iErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(iErr),
		)
	}
	if existed {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindNotFound),
			richerror.WithMessage(errmsg.ErrorMsgNotFound),
		)
	}

	createdCard, cErr := s.repo.CreateNewCard(ctx, newCard)
	if cErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(cErr),
		)
	}

	return &cardparam.CreateNewResponse{Card: createdCard}, nil
}
