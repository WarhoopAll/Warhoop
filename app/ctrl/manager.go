package ctrl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"warhoop/app/config"
	"warhoop/app/svc"
)

var validate = validator.New()
var cfg = config.Get()

type Handler struct {
	ctx      context.Context
	services *svc.Manager
}

func NewHandler(ctx context.Context, svcs *svc.Manager) *Handler {
	validate = validator.New()
	return &Handler{
		ctx:      ctx,
		services: svcs,
	}
}
