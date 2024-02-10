package usermongo

import (
	"context"
	"errors"
	"github.com/tonet-me/tonet-core/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) IsUserExistByEmail(ctx context.Context, email string) (bool, entity.User, error) {
	var user entity.User
	filter := bson.D{{"email", email}}
	err := d.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, entity.User{}, nil
		}

		return false, entity.User{}, err
	}

	return true, user, nil
}
