package store

import (
	"context"
	"github.com/eko/gocache/lib/v4/cache"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type NewStoreParameter struct {
	fx.In
	Core   *cache.Cache[[]byte]
	Ctx    context.Context
	Logger *zap.SugaredLogger
}

func NewStore[T any](
	param NewStoreParameter,
) *Store[T] {
	core, ctx, logger := param.Core, param.Ctx, param.Logger
	return &Store[T]{
		core:   core,
		ctx:    ctx,
		logger: logger,
	}
}
