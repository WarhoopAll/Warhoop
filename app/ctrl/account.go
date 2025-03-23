package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"warhoop/app/config"
	"warhoop/app/ctxs"
	"warhoop/app/model/auth"
	"warhoop/app/model/nexus"
	"warhoop/app/svc/web"
)

func (ctr *Handler) SignIn(ctx *fiber.Ctx) error {
	c, ok := ctx.Locals("s").(*ctxs.Ctx)
	if !ok {
		return ErrResponse(ctx, MsgInternal)
	}

	entry := &auth.Account{}

	err := ParseAndValidate(ctx, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	res, err := ctr.services.Auth.SignIn(ctx.Context(), entry)
	if err != nil {
		return ErrResponse(ctx, err.Error())
	}

	finger := FingerPrint(ctx, res.ID)

	session := &nexus.Session{
		AccountID: res.ID,
		IPs:       c.IPs,
		Finger:    finger,
		UpdatedAt: time.Now(),
		ExpiredAt: time.Now().Add(config.Get().CookieAccessDuration),
	}

	token, err := ctr.services.Web.GenerateAccessToken(res.ID)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	session.Token = token

	_, err = ctr.services.Web.HandleSession(ctx.Context(), session)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	cookie := ctr.services.Web.CreateCookie(token)
	ctx.Cookie(cookie)

	return Response(ctx, MsgSignIn, res)
}

func (ctr *Handler) SignUp(ctx *fiber.Ctx) error {
	c, ok := ctx.Locals("s").(*ctxs.Ctx)
	if !ok {
		return ErrResponse(ctx, MsgInternal)
	}

	entry := &auth.Account{}

	err := ParseAndValidate(ctx, entry)
	if err != nil {
		return ErrResponse(ctx, err.Error())
	}

	res, err := ctr.services.Auth.SignUp(ctx.Context(), entry)
	if err != nil {
		return ErrResponse(ctx, err.Error())
	}

	finger := FingerPrint(ctx, res.ID)

	token, err := web.GenerateTokenAccess(res.ID)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	session := &nexus.Session{
		AccountID: res.ID,
		Token:     token,
		IPs:       c.IPs,
		Finger:    finger,
		UpdatedAt: time.Now(),
		ExpiredAt: time.Now().Add(config.Get().CookieAccessDuration),
	}

	err = ctr.services.Web.CreateSession(ctx.Context(), session)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	cookie := ctr.services.Web.CreateCookie(token)
	ctx.Cookie(cookie)
	return Response(ctx, MsgSignUp, res)
}

func (ctr *Handler) Logout(ctx *fiber.Ctx) error {
	token := ctx.Cookies(config.Get().CookieName)

	cookie, err := ctr.services.Web.DeleteSession(ctx.Context(), token)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	ctx.Cookie(cookie)

	return Response(ctx, MsgLogout, nil)
}
