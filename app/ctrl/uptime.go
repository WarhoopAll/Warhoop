package ctrl

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUptime(ctx *fiber.Ctx) error {
	status, err := h.services.Auth.GetUptime(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	online, err := h.services.Char.GetOnlineCount(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	status.CharOnline = online

	return Response(ctx, MsgSuccess, status)
}
