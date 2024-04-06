package cardmongo

import (
	"context"
	"errors"
	"github.com/tonet-me/tonet-core/entity"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) GetAllCardsByUserID(ctx context.Context, userID string) ([]entity.Card, error) {
	const op = richerror.OP("cardmongo.GetAllCardsByUserID")

	var cards []entity.Card

	filter := bson.D{{"user_id", userID}, {"status", bson.D{{"$ne", entity.CardStatusDelete}}}}
	cur, err := d.collection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return nil, richerror.New(richerror.WithOp(op),
				richerror.WithKind(richerror.ErrKindNotFound),
				richerror.WithMessage(errmsg.ErrorMsgNotFound),
				richerror.WithInnerError(err))
		}

		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	if cur == nil {
		return nil, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindNotFound),
			richerror.WithMessage(errmsg.ErrorMsgNotFound),
			richerror.WithInnerError(err))
	}

	for cur.Next(ctx) {
		var card entity.Card
		if cErr := cur.Decode(&card); cErr != nil {
			return nil, richerror.New(richerror.WithOp(op),
				richerror.WithKind(richerror.ErrKindUnExpected),
				richerror.WithInnerError(cErr))
		}
		cards = append(cards, card)
	}

	return cards, nil
}
