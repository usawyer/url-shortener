package db

import (
	"context"
	"github.com/usawyer/url-shortener/internal/models"
	"gorm.io/gorm"
)

type postgresClient struct {
	db *gorm.DB
}

func (p *postgresClient) AddUrl(ctx context.Context, urls models.Urls) error {
	return nil
}

func (p *postgresClient) GetUrl(ctx context.Context, key string) (string, error) {
	return "", nil
}
