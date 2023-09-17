package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"

	appStd "github.com/telle-bots/std/pkg/app"
)

// Config represents HTTP server config
type Config struct {
	// Port is an HTTP port
	Port uint
	// NoContextCancel if true prevents context cancellation after the request is closed
	NoContextCancel bool
}

// DefaultConfig provides default HTTP server config
func DefaultConfig(app *cli.Context) Config {
	return Config{
		Port:            app.Uint(appStd.FlagHTTPPort.Name),
		NoContextCancel: false,
	}
}

// DefaultFiberConfig provides default Fiber config
func DefaultFiberConfig(app *cli.Context) fiber.Config {
	return fiber.Config{
		BodyLimit:             1024 * 1024 * 2,
		DisableStartupMessage: true,
		AppName:               app.String(appStd.FlagServiceName.Name),
	}
}
