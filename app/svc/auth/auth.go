package auth

import (
	"context"
	"warhoop/app/log"
	"warhoop/app/store"
	"warhoop/app/svc/web"
)

type AuthService struct {
	ctx    context.Context
	store  *store.Store
	logger *log.Logger
	web    *web.WebService
}

// New creates a new account service
func New(ctx context.Context, store *store.Store, logger *log.Logger, web *web.WebService) *AuthService {
	return &AuthService{
		ctx:    ctx,
		store:  store,
		logger: logger,
		web:    web,
	}
}
