package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"grimoire/app/config"
	"grimoire/app/ctrl"
	"grimoire/app/log"
	"grimoire/app/mw"
	"grimoire/app/router"
	"grimoire/app/store"
	"grimoire/app/svc"
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

	// Init repository store (with mysql inside)
	store, err := store.NewBun(ctx, logger)
	if err != nil {
		logger.Error("store.NewBun failed",
			log.String("err", err.Error()),
		)
		return err
	}

	// Init service manager
	serviceManager, err := svc.NewManager(ctx, store, logger)
	if err != nil {
		logger.Error("manager.New failed",
			log.String("err", err.Error()),
		)
		return err
	}

	// Init handlers
	hAccount := ctrl.NewAccount(ctx, serviceManager)

	app := fiber.New()
	app.Use(mw.SetupCors())

	// Routers
	router.SetupRoutes(app, hAccount)

	// Starting api server
	logger.Error("fatal error",
		log.String("err", app.Listen(config.Get().Service.ApiAddrPort).Error()),
	)
	return nil
}
