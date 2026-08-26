package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/korniykom/Chatters-Backend-Go/internal/config"
)

func New(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}

	connConfig.ConnConfig.Host = cfg.DBHost
	connConfig.ConnConfig.Port = 5432
	connConfig.ConnConfig.User = cfg.DBUser
	connConfig.ConnConfig.Password = cfg.DBPassword
	connConfig.ConnConfig.Database = cfg.DBName

	db, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
