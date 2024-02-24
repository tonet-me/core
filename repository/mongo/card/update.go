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

func (d DB) UpdateCard(ctx context.Context, cardID string, card entity.Card) (bool, error) {
	const op = richerror.OP("cardmongo.UpdateCard")

	id, oErr := primitive.ObjectIDFromHex(cardID)
	if oErr != nil {
		return false, richerror.New(richerror.WithOp(op),
			richerror.WithMessage(fmt.Sprintf("cardID %s is not a valid ObjectID", cardID)),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithInnerError(oErr))

	}

	update := bson.D{{"$set", card}}
	updatedResult, err := d.collection.UpdateByID(ctx, id, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, richerror.New(richerror.WithOp(op),
				richerror.WithKind(richerror.ErrKindNotFound),
				richerror.WithMessage(errmsg.ErrorMsgNotFound),
				richerror.WithInnerError(err))

		}

		return false, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	if updatedResult.MatchedCount == 0 {
		return false, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindNotFound),
			richerror.WithMessage(errmsg.ErrorMsgNotFound),
		)
	}

	return true, nil
}
