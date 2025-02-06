package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/internal/config"
	"github.com/Eugene-Usachev/test-task-for-VK/backend/src/pkg"
	"github.com/jackc/pgx/v5/pgxpool"
)

func MustNewPool(cfg config.PostgresConfig) *pgxpool.Pool {
	var (
		pool *pgxpool.Pool
		err  error
	)

	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.User(), cfg.Password(), cfg.Host(), cfg.Port(), cfg.Database())

	err = pkg.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, url)
		if err != nil {
			return err
		}

		return pool.Ping(ctx)
	}, 150, time.Millisecond*100)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	return pool
}
