package web

import (
	"github.com/gofiber/fiber/v2"
	"grimoire/app/config"
	"grimoire/app/log"
	"time"
)

var cfg = config.Get()

func CreateCookie(value string) *fiber.Cookie {
	log.Get().Logger.Debug("Cookie config",
		"name", cfg.Cookie.Name,
		"domain", cfg.Cookie.Domain,
		"httpOnly", cfg.Cookie.HTTPOnly,
		"secure", cfg.Cookie.Secure,
		"sameSite", cfg.Cookie.SameSite,
		"duration", cfg.Cookie.AccessDuration,
	)

	cookie := &fiber.Cookie{
		Name:     cfg.Cookie.Name,
		Value:    value,
		Domain:   cfg.Cookie.Domain,
		HTTPOnly: cfg.Cookie.HTTPOnly,
		Secure:   cfg.Cookie.Secure,
		SameSite: cfg.Cookie.SameSite,
		Expires:  time.Now().Add(cfg.Cookie.AccessDuration),
	}
	log.Get().Logger.Debug("Cookie created",
		"name", cookie.Name,
		"value", cookie.Value,
		"expires", cookie.Expires,
		"secure", cookie.Secure,
		"sameSite", cookie.SameSite,
	)

	return cookie
}

func DeleteCookie() *fiber.Cookie {
	cookie := &fiber.Cookie{
		Name:     cfg.Cookie.Name,
		Value:    "",
		Domain:   cfg.Cookie.Domain,
		HTTPOnly: true,
		Secure:   true,
		SameSite: cfg.Cookie.SameSite,
		Expires:  time.Now().Add(-1 * time.Hour),
	}
	log.Get().Logger.Info("Cookie deleted",
		"name", cookie.Name,
		"expires", cookie.Expires,
	)
	return cookie
}

func (svc *WebService) CreateCookie(token string) *fiber.Cookie {
	return CreateCookie(token)
}
