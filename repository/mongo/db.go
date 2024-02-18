package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	Username string `koanf:"username"`
	Password string `koanf:"password"`
}

type DB struct {
	client *mongo.Client
}

func New(cfg Config) *DB {
	URI := fmt.Sprintf(`mongodb://%s:%s@%s:%d/`, cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}
	client, cErr := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI).SetBSONOptions(bsonOpts))
	if cErr != nil {
		panic(cErr)
	}

	// Check the connection
	pErr := client.Ping(context.TODO(), nil)
	if pErr != nil {
		panic(pErr)
	}

	return &DB{
		client: client,
	}
}

func (d DB) GetClient() *mongo.Client {
	return d.client
}
