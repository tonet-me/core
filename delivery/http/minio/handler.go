package miniohandler

import (
	"github.com/tonet-me/tonet-core/adapter/minio"
	"github.com/tonet-me/tonet-core/service/auth"
	filevalidator "github.com/tonet-me/tonet-core/validator/file"
)

type Handler struct {
	client     *minio.Adapter
	fileVld    filevalidator.Validator
	authSvc    auth.Service
	authConfig auth.Config
}

func New(minioAdp *minio.Adapter, fileVld filevalidator.Validator, authSvc auth.Service, authConfig auth.Config) Handler {
	return Handler{
		client:     minioAdp,
		fileVld:    fileVld,
		authSvc:    authSvc,
		authConfig: authConfig,
	}
}
