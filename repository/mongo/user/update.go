package usermongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) UpdateUser(ctx context.Context, userID string, user entity.User) (bool, error) {
	id, oErr := primitive.ObjectIDFromHex(userID)
	if oErr != nil {
		return false, fmt.Errorf("userID is not a valid ObjectID")
	}

	//filter := bson.D{{"_id", id}}

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
