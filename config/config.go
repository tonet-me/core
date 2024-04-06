package config

import (
	"github.com/tonet-me/tonet-core/adapter/minio"
	oauth "github.com/tonet-me/tonet-core/adapter/oauth"
	httpserver "github.com/tonet-me/tonet-core/delivery/http"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	cardmongo "github.com/tonet-me/tonet-core/repository/mongo/card"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
	visitmongo "github.com/tonet-me/tonet-core/repository/mongo/visit"
	"github.com/tonet-me/tonet-core/service/auth"
	cardservice "github.com/tonet-me/tonet-core/service/card"
)

type Config struct {
	HttpServer  httpserver.Config  `koanf:"http_server"`
	MongoClient mongodb.Config     `koanf:"mongo_client"`
	UserMongo   usermongo.Config   `koanf:"user_mongo"`
	CardMongo   cardmongo.Config   `koanf:"card_mongo"`
	VisitConfig visitmongo.Config  `koanf:"visit_mongo"`
	Auth        auth.Config        `koanf:"auth"`
	Minio       minio.Config       `koanf:"minio"`
	OAuth       oauth.Config       `koanf:"o_auth"`
	CardConfig  cardservice.Config `koanf:"card_config"`
}
