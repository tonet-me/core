package cardmongo

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d DB) CreateNewCard(ctx context.Context, card entity.Card) (entity.Card, error) {
	const op = richerror.OP("cardmongo.CreateNewCard")

	insertResult, err := d.collection.InsertOne(ctx, card)
	if err != nil {
		return entity.Card{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	cardObjectID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return entity.Card{}, richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
		)
	}
	card.ID = cardObjectID.String()

	return card, nil
}
