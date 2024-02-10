package config

import (
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
)

func Default() Config {
	cfx := Config{
		MongoClient: mongodb.Config{
			Host:     "localhost",
			Port:     27017,
			Username: "root",
			Password: "rootpassword",
		},
		UserMongo: usermongo.Config{
			DBName:   "tonet",
			CollName: "users",
		},
	}

	return cfx
}
