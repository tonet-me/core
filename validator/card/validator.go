package cardvalidator

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
)

type Repository interface {
	GetCardByID(ctx context.Context, cardID string) (entity.Card, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository) Validator {
	return Validator{
		repo: repo,
	}
}

func (v Validator) doesStatusExist(value interface{}) error {
	statusType, ok := value.(entity.CardStatus)
	if !ok {
		return fmt.Errorf(errmsg.ErrorMsgCardStatusInvalid)
	}
	validType := statusType.IsValid()
	if !validType {
		return fmt.Errorf(errmsg.ErrorMsgCardStatusInvalid)
	}

	return nil
}

func (v Validator) doesSocialMediaTypeExist(value interface{}) error {
	socialMediaType, ok := value.(entity.SocialMediasType)
	if !ok {
		return fmt.Errorf(errmsg.ErrorMsgCardSocialMediaTypeInvalid)
	}
	validType := socialMediaType.IsValid()
	if !validType {
		return fmt.Errorf(errmsg.ErrorMsgCardSocialMediaTypeInvalid)
	}

	return nil
}
