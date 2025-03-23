package web

import (
	"github.com/gofiber/fiber/v2"
	"time"
	"warhoop/app/config"
	"warhoop/app/log"
)

func CreateCookie(value string) *fiber.Cookie {
	cfg := config.Get()

	//logger.Debug("utils.Cookie.config",
	//	log.Object("name", cfg.Cookie),
	//)

	cookie := &fiber.Cookie{
		Name:     cfg.CookieName,
		Value:    value,
		Domain:   cfg.CookieDomain,
		HTTPOnly: cfg.CookieHTTPOnly,
		Secure:   cfg.CookieSecure,
		SameSite: cfg.CookieSameSite,
		Expires:  time.Now().Add(cfg.CookieAccessDuration),
	}
	log.Get().Debug("utils.Cookie.CreateCookie",
		log.Object("cookie", cookie),
	)

	return cookie
}

func DeleteCookie() *fiber.Cookie {
	cfg := config.Get()

	cookie := &fiber.Cookie{
		Name:     cfg.CookieName,
		Value:    "",
		Domain:   cfg.CookieDomain,
		HTTPOnly: cfg.CookieHTTPOnly,
		Secure:   cfg.CookieSecure,
		SameSite: cfg.CookieSameSite,
		Expires:  time.Now().Add(-1 * time.Hour),
	}

	log.Get().Debug("utils.Cookie.DeleteCookie",
		log.Object("cookie", cookie),
	)
	return cookie
}

func (svc *WebService) CreateCookie(token string) *fiber.Cookie {
	return CreateCookie(token)
}
