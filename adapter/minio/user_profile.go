package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"strings"
	"time"
)

func (a Adapter) UploadUserProfilePhoto(ctx context.Context, userID string, file *multipart.File, fileSize int64) (string, error) {
	fileName := strings.Join([]string{userID, time.Now().String()}, `_`)

	key, pErr := a.upload(ctx, a.userBucketName, fileName, file, fileSize)
	if pErr != nil {
		return "", pErr
	}

	return key, nil
}

func (a Adapter) DownloadUserProfilePhoto(ctx context.Context, fileName string) (*minio.Object, error) {
	object, pErr := a.download(ctx, a.userBucketName, fileName)
	if pErr != nil {
		return nil, pErr
	}

	return object, nil
}
