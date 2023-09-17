package postgres

import (
	"github.com/urfave/cli/v2"

	appStd "github.com/telle-bots/std/pkg/app"
)

// Config represents PostgreSQL config
type Config struct {
	// ConnectionString of the database
	ConnectionString string
}

// DefaultConfig provides default PostgreSQL config
func DefaultConfig(app *cli.Context) Config {
	return Config{
		ConnectionString: app.String(appStd.FlagPostgres.Name),
	}
}
