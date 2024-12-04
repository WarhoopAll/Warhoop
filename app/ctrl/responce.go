package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"grimoire/app/log"
)

func JSONResponse(ctx *fiber.Ctx, status string, message string, data interface{}) error {
	httpStatus := GetStatusByMessage(message)

	return ctx.Status(httpStatus).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func GetStatusByMessage(msg string) int {
	status, exists := ErrorMapping[msg]
	if !exists {
		log.Get().Logger.Warn("Message not found in ErrorMapping",
			"message", msg,
		)
		return fiber.StatusInternalServerError
	}

	log.Get().Logger.Debug("Message mapped to status",
		"message", msg,
		"status", status,
	)
	return status
}

func ErrResponse(ctx *fiber.Ctx, message string) error {
	return JSONResponse(ctx, "error", message, nil)
}

func Response(ctx *fiber.Ctx, message string, data interface{}) error {
	return JSONResponse(ctx, "success", message, data)
}
