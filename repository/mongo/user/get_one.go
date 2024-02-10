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

func (d DB) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	var user entity.User
	id, oErr := primitive.ObjectIDFromHex(userID)
	if oErr != nil {
		return entity.User{}, fmt.Errorf("userID is not a valid ObjectID")
	}

	filter := bson.D{{"id", id}}
	err := d.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return entity.User{}, fmt.Errorf("not found user")
		}

		return entity.User{}, err
	}

	return user, nil
}
