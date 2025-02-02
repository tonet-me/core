package miniohandler

import (
	"github.com/tonet-me/tonet-core/adapter/minio"
	"github.com/tonet-me/tonet-core/service/auth"
)

type Handler struct {
	client     *minio.Adapter
	authSvc    auth.Service
	authConfig auth.Config
}

func New(minioAdp *minio.Adapter, authSvc auth.Service, authConfig auth.Config) Handler {
	return Handler{
		client:     minioAdp,
		authSvc:    authSvc,
		authConfig: authConfig,
	}
}
