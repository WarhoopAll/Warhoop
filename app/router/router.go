package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"io"
	"warhoop/app/config"
	"warhoop/app/ctrl"
	"warhoop/app/ctxs"
	"warhoop/app/log"
	"warhoop/app/mw"
)

func SetupRoutes(app *fiber.App, h *ctrl.Handler) {
	api := app.Group("/", logger.New(logger.Config{
		TimeFormat:    "2006-01-02 15:04:05",
		Format:        "${latency}",
		DisableColors: true,
		Output:        io.Discard,
		Done: func(c *fiber.Ctx, logString []byte) {
			if (c.Response().StatusCode() >= 200) && (c.Response().StatusCode() < 300) {
				log.Get().Info(fmt.Sprintf("response ok: %d", c.Response().StatusCode()),
					log.String("method", string(c.Request().Header.Method())),
					log.Int("status", c.Response().StatusCode()),
					log.String("path", string(c.Request().RequestURI())),
					log.String("duration", string(logString)),
					log.Any("ips", c.IPs()),
				)
			}
			if (c.Response().StatusCode() >= 300) && (c.Response().StatusCode() < 600) {
				log.Get().Warn(fmt.Sprintf("response error: %d", c.Response().StatusCode()),
					log.String("err", string(c.Response().Body())),
					log.String("method", string(c.Request().Header.Method())),
					log.Int("status", c.Response().StatusCode()),
					log.String("path", string(c.Request().RequestURI())),
					log.String("duration", string(logString)),
					log.Any("ips", c.IPs()),
				)
			}
		},
	}))

	api.Use(ctxs.Shared)

	app.Static("/favicon.ico", "templates/images/favicon.ico")
	app.Static("/", config.Get().Service.TemplateStatic)

	char := api.Group("/character")
	char.Get("/:param", func(ctx *fiber.Ctx) error {
		param := ctx.Params("param")
		if isNumeric(param) {
			return h.GetCharByID(ctx)
		}
		return h.GetCharByName(ctx)
	})

	char.Get("/pvp", h.GetCharTop10Kill)

	api.Get("/status", h.GetUptime)
	api.Get("/online", h.GetOnlineSlice)

	auth := api.Group("/auth")
	auth.Post("/signup", h.SignUp)
	auth.Post("/signin", h.SignIn)
	auth.Get("/logout", h.Logout)

	news := api.Group("/news")
	news.Get("/", h.NewsSlice)
	news.Get("/:id", h.NewsGetByID)

	api.Use(mw.Auth)

	news.Post("/", h.CreateNews)
	news.Patch("/", h.UpdateNews)
	news.Delete("/:id", h.DeleteNews)

	auth.Get("/session", h.Session)

	news.Post("/comment", h.CreateComment)
	news.Delete("/comment/:id", h.DeleteComment)
	news.Patch("/comment", h.UpdateComment)

	reaction := api.Group("/reaction")
	reaction.Post("/", h.ToggleReaction)

	prof := api.Group("/profile")
	prof.Post("/avatar", h.UpdateAvatar)

	report := api.Group("/report")
	report.Post("/", h.CreateReport)
	report.Get("/:id", h.GetReportByID)
	report.Get("/", h.GetReports)
	report.Delete("/:id", h.DeleteReport)
}
