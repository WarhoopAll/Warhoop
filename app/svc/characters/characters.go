package characters

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/store"
)

type CharService struct {
	ctx    context.Context
	store  *store.Store
	logger *log.Logger
}

// New creates a new account service
func New(ctx context.Context, store *store.Store, logger *log.Logger) *CharService {
	return &CharService{
		ctx:    ctx,
		store:  store,
		logger: logger,
	}
}
