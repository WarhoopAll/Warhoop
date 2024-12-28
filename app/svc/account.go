package svc

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/store"
	"warhoop/app/svc/web"
)

func NewService(ctx context.Context, store *store.Store, logger *log.Logger) *web.WebService {
	return web.New(ctx, store, logger)
}
