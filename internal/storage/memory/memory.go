package memory

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/usawyer/url-shortener/internal/models"
)

const cacheDuration = 1 * time.Minute

type RedisClient struct {
	rd *redis.Client
}

func (c *RedisClient) AddUrl(ctx context.Context, urls models.Urls) error {
	return c.rd.Set(ctx, urls.Alias, urls.Url, cacheDuration).Err()
}

func (c *RedisClient) GetUrl(ctx context.Context, key string) (string, error) {
	value, err := c.rd.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", fmt.Errorf("URL with alias \"%s\" doesn't found", key)
		}
		return "", err
	}
	return value, nil
}
