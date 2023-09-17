package app

import "github.com/urfave/cli/v2"

const (
	CategorySettings = "Settings"
	CategorySecrets  = "Secrets"
)

var (
	FlagHTTPPort = &cli.StringFlag{
		Name:     "http-port",
		Usage:    "HTTP server port",
		EnvVars:  []string{"HTTP_PORT"},
		Category: CategorySettings,
		Value:    "8080",
	}
	_ cli.Flag = FlagHTTPPort

	FlagPostgres = &cli.StringFlag{
		Name:     "postgres-connection-url",
		Usage:    "PostgreSQL connection URL",
		EnvVars:  []string{"POSTGRES_CONNECTION_URL"},
		Category: CategorySecrets,
		Required: true,
	}
	_ cli.Flag = FlagPostgres
)
