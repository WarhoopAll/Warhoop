package web

import (
	"context"
	"grimoire/app/log"
	"grimoire/app/store"
	"sync"
)

type WebService struct {
	ctx    context.Context
	store  *store.Store
	logger *log.Logger
	mutex  sync.Mutex
}

// New creates a new account service
func New(ctx context.Context, store *store.Store, logger *log.Logger) *WebService {
	return &WebService{
		ctx:    ctx,
		store:  store,
		logger: logger,
	}
}
