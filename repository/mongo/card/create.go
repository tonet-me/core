package cardmongo

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d DB) CreateNewCard(ctx context.Context, card entity.Card) (entity.Card, error) {
	insertResult, err := d.collection.InsertOne(ctx, card)
	fmt.Println("insert res", insertResult.InsertedID)
	fmt.Println("card db", card)
	if err != nil {
		fmt.Println("err", err)
		return entity.Card{}, err
	}

	cardObjectID, ok := insertResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return entity.Card{}, fmt.Errorf("couldn't convert objectID to string")
	}
	card.ID = cardObjectID.String()

	return card, nil
}
