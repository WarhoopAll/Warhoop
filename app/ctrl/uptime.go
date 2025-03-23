package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"warhoop/app/config"
)

func (h *Handler) GetUptime(ctx *fiber.Ctx) error {
	cfg := config.Get()

	status, err := h.services.Auth.GetUptime(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	status.Flag = cfg.RealmFlag
	status.Rate = cfg.RealmRate
	status.Realmlis = cfg.RealmRealmlist

	online, err := h.services.Char.GetOnlineCount(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	status.CharOnline = online

	return Response(ctx, MsgSuccess, status)
}
