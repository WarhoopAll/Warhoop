package mw

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"warhoop/app/config"
)

var cfg = config.Get()

func SetupCors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     cfg.Cors.AllowOrigins,
		AllowMethods:     cfg.Cors.AllowMethods,
		AllowHeaders:     cfg.Cors.AllowHeaders,
		AllowCredentials: cfg.Cors.AllowCredentials,
	})
}
