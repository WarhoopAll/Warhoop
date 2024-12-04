package svc

import (
	"context"
	"grimoire/app/log"
	"grimoire/app/store"
	"grimoire/app/svc/web"
)

func NewService(ctx context.Context, store *store.Store, logger *log.Logger) *web.WebService {
	return web.New(ctx, store, logger)
}
