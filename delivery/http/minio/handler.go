package miniohandler

import (
	"github.com/tonet-me/tonet-core/adapter/minio"
)

type Handler struct {
	client *minio.Adapter
}

func New(minioAdp *minio.Adapter) Handler {
	return Handler{client: minioAdp}
}
