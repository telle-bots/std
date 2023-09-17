package app

import "github.com/urfave/cli/v2"

const (
	// CategorySettings general settings
	CategorySettings = "Settings"
	// CategorySecrets secrets
	CategorySecrets = "Secrets"
)

var (
	// FlagServiceName service name
	FlagServiceName = &cli.StringFlag{
		Name:     "service-name",
		Usage:    "Service name",
		EnvVars:  []string{"SERVICE_NAME"},
		Category: CategorySettings,
		Required: true,
	}

	// FlagHTTPPort HTTP server port
	FlagHTTPPort = &cli.UintFlag{
		Name:     "http-port",
		Usage:    "HTTP server port",
		EnvVars:  []string{"HTTP_PORT"},
		Category: CategorySettings,
		Value:    8080,
	}

	// FlagPostgres PostgreSQL connection URL
	FlagPostgres = &cli.StringFlag{
		Name:     "postgres-connection-url",
		Usage:    "PostgreSQL connection URL",
		EnvVars:  []string{"POSTGRES_CONNECTION_URL"},
		Category: CategorySecrets,
		Required: true,
	}
)
