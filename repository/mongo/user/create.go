package usermongo

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (d DB) CreateNewUser(ctx context.Context, user entity.User) (entity.User, error) {
	const op = richerror.OP("usermongo.CreateNewUser")

	timeNow := time.Now()
	user.CreatedAt = timeNow
	user.UpdatedAt = timeNow

	insertResult, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return entity.User{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	userObjectID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return entity.User{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
		)
	}
	user.ID = userObjectID.Hex()

	return user, nil
}
