package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
)

func (a Adapter) download(ctx context.Context, bucketName, objectName string) (*minio.Object, error) {
	const op = richerror.OP("minio.download")

	object, gErr := a.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if gErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(gErr),
		)
	}
	return object, nil
}
