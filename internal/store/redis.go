package store

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func newRedisClient(ctx context.Context, options *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(options)
	ping := client.Ping(ctx)
	if ping.Err() != nil {
		return nil, ping.Err()
	}

	return client, nil
}
