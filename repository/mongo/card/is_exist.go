package cardmongo

import (
	"context"
	"errors"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// just check card exist by name, we dont need to check deleted or not deleted
func (d DB) IsCardExistByName(ctx context.Context, name string) (bool, error) {
	const op = richerror.OP("cardmongo.IsCardExistByName")

	opts := options.Count().SetHint("_id_")
	filter := bson.D{{"name", bson.M{"$regex": primitive.Regex{
		Pattern: name,
		Options: "i",
	}}}}
	counted, err := d.collection.CountDocuments(ctx, filter, opts)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, nil
		}

		return false, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	if counted != 0 {
		return true, nil
	}

	return false, nil
}
