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

func (d DB) UpdateUser(ctx context.Context, userID string, user entity.User) (bool, error) {
	const op = richerror.OP("usermongo.UpdateUser")

	id, oErr := primitive.ObjectIDFromHex(userID)
	if oErr != nil {
		return false, richerror.New(richerror.WithOp(op),
			richerror.WithMessage(fmt.Sprintf("userID %s is not a valid ObjectID", userID)),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithInnerError(oErr))

	}

	update := bson.D{{"$set", user}}
	updatedResult, err := d.collection.UpdateByID(ctx, id, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, fmt.Errorf("not found user")

		}

		return false, err
	}
	if updatedResult.MatchedCount == 0 {
		return false, fmt.Errorf("not found user")
	}

	return true, nil
}
