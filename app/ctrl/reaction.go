package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"warhoop/app/model/nexus"
)

func (ctr *Handler) ToggleReaction(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	entry := &nexus.Reaction{}
	if err := ctx.BodyParser(&entry); err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	entry.UserID = id

	res, err := ctr.services.Web.ToggleReaction(ctx.Context(), entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, res)
}
