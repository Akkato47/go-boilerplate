package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Akkato47/go-boilerplate/internal/core/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, config *config.PostgresConfig) (*pgxpool.Pool, error) {
	var connURL string
	if config.URL == "" {
		connURL = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", config.Host, config.Port, config.User, config.Password, config.Name, config.SslMode)
	} else {
		connURL = config.URL
	}
	pool, err := pgxpool.New(ctx, connURL)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}

	slog.Info("Successfully connected to PSQL db")
	return pool, nil
}
