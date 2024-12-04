package web

import (
	"github.com/gofiber/fiber/v2"
	"grimoire/app/config"
	"grimoire/app/log"
	"time"
)

var cfg = config.Get()

func CreateCookie(value string) *fiber.Cookie {
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
	log.Get().Logger.Debug("Cookie deleted",
		"name", cookie.Name,
		"expires", cookie.Expires,
	)
	return cookie
}

func (svc *WebService) CreateCookie(token string) *fiber.Cookie {
	return CreateCookie(token)
}
