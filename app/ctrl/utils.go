package ctrl

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"net"
	"strconv"
	"strings"
	"time"
)

func FingerPrint(ctx *fiber.Ctx, id int) string {
	hash := sha256.New()

	idStr := strconv.Itoa(id)

	ips := strings.Join(ctx.IPs(), ",")

	hash.Write([]byte(idStr))
	hash.Write([]byte(ips))
	hash.Write([]byte(ctx.Get("User-Agent")))
	return hex.EncodeToString(hash.Sum(nil))
}

func ParseAndValidate(ctx *fiber.Ctx, entry interface{}) error {
	err := ctx.BodyParser(entry)
	if err != nil {
		return err
	}
	err = validate.Struct(entry)
	if err != nil {
		return err
	}
	return nil
}

func CheckOnline(host, port string) bool {
	address := net.JoinHostPort(host, port)

	timeout := 1 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()

	return true
}
