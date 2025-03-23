package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"warhoop/app/model/nexus"
)

func (h *Handler) CreateReport(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	entry := &nexus.Report{}
	err := ctx.BodyParser(&entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	report, err := h.services.Web.CreateReport(ctx.Context(), id, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, report)
}

func (h *Handler) GetReportByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ErrResponse(ctx, MsgInternal)
	}

	report, err := h.services.Web.GetReportByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgNotFound)
	}

	return Response(ctx, MsgSuccess, report)
}

func (h *Handler) GetReports(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil || limit <= 0 {
		return ErrResponse(ctx, MsgInternal)
	}

	offset, err := strconv.Atoi(ctx.Query("offset", "0"))
	if err != nil || offset < 0 {
		return ErrResponse(ctx, MsgInternal)
	}

	reports, err := h.services.Web.GetReports(ctx.Context(), limit, offset)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, reports)
}

func (h *Handler) DeleteReport(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id <= 0 {
		return ErrResponse(ctx, MsgInternal)
	}

	err = h.services.Web.DeleteReportByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, nil)
}
