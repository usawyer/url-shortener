package db

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/usawyer/url-shortener/internal/config"
	"github.com/usawyer/url-shortener/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func New(logger *zap.Logger, cfg *config.Config) (*PostgresClient, error) {
	logger = logger.Named("PostgreSQL")
	dsn := makeDsnStr(cfg)

	db, err := getClient(dsn, logger)
	if err != nil {
		return nil, err
	}

	db = db.Debug()
	err = db.AutoMigrate(&models.Urls{})
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("migrate ok")
	}

	return &PostgresClient{db: db}, err
}

func makeDsnStr(cfg *config.Config) string {
	parameters := map[string]string{
		"host":     cfg.DbHost,
		"user":     cfg.PgUser,
		"password": cfg.PgPassword,
		"dbname":   cfg.DbName,
		"port":     cfg.DbPort,
		"sslmode":  "disable",
	}

	var pairs []string
	for key, value := range parameters {
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
	}
	return strings.Join(pairs, " ")
}

func getClient(dsn string, logger *zap.Logger) (*gorm.DB, error) {
	var value time.Duration = 1
	ticker := time.NewTicker(value * time.Nanosecond)
	timeout := time.After(1 * time.Minute)
	gormLogger := zapgorm2.New(logger)

	for {
		select {
		case <-ticker.C:
			ticker.Stop()

			client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})

			if err != nil {
				logger.With(zap.Error(err)).Warn("fail to set connection with PostgreSQL")
				ticker = time.NewTicker(value * time.Second)
				value *= 2
				continue
			}
			return client, nil
		case <-timeout:
			return nil, errors.New("PostgreSQL: connection timeout")
		}
	}
}
