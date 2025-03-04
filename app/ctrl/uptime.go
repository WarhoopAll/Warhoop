package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"warhoop/app/config"
)

func (h *Handler) GetUptime(ctx *fiber.Ctx) error {
	status, err := h.services.Auth.GetUptime(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	status.Flag = config.Get().Realm.Flag
	status.Rate = config.Get().Realm.Rate
	status.Realmlis = config.Get().Realm.Realmlist

	online, err := h.services.Char.GetOnlineCount(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	status.CharOnline = online

	return Response(ctx, MsgSuccess, status)
}
