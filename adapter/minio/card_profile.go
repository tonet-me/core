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

func (a Adapter) UploadCardProfile(ctx context.Context, cardID string, file *multipart.File, fileSize int64) (string, error) {
	const op = richerror.OP("minio.UploadCardProfile")

	fileName := strings.Join([]string{cardID, strconv.FormatInt(time.Now().Unix(), 10)}, `_`)

	key, uErr := a.upload(ctx, a.cardProfileBucketName, fileName, file, fileSize)
	if uErr != nil {
		return "", richerror.New(
			richerror.WithOp(op),
			richerror.WithInnerError(uErr),
		)
	}

	return key, nil
}

func (a Adapter) DownloadCardProfile(ctx context.Context, fileName string) (*minio.Object, error) {
	const op = richerror.OP("minio.DownloadCardProfile")

	object, dErr := a.download(ctx, a.cardProfileBucketName, fileName)
	if dErr != nil {
		return nil, richerror.New(
			richerror.WithOp(op),
			richerror.WithInnerError(dErr),
		)
	}

	return object, nil
}
