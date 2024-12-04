package svc

import (
	"context"
	"grimoire/app/log"
	"grimoire/app/store"
	"grimoire/app/svc/web"
	"grimoire/app/utils"
)

type Manager struct {
	Account Account
	Sait    Sait
	Char    Characters
	Logger  *log.Logger
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store, logger *log.Logger) (*Manager, error) {
	if store == nil {
		return nil, utils.ErrNoData
	}
	return &Manager{
		Logger:  logger,
		Account: web.New(ctx, store, logger),
		Char:    web.New(ctx, store, logger),
		Sait:    web.New(ctx, store, logger),
	}, nil
}
