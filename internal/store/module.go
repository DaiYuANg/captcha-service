package store

import (
	"captcha-service/internal/config"
	"captcha-service/internal/constant"
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	bigcachestore "github.com/eko/gocache/store/bigcache/v4"
	redisstore "github.com/eko/gocache/store/redis/v4"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
)

var Module = fx.Module("store",
	fx.Provide(
		storeContext,
		newStoreCore,
	),
)

func newStoreCore(config *config.Config, ctx context.Context, logger *zap.SugaredLogger) (*cache.Cache[[]byte], error) {
	var backend store.StoreInterface
	logger.Debugf("current backend %s", config.Store.Type)
	switch config.Store.Type {
	case constant.Memory:
		bgConf := bigcache.DefaultConfig(2 * time.Minute)
		bgConf.StatsEnabled = true
		bigcacheClient, _ := bigcache.New(ctx, bgConf)
		backend = bigcachestore.NewBigcache(bigcacheClient)
	case constant.Redis:
		redisClient, err := newRedisClient(ctx, &redis.Options{
			Addr:         config.Store.Url,
			PoolSize:     100,
			MinIdleConns: 5,
			DialTimeout:  2 * time.Second,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
		})
		if err != nil {
			return nil, err
		}
		backend = redisstore.NewRedis(
			redisClient,
		)
	}
	return cache.New[[]byte](backend), nil
}

func storeContext() context.Context {
	return context.Background()
}
