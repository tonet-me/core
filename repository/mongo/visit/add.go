package visitmongo

import (
	"context"
	"github.com/tonet-me/tonet-core/entity"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"time"
)

func (d DB) AddVisitToCard(ctx context.Context, visit entity.Visit) error {
	const op = richerror.OP("visitmongo.AddVisitToCard")

	timeNow := time.Now()
	visit.CreatedAt = timeNow
	visit.UpdatedAt = timeNow

	_, err := d.collection.InsertOne(ctx, visit)
	if err != nil {
		return richerror.New(richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(err))
	}

	return nil
}
