package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/adapter/minio"
	"github.com/tonet-me/tonet-core/adapter/oauth"
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
	uservalidator "github.com/tonet-me/tonet-core/validator/user"
)

type Serve struct {
	Handlers []httpserver.Handler
}

func StartServe(cfg config.Config) {
	mongoClient := mongodb.New(cfg.MongoClient)
	authGenerator := auth.New(cfg.Auth)

	minioHandler := createMinioHandler(cfg, authGenerator)
	userHandler := createUserHandler(cfg, mongoClient, authGenerator)
	cardHandler := createCardHandler(cfg, mongoClient, authGenerator)

	e := echo.New()
	server := httpserver.New(cfg.HttpServer, e, userHandler, cardHandler, minioHandler)
	server.StartListening()
}

func createUserHandler(cfg config.Config, client *mongodb.DB, authGenerator auth.Service) httpserver.Handler {
	userDB := usermongo.New(cfg.UserMongo, client)
	googleOauth := oauth.NewGoogle(cfg.OAuth.Google)
	oAuthAdapter := oauth.New(googleOauth)
	userValidator := uservalidator.New()
	userSvc := userservice.New(userDB, authGenerator, oAuthAdapter)

	return userhandler.New(userSvc, userValidator, authGenerator, cfg.Auth)
}

func createCardHandler(cfg config.Config, client *mongodb.DB, authGenerator auth.Service) httpserver.Handler {
	cardDB := cardmongo.New(cfg.CardMongo, client)
	cardSvc := cardservice.New(cardDB)

	return cardhandler.New(cardSvc, authGenerator, cfg.Auth)
}

func createMinioHandler(cfg config.Config, authGenerator auth.Service) httpserver.Handler {
	minioClient := minio.New(cfg.Minio)

	return miniohandler.New(minioClient, authGenerator, cfg.Auth)
}
