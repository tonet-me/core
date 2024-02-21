package usermongo

import (
	"context"
	"errors"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) IsUserExistByEmail(ctx context.Context, email string) (bool, entity.User, error) {
	const op = richerror.OP("usermongo.IsUserExistByEmail")

	var user entity.User
	filter := bson.D{{"email", email}}
	err := d.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, entity.User{}, nil
		}

		return false, entity.User{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	return true, user, nil
}
