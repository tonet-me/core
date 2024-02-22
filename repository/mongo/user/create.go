package usermongo

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d DB) CreateNewUser(ctx context.Context, user entity.User) (entity.User, error) {
	const op = richerror.OP("usermongo.CreateNewUser")

	insertResult, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return entity.User{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	userObjectID, ok := insertResult.InsertedID.(primitive.ObjectID)
	fmt.Println("created id", userObjectID)
	if !ok {
		return entity.User{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
		)
	}
	user.ID = userObjectID.String()

	return user, nil
}
