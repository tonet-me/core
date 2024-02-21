package usermongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	const op = richerror.OP("usermongo.GetUserByID")

	var user entity.User
	id, oErr := primitive.ObjectIDFromHex(userID)
	if oErr != nil {
		return entity.User{}, richerror.New(richerror.WithOp(op),
			richerror.WithMessage(fmt.Sprintf("userID %s is not a valid ObjectID", userID)),
			richerror.WithKind(richerror.ErrKindInvalid),
			richerror.WithInnerError(oErr))
	}

	filter := bson.D{{"id", id}}
	err := d.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return entity.User{}, richerror.New(richerror.WithOp(op),
				richerror.WithKind(richerror.ErrKindNotFound),
				richerror.WithMessage(errmsg.ErrorMsgNotFound),
				richerror.WithInnerError(err))
		}

		return entity.User{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	return user, nil
}
