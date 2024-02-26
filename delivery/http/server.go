package httpserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
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

	log.Println("server shutting down in %s...", s.gracefulShutdown) //replace with slog
	c, cancel := context.WithTimeout(context.Background(), s.gracefulShutdown)
	defer cancel()

	if err := s.echo.Shutdown(c); err != nil {
		log.Fatal("server shutdown", err)
	}

	<-c.Done()
	log.Println("Good Luck!")
}
