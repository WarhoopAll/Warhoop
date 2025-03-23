package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"warhoop/app/model/auth"
	"warhoop/app/model/nexus"
)

func (ctr *Handler) CreateComment(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	entry := &nexus.Comment{}
	err := ParseAndValidate(ctx, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	res, err := ctr.services.Web.CreateComment(ctx.Context(), id, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, res)
}

func (ctr *Handler) DeleteComment(ctx *fiber.Ctx) error {
	idacc, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	res, err := ctr.services.Auth.GetByID(ctx.Context(), idacc)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	if res.Access == nil {
		res.Access = &auth.Access{SecurityLevel: 0}
	}

	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	comment, err := ctr.services.Web.GetCommentByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgNotFound)
	}

	if comment.Profile == nil || comment.Profile.AccountID != idacc && res.Access.SecurityLevel <= 0 {
		return ErrResponse(ctx, MsgForbidden)
	}

	err = ctr.services.Web.DeleteComment(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, nil)
}

func (ctr *Handler) UpdateComment(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	entry := &nexus.Comment{}
	err := ctx.BodyParser(&entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	res, err := ctr.services.Web.UpdateComment(ctx.Context(), id, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}
	return Response(ctx, MsgSuccess, res)
}
