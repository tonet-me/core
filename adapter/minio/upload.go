package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

func (a Adapter) upload(ctx context.Context, bucketName, objectName string, file *multipart.File, fileSize int64) (string, error) {

	uploadInfo, pErr := a.client.PutObject(ctx, bucketName, objectName, *file, fileSize, minio.PutObjectOptions{})
	if pErr != nil {
		return "", pErr
	}

	return uploadInfo.Key, nil
}
