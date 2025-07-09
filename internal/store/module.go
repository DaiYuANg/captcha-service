package store

import (
	"captcha-service/internal/config"
	"captcha-service/internal/constant"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/storage/badger/v2"
	"github.com/gofiber/storage/bbolt/v2"
	"github.com/gofiber/storage/etcd/v2"
	"github.com/gofiber/storage/memory/v2"
	"github.com/gofiber/storage/mysql/v2"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/storage/redis/v3"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/gofiber/storage/valkey"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
	"path"
	"runtime"
	"time"
)

var Module = fx.Module("store",
	fx.Provide(
		newStoreCore,
	),
)

func newStoreCore(config *config.Config, logger *zap.SugaredLogger) fiber.Storage {
	temp := os.TempDir()
	logger.Debugf("current backend %s", config.Store.Type)
	logger.Debugf("temp dir:%s", temp)

	commonGCInterval := 10 * time.Second

	// 构建后端类型与初始化函数的映射
	backendFactories := map[constant.StoreType]func() fiber.Storage{
		constant.Memory: func() fiber.Storage {
			return memory.New(memory.Config{
				GCInterval: commonGCInterval,
			})
		},
		constant.Badger: func() fiber.Storage {
			return badger.New(badger.Config{
				Database:   path.Join(temp, "captcha.badger"),
				GCInterval: commonGCInterval,
			})
		},
		constant.Etcd: func() fiber.Storage {
			return etcd.New(etcd.Config{
				Endpoints: []string{config.Store.Url},
			})
		},
		constant.Sqlite: func() fiber.Storage {
			return sqlite3.New(sqlite3.Config{
				Database:        path.Join(temp, "captcha.sqlite"),
				Table:           "fiber_storage",
				GCInterval:      commonGCInterval,
				MaxOpenConns:    100,
				MaxIdleConns:    100,
				ConnMaxLifetime: 1 * time.Second,
			})
		},
		constant.Bbolt: func() fiber.Storage {
			return bbolt.New(bbolt.Config{
				Database: path.Join(temp, "captcha.db"),
				Bucket:   "captcha",
				Reset:    true,
			})
		},
		constant.Valkey: func() fiber.Storage {
			return valkey.New(valkey.Config{
				InitAddress: []string{config.Store.Url},
				Username:    "",
				Password:    "",
				TLSConfig:   nil,
			})
		},
		constant.Redis: func() fiber.Storage {
			return redis.New(redis.Config{
				Host:      config.Store.Url,
				Port:      config.Store.Port,
				Username:  config.Store.Username,
				Password:  config.Store.Password,
				Database:  0,
				TLSConfig: nil,
				PoolSize:  10 * runtime.GOMAXPROCS(0),
			})
		},
		constant.Postgres: func() fiber.Storage {
			return postgres.New(postgres.Config{
				ConnectionURI: config.Store.Url,
				Table:         "fiber_storage",
				GCInterval:    commonGCInterval,
			})
		},
		constant.Mysql: func() fiber.Storage {
			return mysql.New(mysql.Config{
				Host:       config.Store.Url,
				Port:       config.Store.Port,
				Database:   config.Store.Database,
				Table:      "captcha-" + time.Now().Format("20060102150405"), // 避免特殊字符
				GCInterval: commonGCInterval,
			})
		},
	}

	if factory, ok := backendFactories[config.Store.Type]; ok {
		return factory()
	}

	// fallback，避免 nil 返回导致 panic
	logger.Panicf("unsupported backend type: %s", config.Store.Type)
	return nil
}
