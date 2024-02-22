package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	"log"
)

type Config struct {
	Endpoint        string `koanf:"endpoint"`
	AccessKeyID     string `koanf:"access_key_id"`
	SecretAccessKey string `koanf:"secret_access_key"`
	UseSSL          bool   `koanf:"use_ssl"`
	UserBucketName  string `koanf:"user_bucket_name"`
}
type Adapter struct {
	client         *minio.Client
	userBucketName string
}

func New(cfg Config) *Adapter {
	const op = richerror.OP("minio.New")

	// Initialize minio client object.
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		panic(fmt.Errorf("op:%v,\nwith err:%v", op, err))
	}

	newAdapter := Adapter{
		client:         minioClient,
		userBucketName: cfg.UserBucketName,
	}

	//create buckets
	newAdapter.createBuckets(cfg.UserBucketName)

	return &newAdapter
}

func (a Adapter) createBuckets(bucketNames ...string) {
	const op = richerror.OP("minio.createBuckets")

	for _, bucketName := range bucketNames {
		mErr := a.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if mErr != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists := a.client.BucketExists(context.Background(), bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				panic(fmt.Errorf("op:%v,\nwith err:%v", op, mErr))
			}
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
	}
}
