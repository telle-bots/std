package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"

	"github.com/telle-bots/std/pkg/runner"
)

// Server provides Fiber server
func Server(ctx context.Context, cfg Config, fiberCfg fiber.Config, run runner.Runner) *fiber.App {
	app := fiber.New(fiberCfg)
	app.Use(middlewareContext(cfg.NoContextCancel))
	app.Group("/debug/pprof", pprof.New())

	run.Add(&listenCmd{
		ctx: ctx,
		cfg: cfg,
		app: app,
	})

	return app
}
