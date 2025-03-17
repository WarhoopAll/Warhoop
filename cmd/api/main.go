package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"warhoop/app/cache"
	"warhoop/app/config"
	"warhoop/app/ctrl"
	"warhoop/app/log"
	"warhoop/app/mw"
	"warhoop/app/router"
	"warhoop/app/store"
	"warhoop/app/svc"
	"warhoop/app/utils"
)

func main() {
	logger := log.Get()

	if err := run(logger); err != nil {
		logger.Error(err.Error())
		panic(err.Error())
	}
}

func run(logger *log.Logger) error {
	ctx := context.Background()

	// Init welcome page
	err := utils.LoadAndGenerateHTML(config.Get().Service.GitInfo)
	if err != nil {
		logger.Error("error generate welcome page",
			log.String("err", err.Error()))
		return nil
	}

	// Init repository store (with mysql inside)
	store, err := store.NewBun(logger)
	if err != nil {
		logger.Error("store.NewBun failed",
			log.String("err", err.Error()),
		)
		return err
	}
	redisCache := cache.NewRedisCache()

	// Init service manager
	serviceManager, err := svc.NewManager(ctx, store, logger)
	if err != nil {
		logger.Error("manager.New failed",
			log.String("err", err.Error()),
		)
		return err
	}

	hAccount := ctrl.NewHandler(ctx, serviceManager, redisCache)

	app := fiber.New()
	app.Get("/metrics", monitor.New())
	app.Use(mw.SetupCors())
	app.Use(healthcheck.New())

	// Routers
	router.SetupRoutes(app, hAccount)

	// Starting api server
	logger.Error("fatal error",
		log.String("err", app.Listen(config.Get().Service.ApiAddrPort).Error()),
	)
	return nil
}
