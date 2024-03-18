package storage

import (
	"context"
	"github.com/usawyer/url-shortener/internal/config"
	"github.com/usawyer/url-shortener/internal/models"
	db "github.com/usawyer/url-shortener/internal/storage/database"
	"github.com/usawyer/url-shortener/internal/storage/memory"
	"go.uber.org/zap"
)

type Storage interface {
	AddUrl(context.Context, models.Urls) error
	GetUrl(context.Context, string) (string, error)
}

func New(storageType string, logger *zap.Logger, cfg *config.Config) Storage {
	var store Storage
	var err error
	switch storageType {
	case "memory":
		store, err = memory.New(logger, cfg)
		if err != nil {
			logger.Fatal(err.Error())
		}
	case "db":
		store, err = db.New(logger, cfg)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	return store
}
