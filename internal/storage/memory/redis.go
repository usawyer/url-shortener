package memory

import (
	"context"
	"net"

	"github.com/redis/go-redis/v9"
	"github.com/usawyer/url-shortener/internal/config"
	"go.uber.org/zap"
)

func New(logger *zap.Logger, cfg *config.Config) (*RedisClient, error) {
	logger = logger.Named("Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(cfg.RdHost, cfg.RdPort),
		Password: cfg.RdPassword,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("connected to Redis")
	}

	return &RedisClient{rd: client}, err
}
