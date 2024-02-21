package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
	// Initialize minio client object.
	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		log.Fatalln(`minio service can't started:`, err)
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
	for _, bucketName := range bucketNames {
		mErr := a.client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if mErr != nil {
			// Check to see if we already own this bucket (which happens if you run this twice)
			exists, errBucketExists := a.client.BucketExists(context.Background(), bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(mErr)
			}
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
	}
}
