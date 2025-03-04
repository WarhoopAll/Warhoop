package svc

import (
	"context"
	"warhoop/app/cache"
	"warhoop/app/log"
	"warhoop/app/store"
	"warhoop/app/svc/auth"
	"warhoop/app/svc/characters"
	"warhoop/app/svc/soap"
	"warhoop/app/svc/web"
	"warhoop/app/utils"
)

type Manager struct {
	Auth       Auth
	Web        Web
	Char       Characters
	Soap       Soap
	Logger     *log.Logger
	redisCache *cache.RedisCache
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store, logger *log.Logger, redisCache *cache.RedisCache) (*Manager, error) {
	if store == nil {
		return nil, utils.ErrNoData
	}

	webSvc := web.New(ctx, store, logger, redisCache)
	soapSvc := soap.New(logger)

	return &Manager{
		Logger: logger,
		Auth:   auth.New(ctx, store, logger, webSvc),
		Char:   characters.New(ctx, store, logger, redisCache),
		Web:    webSvc,
		Soap:   soapSvc,
	}, nil
}
