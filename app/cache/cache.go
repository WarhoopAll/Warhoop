package cache

import (
	"context"
	"time"
	"warhoop/app/config"
)

type Cache interface {
	Get(ctx context.Context, key string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

type CacheManager struct {
	Cache Cache
}

func NewCacheManager() *CacheManager {
	if config.Get().RedisEnable {
		return &CacheManager{Cache: NewRedisCache()}
	}
	return &CacheManager{Cache: &NoopCache{}}
}
