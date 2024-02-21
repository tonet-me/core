package usermongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) DeActiveUser(ctx context.Context, userID string) (bool, error) {
	const op = richerror.OP("usermongo.DeActiveUser")

	id, oErr := primitive.ObjectIDFromHex(userID)
	if oErr != nil {
		return false, richerror.New(richerror.WithOp(op),
			richerror.WithMessage(fmt.Sprintf("userID %s is not a valid ObjectID", userID)),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithInnerError(oErr))
	}

	filter := bson.D{{"_id", id}, {"status", entity.UserStatusActive}}

	update := bson.D{{"$set", bson.D{{"status", entity.UserStatusDeActive}}}}

	err := d.collection.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, nil
		}

		return false, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	return true, nil
}
