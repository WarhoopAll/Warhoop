package ctxs

import (
	"github.com/gofiber/fiber/v2"
	"grimoire/app/config"
	"strings"
)

type Ctx struct {
	IPs    []string
	Locale string
}

func Shared(c *fiber.Ctx) error {
	lang := parseLang(c.Get("Accept-Language", ""))
	ips := getIPs(c)

	c.Locals("s", &Ctx{
		Locale: lang,
		IPs:    ips,
	})

	return c.Next()
}

func parseLang(header string) string {
	if header == "" {
		return config.Get().Service.DefaultLocale
	}

	locales := strings.Split(header, ",")
	loc := strings.ToLower(strings.Split(locales[0], ";")[0])

	switch {
	case loc == "ru" || strings.HasPrefix(loc, "ru-"):
		return "ru"
	case loc == "en" || strings.HasPrefix(loc, "en-"):
		return "en"
	default:
		return config.Get().Service.DefaultLocale
	}
}

func getIPs(c *fiber.Ctx) []string {
	ips := c.IPs()
	if len(ips) == 0 {
		return []string{c.IP()}
	}
	return ips
}
