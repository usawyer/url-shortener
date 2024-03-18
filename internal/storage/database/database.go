package db

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/usawyer/url-shortener/internal/models"
	"gorm.io/gorm"
)

var (
	ErrUrlExist   = errors.New("url already exists")
	ErrAliasExist = errors.New("collision: alias already exists")
)

type PostgresClient struct {
	db *gorm.DB
}

func (p *PostgresClient) AddUrl(ctx context.Context, urls models.Urls) error {
	if p.findUrl(urls.Url, "url") {
		return ErrUrlExist
	}

	if p.findUrl(urls.Alias, "alias") {
		return ErrAliasExist
	}

	return p.db.WithContext(ctx).Create(&urls).Error
}

func (p *PostgresClient) GetUrl(ctx context.Context, key string) (string, error) {
	var url models.Urls
	if err := p.db.WithContext(ctx).Where("alias = ?", key).First(&url).Error; err != nil {
		return "", err
	}
	return url.Url, nil
}

func (p *PostgresClient) findUrl(str string, columnName string) bool {
	var url models.Urls
	condition := fmt.Sprintf("%s = ?", columnName)
	res := p.db.Where(condition, str).Find(&url)
	return res.RowsAffected == 1
}
