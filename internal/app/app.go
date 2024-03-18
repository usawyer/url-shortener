package app

import (
	"net"

	"github.com/usawyer/url-shortener/internal/config"
	"github.com/usawyer/url-shortener/internal/handler"
	"github.com/usawyer/url-shortener/internal/router"
	"github.com/usawyer/url-shortener/internal/service"
	"github.com/usawyer/url-shortener/internal/storage"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	Router *router.Router
}

func New(logger *zap.Logger, cfg *config.Config, storeType string) *App {
	logger = logger.Named("App")

	repository := storage.New(storeType, logger, cfg)
	srvc := service.New(repository, logger)
	hndlr := handler.New(srvc, logger)
	rtr := router.New(hndlr)

	return &App{
		Config: cfg,
		Logger: logger,
		Router: rtr,
	}

}

func (a *App) Run() {
	a.Logger.Info("starting server")
	if err := a.Router.Run(net.JoinHostPort(a.Config.Host, a.Config.Port)); err != nil {
		a.Logger.Fatal(err.Error())
	}
}
