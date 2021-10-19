package clients

import (
	"context"
	"fmt"

	"github.com/HotPotatoC/go-microservice/internal/apps/postservice/config"
	"github.com/HotPotatoC/go-microservice/pkg/database"
	"github.com/HotPotatoC/go-microservice/pkg/kv"
	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/errgroup"
)

// Clients contains dependencies to fetch data from resources.
type Clients struct {
	Redis    kv.KV
	Database database.SQL
}

// NewClients creates a new Clients.
func NewClients(ctx context.Context, conf *config.Config) (Clients, error) {
	var g errgroup.Group

	c := Clients{}

	g.Go(func() error {
		var err error
		c.Redis, err = kv.NewRedisClient(ctx, &redis.Options{
			Addr:         conf.Redis.Addr,
			Password:     conf.Redis.Password,
			DB:           conf.Redis.DB,
			WriteTimeout: conf.Redis.WriteTimeout,
			ReadTimeout:  conf.Redis.ReadTimeout,
		})
		if err != nil {
			return fmt.Errorf("failed to connect to redis: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		var err error
		c.Database, err = database.NewPostgreSQL(ctx, conf.Database.DSN)
		if err != nil {
			return fmt.Errorf("failed to connect to database: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return Clients{}, err
	}

	return c, nil
}
