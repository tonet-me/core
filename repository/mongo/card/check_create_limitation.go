package cardmongo

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (d DB) CheckIsCreateCardLimitation(ctx context.Context, userID string, limit uint) (bool, error) {
	const op = richerror.OP("cardmongo.CheckCreateCardLimitation")

	opts := options.Count().SetHint("_id_")
	filter := bson.D{{"user_id", userID}, {"status", bson.D{{"$ne", entity.CardStatusDelete}}}}
	counted, err := d.collection.CountDocuments(ctx, filter, opts)
	if err != nil {
		return false, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	return counted >= int64(limit), nil

}
