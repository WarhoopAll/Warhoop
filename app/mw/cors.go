package mw

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"warhoop/app/config"
)

//var cfg = config.Get()

func SetupCors() fiber.Handler {
	cfg := config.Get()
	return cors.New(cors.Config{
		AllowOrigins:     cfg.CorsAllowOrigins,
		AllowMethods:     cfg.CorsAllowMethods,
		AllowHeaders:     cfg.CorsAllowHeaders,
		AllowCredentials: cfg.CorsAllowCredentials,
	})
}
