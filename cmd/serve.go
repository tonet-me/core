package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/adapter/minio"
	"github.com/tonet-me/tonet-core/config"
	httpserver "github.com/tonet-me/tonet-core/delivery/http"
	cardhandler "github.com/tonet-me/tonet-core/delivery/http/card"
	miniohandler "github.com/tonet-me/tonet-core/delivery/http/minio"
	userhandler "github.com/tonet-me/tonet-core/delivery/http/user"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	cardmongo "github.com/tonet-me/tonet-core/repository/mongo/card"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
	"github.com/tonet-me/tonet-core/service/auth"
	cardservice "github.com/tonet-me/tonet-core/service/card"
	userservice "github.com/tonet-me/tonet-core/service/user"
)

type Serve struct {
	Handlers []httpserver.Handler
}

func StartServe(cfg config.Config) {
	mongoClient := mongodb.New(cfg.MongoClient)
	minioHandler := createMinioHandler(cfg)
	userHandler := createUserHandler(cfg, mongoClient)
	cardHandler := createCardHandler(cfg, mongoClient)
	e := echo.New()
	server := httpserver.New(cfg.HttpServer, e, userHandler, cardHandler, minioHandler)
	server.StartListening()
}

func createUserHandler(cfg config.Config, client *mongodb.DB) httpserver.Handler {
	userDB := usermongo.New(cfg.UserMongo, client)
	authGenerator := auth.New(cfg.Auth)
	oauth := new(userservice.OAuthService)
	userSvc := userservice.New(userDB, authGenerator, *oauth)
	return userhandler.New(userSvc)
}

func createCardHandler(cfg config.Config, client *mongodb.DB) httpserver.Handler {
	cardDB := cardmongo.New(cfg.CardMongo, client)
	cardSvc := cardservice.New(cardDB)
	return cardhandler.New(cardSvc)
}

func createMinioHandler(cfg config.Config) httpserver.Handler {
	minioClient := minio.New(cfg.Minio)

	return miniohandler.New(minioClient)
}
