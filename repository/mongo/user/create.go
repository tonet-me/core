package usermongo

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d DB) CreateNewUser(ctx context.Context, user entity.User) (entity.User, error) {
	insertResult, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("err", err)
		return entity.User{}, err
	}
	fmt.Println("id", insertResult.InsertedID)
	userObjectID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return entity.User{}, fmt.Errorf("couldn't convert objectID to string")
	}
	user.ID = userObjectID.String()

	return user, nil
}
