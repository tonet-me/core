package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

func (a Adapter) UploadUserProfilePhoto(ctx context.Context, userID string, file *multipart.File, fileSize int64) (string, error) {
	const op = richerror.OP("minio.DownloadUserProfilePhoto")

	fileName := strings.Join([]string{userID, strconv.FormatInt(time.Now().Unix(), 10)}, `_`)

	key, uErr := a.upload(ctx, a.userProfileBucketName, fileName, file, fileSize)
	if uErr != nil {
		return "", richerror.New(
			richerror.WithOp(op),
			richerror.WithInnerError(uErr),
		)
	}

	return key, nil
}

func (a Adapter) DownloadUserProfilePhoto(ctx context.Context, fileName string) (*minio.Object, error) {
	const op = richerror.OP("minio.DownloadUserProfilePhoto")

	object, dErr := a.download(ctx, a.userProfileBucketName, fileName)
	if dErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithInnerError(dErr),
		)
	}

	return object, nil
}
