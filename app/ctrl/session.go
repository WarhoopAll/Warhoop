package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"warhoop/app/config"
	"warhoop/app/ctxs"
	"warhoop/app/model/nexus"
	"warhoop/app/svc/web"
	"warhoop/app/utils"
)

func (ctr *Handler) Session(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals("id").(int)
	if !ok {
		return ErrResponse(ctx, MsgUnauthorized)
	}

	c, ok := ctx.Locals("s").(*ctxs.Ctx)
	if !ok {
		return ErrResponse(ctx, MsgInternal)
	}

	oldToken := ctx.Cookies(config.Get().CookieName)

	session := &nexus.Session{
		Token:     oldToken,
		AccountID: id,
		IPs:       c.IPs,
		UpdatedAt: time.Now(),
		ExpiredAt: time.Now().Add(config.Get().CookieAccessDuration),
	}

	newToken, err := web.GenerateTokenAccess(id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	fnd, err := ctr.services.Auth.GetByID(ctx.Context(), id)
	if err != nil {
		return err
	}

	finger := FingerPrint(ctx, fnd.ID)

	session.Finger = finger
	err = ctr.services.Web.UpdateOrCreateSession(ctx.Context(), session, newToken)
	if err != nil {
		if err == utils.ErrDataBase {
			return ErrResponse(ctx, MsgUnauthorized)
		}
		return ErrResponse(ctx, MsgInternal)
	}

	ctx.Cookie(web.CreateCookie(newToken))

	return Response(ctx, MsgSuccess, fnd)
}
