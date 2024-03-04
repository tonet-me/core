package cardmongo

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (d DB) CreateNewCard(ctx context.Context, card entity.Card) (entity.Card, error) {
	const op = richerror.OP("cardmongo.CreateNewCard")

	timeNow := time.Now()
	card.CreatedAt = timeNow
	card.UpdatedAt = timeNow

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
	card.ID = cardObjectID.Hex()

	return card, nil
}
