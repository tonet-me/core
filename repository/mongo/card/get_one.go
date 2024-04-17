package cardmongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) GetCardByID(ctx context.Context, cardID string) (entity.Card, error) {
	const op = richerror.OP("cardmongo.GetCardByID")

	var card entity.Card
	id, oErr := primitive.ObjectIDFromHex(cardID)
	if oErr != nil {
		return entity.Card{}, richerror.New(richerror.WithOp(op),
			richerror.WithMessage(fmt.Sprintf("cardID %s is not a valid ObjectID", cardID)),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithInnerError(oErr))
	}

	filter := bson.D{{"_id", id}, {"status", bson.D{{"$ne", entity.CardStatusDelete}}}}
	err := d.collection.FindOne(ctx, filter).Decode(&card)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return entity.Card{}, richerror.New(richerror.WithOp(op),
				richerror.WithKind(richerror.ErrKindNotFound),
				richerror.WithMessage(errmsg.ErrorMsgNotFound),
				richerror.WithInnerError(err))
		}

		return entity.Card{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	return card, nil
}

// Warning: Just use with visit service
func (d DB) GetOnlyActiveCardByName(ctx context.Context, name string) (entity.Card, error) {
	const op = richerror.OP("cardmongo.GetOnlyActiveCardByName")

	var card entity.Card

	//Just active card can show
	ninArrayToFindOnlyActiveCard := []entity.CardStatus{entity.CardStatusDelete, entity.CardStatusDeActive}

	filter := bson.D{{"name", name}, {"status", bson.M{"$nin": ninArrayToFindOnlyActiveCard}}}
	err := d.collection.FindOne(ctx, filter).Decode(&card)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return entity.Card{}, richerror.New(richerror.WithOp(op),
				richerror.WithKind(richerror.ErrKindNotFound),
				richerror.WithMessage(errmsg.ErrorMsgNotFound),
				richerror.WithInnerError(err))
		}

		return entity.Card{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	return card, nil
}
