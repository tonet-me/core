package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/tonet-me/tonet-core/config"
	httpserver "github.com/tonet-me/tonet-core/delivery/http"
	userhandler "github.com/tonet-me/tonet-core/delivery/http/user"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	usermongo "github.com/tonet-me/tonet-core/repository/mongo/user"
	"github.com/tonet-me/tonet-core/service/auth"
	userservice "github.com/tonet-me/tonet-core/service/user"
)

type Serve struct {
	Handlers []httpserver.Handler
}

func StartServe(cfg config.Config) {
	mongoClient := mongodb.New(cfg.MongoClient)
	userHandler := createUserHandler(cfg, mongoClient)
	e := echo.New()
	server := httpserver.New(cfg.HttpServer, e, userHandler)
	server.StartListening()
}

func createUserHandler(cfg config.Config, client *mongodb.DB) httpserver.Handler {
	userDB := usermongo.New(cfg.UserMongo, client)
	authGenerator := auth.New(cfg.Auth)
	oauth := new(userservice.OAuthService)
	userSvc := userservice.New(userDB, authGenerator, *oauth)
	return userhandler.New(userSvc)

}
