package usermongo

import (
	"context"
	"fmt"
	"github.com/tonet-me/tonet-core/logger"
	errmsg "github.com/tonet-me/tonet-core/pkg/err_msg"
	richerror "github.com/tonet-me/tonet-core/pkg/rich_error"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
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
	const op = richerror.OP("usermongo.initialCollection")

	collections, lErr := client.GetClient().Database(cfg.DBName).ListCollectionNames(context.Background(), bson.D{{}})
	if lErr != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, lErr.Error()))

		panic(fmt.Errorf("op:%v,\nwith err:%v", op, lErr))
	}
	if !slices.Contains(collections, cfg.CollName) {
		err := client.GetClient().Database(cfg.DBName).CreateCollection(context.TODO(), cfg.CollName)

		if err != nil {
			logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, err.Error()))

			panic(fmt.Errorf("op:%v,\nwith err:%v", op, err))
		}

	}

	userCollection := client.GetClient().Database(cfg.DBName).Collection(cfg.CollName)
	indexModelEmail := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}

	indexModelPhoneNumber := mongo.IndexModel{
		Keys: bson.D{{"phone_number", 1}},
	}
	indexModelEmailAndStatus := mongo.IndexModel{
		Keys: bson.D{{"email", 1}, {"status", 1}},
	}
	_, iErr := userCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{indexModelEmail, indexModelPhoneNumber, indexModelEmailAndStatus})
	if iErr != nil {
		logger.GetLogger().Error(string(op), slog.String(errmsg.ErrorMsg, iErr.Error()))

		panic(fmt.Errorf("op:%v,\nwith err:%v", op, iErr))
	}

	return userCollection
}
