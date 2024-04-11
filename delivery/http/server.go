package httpserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tonet-me/tonet-core/logger"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Host             string        `koanf:"host"`
	Port             int           `koanf:"port"`
	gracefulShutdown time.Duration `koanf:"graceful_shutdown"`
}
type Server struct {
	addr             string
	gracefulShutdown time.Duration
	echo             *echo.Echo
	handlers         []Handler
}

func New(cfg Config, echo *echo.Echo, handlers ...Handler) *Server {
	return &Server{
		addr:             fmt.Sprintf("%s:%v", cfg.Host, cfg.Port),
		echo:             echo,
		gracefulShutdown: cfg.gracefulShutdown,
		handlers:         handlers,
	}
}

func (s *Server) StartListening() {
	s.echo.Use(middleware.CORS())

	sLogger := logger.GetLogger()
	s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				sLogger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST_SUCCESS",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {

				var slogLevel slog.Level
				var slogMsg string

				if v.Status >= 500 {
					slogLevel = slog.LevelError
					slogMsg = "REQUEST_ERROR"
				} else if v.Status >= 200 && v.Status <= 299 {
					slogLevel = slog.LevelInfo
					slogMsg = "REQUEST_SUCCESS"
				} else {
					slogLevel = slog.LevelInfo
					slogMsg = "REQUEST_INFO"
				}

				sLogger.LogAttrs(context.Background(), slogLevel, slogMsg,
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err-msg", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	for _, handler := range s.handlers {
		handler.SetRoutes(s.echo)
	}

	go func() {
		if err := s.echo.Start(s.addr); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("server start", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	<-quit

	logger.GetLogger().Warn("server graceful shutdown", slog.Any("server shutting down in", s.gracefulShutdown))

	c, cancel := context.WithTimeout(context.Background(), s.gracefulShutdown)
	defer cancel()

	if err := s.echo.Shutdown(c); err != nil {
		log.Fatal("server shutdown", err)
	}

	<-c.Done()
	logger.GetLogger().Info("Good Luck!")

	log.Println("Good Luck!")
}
