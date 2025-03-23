package mw

import (
	"github.com/gofiber/fiber/v2"
	"warhoop/app/config"
	"warhoop/app/ctrl"
	"warhoop/app/svc/web"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Cookies(config.Get().CookieName)
	if token == "" {
		return ctrl.ErrResponse(ctx, ctrl.MsgUnauthorized)
	}

	tokenInfo, err := web.TokenVerify(token)
	if err != nil {
		return ctrl.ErrResponse(ctx, ctrl.MsgUnauthorized)
	}

	ctx.Locals("token", token)
	ctx.Locals("id", tokenInfo.ID)
	ctx.Locals("exp", tokenInfo.ExpiresAt)

	return ctx.Next()
}
