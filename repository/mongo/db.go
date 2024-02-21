package mongodb

import (
	"context"
	"fmt"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
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
	const op = richerror.OP("mongodb.New")

	URI := fmt.Sprintf(`mongodb://%s:%s@%s:%d/`, cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}
	client, cErr := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI).SetBSONOptions(bsonOpts))
	if cErr != nil {
		panic(fmt.Errorf("op:%v,\nwith err:%v", op, cErr))

	}

	// Check the connection
	pErr := client.Ping(context.TODO(), nil)
	if pErr != nil {
		panic(fmt.Errorf("op:%v,\nwith err:%v", op, pErr))
	}

	return &DB{
		client: client,
	}
}

func (d DB) GetClient() *mongo.Client {
	return d.client
}
