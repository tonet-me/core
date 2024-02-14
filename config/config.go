package config

import (
	httpserver "github.com/tonet-me/tonet-core/delivery/http"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
	"github.com/tonet-me/tonet-core/service/auth"
)

type Config struct {
	HttpServer  httpserver.Config `koanf:"http_server"`
	MongoClient mongodb.Config    `koanf:"mongo_client"`
	UserMongo   usermongo.Config  `koanf:"user_mongo"`
	Auth        auth.Config       `koanf:"auth"`
}
