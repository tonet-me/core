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
	visithandler "github.com/tonet-me/tonet-core/delivery/http/visit"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	cardmongo "github.com/tonet-me/tonet-core/repository/mongo/card"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
	visitmongo "github.com/tonet-me/tonet-core/repository/mongo/visit"
	"github.com/tonet-me/tonet-core/service/auth"
	cardservice "github.com/tonet-me/tonet-core/service/card"
	userservice "github.com/tonet-me/tonet-core/service/user"
	visitservice "github.com/tonet-me/tonet-core/service/visit"
	cardvalidator "github.com/tonet-me/tonet-core/validator/card"
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
	visitHandler := creatVisitHandler(cfg, mongoClient)

	e := echo.New()
	server := httpserver.New(cfg.HttpServer, e, userHandler, cardHandler, minioHandler, visitHandler)
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
	cardSvc := cardservice.New(cfg.CardConfig, cardDB)
	cardVld := cardvalidator.New(cardDB)
	return cardhandler.New(cardSvc, cardVld, authGenerator, cfg.Auth)
}

func createMinioHandler(cfg config.Config, authGenerator auth.Service) httpserver.Handler {
	minioClient := minio.New(cfg.Minio)

	return miniohandler.New(minioClient, authGenerator, cfg.Auth)
}

func creatVisitHandler(cfg config.Config, client *mongodb.DB) httpserver.Handler {
	cardDB := cardmongo.New(cfg.CardMongo, client)
	cardSvc := cardservice.New(cfg.CardConfig, cardDB)

	visitDB := visitmongo.New(cfg.VisitConfig, client)
	visitSvc := visitservice.New(visitDB, cardSvc)

	return visithandler.New(visitSvc)
}
