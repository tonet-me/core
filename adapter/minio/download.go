package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (a Adapter) download(ctx context.Context, bucketName, objectName string) (*minio.Object, error) {

	object, gErr := a.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if gErr != nil {
		return nil, gErr
	}
	return object, nil
}
