package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"warhoop/app/log"
	"warhoop/app/model/nexus"
)

func (ctr *Handler) NewsSlice(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 2)
	offset := ctx.QueryInt("offset", 0)

	entry, total, err := ctr.services.Web.GetNewsSlice(ctx.Context(), limit, offset)
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

func (ctr *Handler) NewsGetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	entry, err := ctr.services.Web.GetNewsByID(ctx.Context(), idInt)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}

func (ctr *Handler) CreateNews(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	access, err := ctr.services.Auth.AccessByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	if access.SecurityLevel < 3 {
		log.Get().Warn("Forbidden access attempt",
			log.Int("account_id", id))
		return ErrResponse(ctx, MsgForbidden)
	}

	entry := &nexus.News{}

	err = ctx.BodyParser(&entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	res, err := ctr.services.Web.CreateNews(ctx.Context(), id, entry)
	if err != nil {
		return err
	}
	return Response(ctx, MsgSuccess, res)
}

func (ctr *Handler) UpdateNews(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	entry := &nexus.News{}
	err := ParseAndValidate(ctx, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	access, err := ctr.services.Auth.AccessByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	if access.SecurityLevel < 3 {
		log.Get().Warn("Forbidden access attempt",
			log.Int("account_id", id))
		return ErrResponse(ctx, MsgForbidden)
	}

	res, err := ctr.services.Web.UpdateNews(ctx.Context(), entry)
	if err != nil {
		return err
	}
	return Response(ctx, MsgSuccess, res)
}

func (ctr *Handler) DeleteNews(ctx *fiber.Ctx) error {
	idAcc, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	id := ctx.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	access, err := ctr.services.Auth.AccessByID(ctx.Context(), idAcc)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	if access.SecurityLevel < 2 {
		log.Get().Warn("Forbidden access attempt",
			log.Int("account_id", idAcc))
		return ErrResponse(ctx, MsgForbidden)
	}

	err = ctr.services.Web.DeleteNews(ctx.Context(), idInt)
	if err != nil {
		return err
	}
	return Response(ctx, MsgSuccess, nil)
}
