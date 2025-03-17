package ctrl

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func (ctr *Handler) GetArmoryCharactersSlice(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 8)
	offset := ctx.QueryInt("offset", 0)

	cacheKey := fmt.Sprintf("armory_characters:%d:%d", limit, offset)

	var cachedResponse fiber.Map
	err := ctr.cache.Get(ctx.Context(), cacheKey, &cachedResponse)
	if err == nil {
		log.Println("Returning data from cache")
		return ctx.JSON(cachedResponse)
	} else if err != redis.Nil {
		log.Printf("Failed to get data from Redis: %v", err)
	}

	entry, total, err := ctr.services.Char.GetArmoryCharactersSlice(ctx.Context(), limit, offset)
	if err != nil {
		log.Printf("Failed to get data from database: %v", err)
		return ErrResponse(ctx, MsgInternal)
	}

	response := fiber.Map{
		"status":  "success",
		"message": "Characters retrieved successfully",
		"data":    entry,
		"total":   total,
	}

	err = ctr.cache.Set(ctx.Context(), cacheKey, response, time.Hour)
	if err != nil {
		log.Printf("Failed to save data to Redis: %v", err)
	} else {
		log.Println("Data saved to Redis")
	}

	return ctx.JSON(response)
}
