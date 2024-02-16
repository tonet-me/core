package cardservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	cardparam "github.com/tonet-me/tonet-core/param/card"
)

func (s Service) CreateNew(ctx context.Context, req cardparam.CreateNewRequest) (*cardparam.CreateNewResponse, error) {
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
		return nil, iErr
	}
	if existed {
		return nil, errors.New(`the card name is available`)
	}
	fmt.Println("new card", newCard)
	createdCard, cErr := s.repo.CreateNewCard(ctx, newCard)
	if cErr != nil {
		return nil, cErr
	}

	return &cardparam.CreateNewResponse{Card: createdCard}, nil
}
