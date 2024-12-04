package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"grimoire/app/model"
	"strconv"
)

func (ctr *AccountHandler) News(ctx *fiber.Ctx) error {
	entry, err := ctr.services.Sait.GetNewsSlice(ctx.Context())
	if err != nil {
		return err
	}
	return Response(ctx, MsgSuccess, entry)
}

func (ctr *AccountHandler) NewsByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	entry := &model.News{
		ID: idInt,
	}

	entry, err = ctr.services.Sait.GetNewsByID(ctx.Context(), entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}
