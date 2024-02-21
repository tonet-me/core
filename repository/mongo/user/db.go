package usermongo

import (
	"context"
	"fmt"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DBName   string `koanf:"db_name"`
	CollName string `koanf:"coll_name"`
}

type DB struct {
	collection *mongo.Collection
}

func New(cfg Config, client *mongodb.DB) *DB {
	return &DB{
		collection: initialCollection(cfg, client),
	}
}

func initialCollection(cfg Config, client *mongodb.DB) *mongo.Collection {
	const op = richerror.OP("usermongo.initialCollection")

	err := client.GetClient().Database(cfg.DBName).CreateCollection(context.TODO(), cfg.CollName)
	if err != nil {
		panic(fmt.Errorf("op:%v,\nwith err:%v", op, err))
	}

	userCollection := client.GetClient().Database(cfg.DBName).Collection(cfg.CollName)
	indexModelEmail := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	indexModelPhoneNumber := mongo.IndexModel{
		Keys:    bson.D{{"phone_number", 1}},
		Options: options.Index().SetUnique(true),
	}
	indexModelEmailAndStatus := mongo.IndexModel{
		Keys: bson.D{{"email", 1}, {"status", 1}},
	}
	_, iErr := userCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{indexModelEmail, indexModelPhoneNumber, indexModelEmailAndStatus})
	if iErr != nil {
		panic(fmt.Errorf("op:%v,\nwith err:%v", op, iErr))
	}

	return userCollection
}
