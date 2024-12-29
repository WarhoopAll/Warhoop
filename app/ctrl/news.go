package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"grimoire/app/log"
	"grimoire/app/model"
	"strconv"
)

func (ctr *AccountHandler) NewsSlice(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 2)
	offset := ctx.QueryInt("offset", 0)

	entry, total, err := ctr.services.Sait.GetNewsSlice(ctx.Context(), limit, offset)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "News retrieved successfully",
		"data":    entry,
		"total":   total,
	})
}

func (ctr *AccountHandler) NewsGetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	entry, err := ctr.services.Sait.GetNewsByID(ctx.Context(), idInt)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}

func (ctr *AccountHandler) CreateNews(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	access, err := ctr.services.Account.AccessByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	if access.SecurityLevel < 3 {
		log.Get().Warn("Forbidden access attempt",
			log.Int("account_id", id))
		return ErrResponse(ctx, MsgForbidden)
	}

	entry := &model.News{
		Author: id,
	}
	if err := ctx.BodyParser(&entry); err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	res, err := ctr.services.Sait.CreateNews(ctx.Context(), entry)
	if err != nil {
		return err
	}
	return Response(ctx, MsgSuccess, res)
}

func (ctr *AccountHandler) DeleteNews(ctx *fiber.Ctx) error {
	idAcc, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	id := ctx.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	access, err := ctr.services.Account.AccessByID(ctx.Context(), idAcc)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	if access.SecurityLevel < 2 {
		log.Get().Warn("Forbidden access attempt",
			log.Int("account_id", idAcc))
		return ErrResponse(ctx, MsgForbidden)
	}

	err = ctr.services.Sait.DeleteNews(ctx.Context(), idInt)
	if err != nil {
		return err
	}
	return Response(ctx, MsgSuccess, nil)
}
