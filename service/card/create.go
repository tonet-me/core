package cardservice

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) CreateNew(ctx context.Context, req cardparam.CreateNewRequest) (*cardparam.CreateNewResponse, error) {
	const op = richerror.OP("cardservice.CreateNew")

	isLimited, cErr := s.repo.CheckIsCreateCardLimitation(ctx, req.AuthenticatedUserID, s.config.CreateCardLimitation)
	if cErr != nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithInnerError(cErr),
		)
	}
	if isLimited {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindForbidden),
			richerror.WithMessage(fmt.Sprintf(errmsg.ErrorMsgCreatCardLimitation, s.config.CreateCardLimitation)),
		)
	}

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

	//created to fills zero value to pointer fields in request
	var optionalCardField entity.Card
	s.checkOptionalOnCreate(req, &optionalCardField)

	newCard := entity.Card{
		UserID:       req.AuthenticatedUserID,
		Name:         req.CreateData.Name,
		Title:        req.CreateData.Title,
		About:        optionalCardField.About,
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

func (s Service) checkOptionalOnCreate(req cardparam.CreateNewRequest, optionalCardField *entity.Card) {
	if req.CreateData.About != nil {
		optionalCardField.About = *req.CreateData.About
	}
	if req.CreateData.PhotoURL != nil {
		optionalCardField.PhotoURL = *req.CreateData.PhotoURL
	}
	if req.CreateData.PhoneNumbers != nil {
		optionalCardField.PhoneNumbers = *req.CreateData.PhoneNumbers
	}
	if req.CreateData.Emails != nil {
		optionalCardField.Emails = *req.CreateData.Emails
	}
	if req.CreateData.SocialMedias != nil {
		optionalCardField.SocialMedias = *req.CreateData.SocialMedias
	}
	if req.CreateData.Links != nil {
		optionalCardField.Links = *req.CreateData.Links
	}
}
