package memory

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/usawyer/url-shortener/internal/models"
	"time"
)

const cacheDuration = 1 * time.Minute

type redisClient struct {
	rd *redis.Client
}

func (c *redisClient) AddUrl(ctx context.Context, urls models.Urls) error {
	return c.rd.Set(ctx, urls.Alias, urls.Url, cacheDuration).Err()
}

func (c *redisClient) GetUrl(ctx context.Context, key string) (string, error) {
	return c.rd.Get(ctx, key).Result()
}
