package visitservice

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	"github.com/tonet-me/tonet-core/logger"
	visitparam "github.com/tonet-me/tonet-core/param/visit"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"log/slog"
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
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()), slog.Any("visit-data", newVisit))
	}
	return &visitparam.GetCardInfoByNameResponse{Card: getCard.Card}, nil

}
