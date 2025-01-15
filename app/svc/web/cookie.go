package web

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"warhoop/app/config"
	"warhoop/app/log"
)

var cfg = config.Get()
var logger = log.Get()

func CreateCookie(value string) *fiber.Cookie {
	logger.Debug("utils.Cookie.config",
		log.Object("name", cfg.Cookie),
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
	logger.Debug("utils.Cookie.CreateCookie",
		log.Object("cookie", cookie),
	)

	return cookie
}

func DeleteCookie() *fiber.Cookie {
	cookie := &fiber.Cookie{
		Name:     cfg.Cookie.Name,
		Value:    "",
		Domain:   cfg.Cookie.Domain,
		HTTPOnly: cfg.Cookie.HTTPOnly,
		Secure:   cfg.Cookie.Secure,
		SameSite: cfg.Cookie.SameSite,
		Expires:  time.Now().Add(-1 * time.Hour),
	}

	logger.Debug("utils.Cookie.DeleteCookie",
		log.Object("cookie", cookie),
	)
	return cookie
}

func (svc *WebService) CreateCookie(token string) *fiber.Cookie {
	return CreateCookie(token)
}
