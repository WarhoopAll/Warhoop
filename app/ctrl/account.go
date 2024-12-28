package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"warhoop/app/ctxs"
	"warhoop/app/model"
	"warhoop/app/svc/web"
	"time"
)

func (ctr *AccountHandler) SignIn(ctx *fiber.Ctx) error {
	c, ok := ctx.Locals("s").(*ctxs.Ctx)
	if !ok {
		return ErrResponse(ctx, MsgInternal)
	}

	entry := &model.Account{}

	err := ParseAndValidate(ctx, entry)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	res, err := ctr.services.Account.SignIn(ctx.Context(), entry)
	if err != nil {
		return ErrResponse(ctx, err.Error())
	}

	finger := FingerPrint(ctx, res.ID)

	session := &model.Session{
		AccountID: res.ID,
		IPs:       c.IPs,
		Finger:    finger,
		UpdatedAt: time.Now(),
		ExpiredAt: time.Now().Add(cfg.Cookie.AccessDuration),
	}

	token, err := ctr.services.Sait.GenerateAccessToken(res.ID)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	session.Token = token

	_, err = ctr.services.Sait.HandleSession(ctx.Context(), session)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	cookie := ctr.services.Sait.CreateCookie(token)
	ctx.Cookie(cookie)

	return Response(ctx, MsgSignIn, res)
}

func (ctr *AccountHandler) SignUp(ctx *fiber.Ctx) error {
	c, ok := ctx.Locals("s").(*ctxs.Ctx)
	if !ok {
		return ErrResponse(ctx, MsgInternal)
	}

	entry := &model.Account{}

	err := ParseAndValidate(ctx, entry)
	if err != nil {
		return ErrResponse(ctx, err.Error())
	}

	res, err := ctr.services.Account.SignUp(ctx.Context(), entry)
	if err != nil {
		return ErrResponse(ctx, err.Error())
	}

	finger := FingerPrint(ctx, res.ID)

	token, err := web.GenerateTokenAccess(res.ID)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	session := &model.Session{
		AccountID: res.ID,
		Token:     token,
		IPs:       c.IPs,
		Finger:    finger,
		UpdatedAt: time.Now(),
		ExpiredAt: time.Now().Add(cfg.Cookie.AccessDuration),
	}

	err = ctr.services.Sait.CreateSession(ctx.Context(), session)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	cookie := ctr.services.Sait.CreateCookie(token)
	ctx.Cookie(cookie)
	return Response(ctx, MsgSignUp, res)
}

func (ctr *AccountHandler) Logout(ctx *fiber.Ctx) error {
	token := ctx.Cookies(cfg.Cookie.Name)

	cookie, err := ctr.services.Sait.DeleteSession(ctx.Context(), token)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	ctx.Cookie(cookie)

	return Response(ctx, MsgLogout, nil)
}
