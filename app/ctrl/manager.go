package ctrl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"warhoop/app/cache"
	"warhoop/app/svc"
)

var validate = validator.New()

type Handler struct {
	ctx      context.Context
	services *svc.Manager
	cache    cache.Cache
}

func NewHandler(ctx context.Context, svcs *svc.Manager, cache cache.Cache) *Handler {
	validate = validator.New()
	return &Handler{
		ctx:      ctx,
		services: svcs,
		cache:    cache,
	}
}
