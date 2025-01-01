package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (ctr *Handler) GetCharByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("param")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	entry, err := ctr.services.Char.GetByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}

func (ctr *Handler) GetCharByName(ctx *fiber.Ctx) error {
	name := ctx.Params("param")
	entry, err := ctr.services.Char.GetByName(ctx.Context(), name)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}

func (ctr *Handler) GetCharTop10Kill(ctx *fiber.Ctx) error {
	entry, err := ctr.services.Char.GetTop10Kill(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}
