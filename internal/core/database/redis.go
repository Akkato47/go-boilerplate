package database

import (
	"context"
	"log/slog"

	"github.com/Akkato47/go-boilerplate/internal/core/config"
	"github.com/redis/go-redis/v9"
)

func CreateRedisClient(ctx context.Context, config *config.RedisConfig) (*redis.Client, error) {
	opts, err := redis.ParseURL(config.URL)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, err
	}

	slog.Info("Successfully connected to Redis")
	return client, nil
}
