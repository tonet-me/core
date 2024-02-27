package visitservice

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) AddNewVisitToCard(ctx context.Context, cardID string, userAgent map[string]string) {
	const op = richerror.OP("visitservice.AddNewVisitToCard")

	newVisit := entity.Visit{
		CardID:    cardID,
		UserAgent: userAgent,
	}
	if err := s.repo.AddVisitToCard(ctx, newVisit); err != nil {
		log.Errorf("error in op:%s, with message:%s", op, err.Error())
	}

}
