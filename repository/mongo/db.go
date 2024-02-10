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
	config Config
	client *mongo.Client
}

func New(cfg Config) *DB {
	client, cErr := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.URI))
	if cErr != nil {
		panic(cErr)
	}

	// Check the connection
	pErr := client.Ping(context.TODO(), nil)
	if pErr != nil {
		panic(pErr)
	}

	return &DB{
		config: cfg,
		client: client,
	}
}

func (d DB) GetClient() *mongo.Client {
	return d.client
}
