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

func (d DB) DeActiveCard(ctx context.Context, cardID string) (bool, error) {
	const op = richerror.OP("cardmongo.DeActiveUser")

	id, oErr := primitive.ObjectIDFromHex(cardID)
	if oErr != nil {
		return false, richerror.New(richerror.WithOp(op),
			richerror.WithMessage(fmt.Sprintf("userID %s is not a valid ObjectID", cardID)),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithInnerError(oErr))
	}

	filter := bson.D{{"_id", id}, {"status", entity.CardStatusActive}}

	update := bson.D{{"$set", bson.D{{"status", entity.CardStatusDeActive}}}}

	err := d.collection.FindOneAndUpdate(ctx, filter, update).Err()
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

	return true, nil
}
