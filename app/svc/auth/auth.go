package auth

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/store"
)

type AuthService struct {
	ctx    context.Context
	store  *store.Store
	logger *log.Logger
}

// New creates a new account service
func New(ctx context.Context, store *store.Store, logger *log.Logger) *AuthService {
	return &AuthService{
		ctx:    ctx,
		store:  store,
		logger: logger,
	}
}
