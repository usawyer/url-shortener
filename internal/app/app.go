package app

import (
	"github.com/usawyer/url-shortener/internal/config"
	"github.com/usawyer/url-shortener/internal/handler"
	"github.com/usawyer/url-shortener/internal/router"
	"github.com/usawyer/url-shortener/internal/service"
	"github.com/usawyer/url-shortener/internal/storage"
	"go.uber.org/zap"
	"net"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	Router *router.Router
}

func New(logger *zap.Logger, config *config.Config, storeType string) *App {
	logger = logger.Named("App")

	repository := storage.New(storeType)
	service := service.New(repository)
	handler := handler.New(service)
	router := router.New(handler)

	return &App{
		Config: config,
		Logger: logger,
		Router: router,
	}

}

func (a *App) Run() {
	a.Logger.Info("starting server")
	if err := a.Router.Run(net.JoinHostPort(a.Config.Host, a.Config.Port)); err != nil {
		a.Logger.Fatal(err.Error())
	}
}
