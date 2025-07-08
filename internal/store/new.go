package store

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type NewStoreParameter struct {
	fx.In
	Core   fiber.Storage
	Logger *zap.SugaredLogger
}

func NewStore[T any](
	param NewStoreParameter,
) *Store[T] {
	core, logger := param.Core, param.Logger
	return &Store[T]{
		core:   core,
		logger: logger,
	}
}
