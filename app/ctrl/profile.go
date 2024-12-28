package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"warhoop/app/model"
)

func (ctr *AccountHandler) UpdateAvatar(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}
	entry := &model.Profile{}
	err := ParseAndValidate(ctx, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	entry.AccountID = id
	res, err := ctr.services.Sait.UpdateAvatar(ctx.Context(), entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	if res == nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgAvatarUpdate, res)
}
