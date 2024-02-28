package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"mime/multipart"
)

func (a Adapter) upload(ctx context.Context, bucketName, objectName string, file *multipart.File, fileSize int64) (string, error) {
	const op = richerror.OP("minio.upload")

	uploadInfo, pErr := a.client.PutObject(ctx, bucketName, objectName, *file, fileSize, minio.PutObjectOptions{})
	if pErr != nil {
		return "", richerror.New(
			richerror.WithOp(op),
			richerror.WithKind(richerror.ErrKindUnExpected),
			richerror.WithInnerError(pErr),
		)
	}

	return uploadInfo.Key, nil
}
