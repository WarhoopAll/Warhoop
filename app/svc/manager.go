package svc

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/store"
	"warhoop/app/svc/auth"
	"warhoop/app/svc/characters"
	"warhoop/app/svc/web"
	"warhoop/app/utils"
)

type Manager struct {
	Auth   Auth
	Web    Web
	Char   Characters
	Logger *log.Logger
}

// NewManager creates new service manager
func NewManager(ctx context.Context, store *store.Store, logger *log.Logger) (*Manager, error) {
	if store == nil {
		return nil, utils.ErrNoData
	}
	return &Manager{
		Logger: logger,
		Auth:   auth.New(ctx, store, logger),
		Char:   characters.New(ctx, store, logger),
		Web:    web.New(ctx, store, logger),
	}, nil
}
