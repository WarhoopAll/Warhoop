package ctrl

import (
	"context"
	"github.com/go-playground/validator/v10"
	"grimoire/app/config"
	"grimoire/app/svc"
)

var validate = validator.New()
var cfg = config.Get()

type AccountHandler struct {
	ctx      context.Context
	services *svc.Manager
}

func NewAccount(ctx context.Context, svcs *svc.Manager) *AccountHandler {
	validate = validator.New()
	return &AccountHandler{
		ctx:      ctx,
		services: svcs,
	}
}
