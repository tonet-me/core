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

func (d DB) DeActiveUser(ctx context.Context, userID string) (bool, error) {
	id, oErr := primitive.ObjectIDFromHex(userID)
	if oErr != nil {
		return false, fmt.Errorf("userID is not a valid ObjectID")
	}

	filter := bson.D{{"_id", id}, {"status", entity.UserStatusActive}}

	update := bson.D{{"$set", bson.D{{"status", entity.UserStatusDeActive}}}}

	err := d.collection.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, nil
		}

		return false, err
	}

	return true, nil
}
