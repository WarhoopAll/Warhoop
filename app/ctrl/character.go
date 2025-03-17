package ctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

func (ctr *Handler) GetCharByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("param")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	entry, err := ctr.services.Char.GetByID(ctx.Context(), id)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}

func (ctr *Handler) GetCharByName(ctx *fiber.Ctx) error {
	name := ctx.Params("param")

	cacheKey := "characters:" + name

	var cachedResponse fiber.Map
	err := ctr.cache.Get(ctx.Context(), cacheKey, &cachedResponse)
	if err == nil {
		log.Println("Returning data from cache")
		return ctx.JSON(cachedResponse)
	} else if err != redis.Nil {
		log.Printf("Failed to get data from Redis: %v", err)
	}

	entry, err := ctr.services.Char.GetByName(ctx.Context(), name)
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	response := fiber.Map{
		"data":   entry,
		"status": "success",
	}

	err = ctr.cache.Set(ctx.Context(), cacheKey, response, time.Second)
	if err != nil {
		log.Printf("Failed to save data to Redis: %v", err)
	} else {
		log.Println("Data saved to Redis")
	}

	return ctx.JSON(response)
}

func (ctr *Handler) GetCharTop10Kill(ctx *fiber.Ctx) error {
	entry, err := ctr.services.Char.GetTop10Kill(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}

func (ctr *Handler) GetOnlineSlice(ctx *fiber.Ctx) error {
	entry, err := ctr.services.Char.GetOnlineSlice(ctx.Context())
	if err != nil {
		return ErrResponse(ctx, MsgInternal)
	}

	return Response(ctx, MsgSuccess, entry)
}
