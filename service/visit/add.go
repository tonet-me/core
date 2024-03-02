package visitservice

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/tonet-me/tonet-core/entity"
	visitparam "github.com/tonet-me/tonet-core/param/visit"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (s Service) AddNewVisitToCard(ctx context.Context, req visitparam.AddNewCardVisitRequest) (*visitparam.GetCardInfoByNameResponse, error) {
	const op = richerror.OP("visitservice.AddNewVisitToCard")

	getCard, gErr := s.GetCardInfoByName(ctx, visitparam.GetCardInfoByNameRequest{Name: req.CardName})
	if gErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithInnerError(gErr))
	}

	newVisit := entity.Visit{
		CardID: getCard.Card.ID,
		//UserAgent: ctx.Request().UserAgent(), //TODO: get user agent info
	}
	//TODO: check to use go routine and then -> need handle error with go?
	if err := s.repo.AddVisitToCard(ctx, newVisit); err != nil {
		log.Errorf("error in op:%s, with message:%s", op, err.Error())
	}
	return &visitparam.GetCardInfoByNameResponse{Card: getCard.Card}, nil

}
