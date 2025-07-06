package store

import (
	"context"
	"encoding/json"
	"github.com/eko/gocache/lib/v4/cache"
)

type Store[T any] struct {
	core *cache.Cache[[]byte]
	ctx  context.Context
}

func newStore(core *cache.Cache[[]byte], ctx context.Context) *Store[any] {
	return &Store[any]{
		core: core,
		ctx:  ctx,
	}
}

func (s *Store[T]) Set(key string, value T) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.core.Set(s.ctx, key, data)
}

func (s *Store[T]) Get(key string) (T, error) {
	var result T

	data, err := s.core.Get(s.ctx, key)
	if err != nil {
		var nilVal T
		return nilVal, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		var nilVal T
		return nilVal, err
	}
	return result, nil
}

func (s *Store[T]) Delete(key string) error {
	return s.core.Delete(s.ctx, key)
}
