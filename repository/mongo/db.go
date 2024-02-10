package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI string
}
type DB struct {
	client *mongo.Client
}

func New(ctx context.Context, cfg Config) *DB {
	client, cErr := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if cErr != nil {
		panic(cErr)
	}

	return &DB{client: client}
}

func (d DB) GetClient() *mongo.Client {
	return d.client
}
