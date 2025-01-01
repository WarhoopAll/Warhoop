package ctrl

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUptime(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	status, err := h.services.Auth.GetUptimeByID(ctx.Context(), id)
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
