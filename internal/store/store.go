package store

import (
	"context"
	"encoding/json"
	"github.com/eko/gocache/lib/v4/cache"
	"go.uber.org/zap"
)

type Store[T any] struct {
	core   *cache.Cache[[]byte]
	ctx    context.Context
	logger *zap.SugaredLogger
}

func (s *Store[T]) Set(key string, value T) error {
	s.logger.Infof("Store.Set called key=%s", key)

	data, err := json.Marshal(value)
	if err != nil {
		s.logger.Errorf("Store.Set marshal error key=%s err=%v", key, err)
		return err
	}

	err = s.core.Set(s.ctx, key, data)
	if err != nil {
		s.logger.Errorf("Store.Set core.Set error key=%s err=%v", key, err)
		return err
	}

	s.logger.Infof("Store.Set success key=%s", key)
	return nil
}

func (s *Store[T]) Get(key string) (T, error) {
	s.logger.Infof("Store.Get called key=%s", key)

	var result T

	data, err := s.core.Get(s.ctx, key)
	if err != nil {
		s.logger.Errorf("Store.Get core.Get error key=%s err=%v", key, err)
		var nilVal T
		return nilVal, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		s.logger.Errorf("Store.Get unmarshal error key=%s err=%v", key, err)
		var nilVal T
		return nilVal, err
	}

	s.logger.Infof("Store.Get success key=%s", key)
	return result, nil
}

func (s *Store[T]) Delete(key string) error {
	s.logger.Infof("Store.Delete called key=%s", key)

	err := s.core.Delete(s.ctx, key)
	if err != nil {
		s.logger.Errorf("Store.Delete error key=%s err=%v", key, err)
		return err
	}

	s.logger.Infof("Store.Delete success key=%s", key)
	return nil
}
