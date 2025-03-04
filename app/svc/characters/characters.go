package characters

import (
	"context"
	"warhoop/app/cache"
	"warhoop/app/log"
	"warhoop/app/store"
)

type CharService struct {
	ctx        context.Context
	store      *store.Store
	logger     *log.Logger
	redisCache *cache.RedisCache
}

// New creates a new account service
func New(ctx context.Context, store *store.Store, logger *log.Logger, redisCache *cache.RedisCache) *CharService {
	return &CharService{
		ctx:        ctx,
		store:      store,
		logger:     logger,
		redisCache: redisCache,
	}
}
