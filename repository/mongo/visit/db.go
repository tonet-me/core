package visitmongo

import (
	"context"
	"fmt"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"slices"
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
	const op = richerror.OP("visitmongo.initialCollection")

	collections, lErr := client.GetClient().Database(cfg.DBName).ListCollectionNames(context.TODO(), bson.D{{}})
	if lErr != nil {
		panic(fmt.Errorf("op:%v,\nwith err:%v", op, lErr))
	}
	if !slices.Contains(collections, cfg.CollName) {
		err := client.GetClient().Database(cfg.DBName).CreateCollection(context.TODO(), cfg.CollName)
		if err != nil {
			panic(fmt.Errorf("op:%v,\nwith err:%v", op, err))
		}
	}
	cardCollection := client.GetClient().Database(cfg.DBName).Collection(cfg.CollName)

	indexModelCardName := mongo.IndexModel{
		Keys: bson.D{{"card_id", 1}},
	}

	_, iErr := cardCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{indexModelCardName})
	if iErr != nil {
		panic(fmt.Errorf("op:%v,\nwith err:%v", op, iErr))
	}

	return cardCollection
}
