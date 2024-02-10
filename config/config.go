package config

import (
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
)

type Config struct {
	MongoClient mongodb.Config   `koanf:"mongo_client"`
	UserMongo   usermongo.Config `koanf:"user_mongo"`
}
