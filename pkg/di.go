package pkg

import (
	"context"

	"github.com/mymmrac/mdi"
	"github.com/urfave/cli/v2"

	"github.com/telle-bots/std/pkg/http"
	"github.com/telle-bots/std/pkg/id"
	"github.com/telle-bots/std/pkg/postgres"
	"github.com/telle-bots/std/pkg/runner"
)

// InitDI provides DI with default provides
func InitDI(app *cli.Context) *mdi.DI {
	ctx, cancel := context.WithCancel(app.Context)
	return mdi.New().
		MustProvide(app).
		MustProvide(func() context.Context { return ctx }).
		MustProvide(func() context.CancelFunc { return cancel }).
		MustProvide(id.NewGenerator).
		MustProvide(runner.NewRunner).
		MustProvide(postgres.DefaultConfig).
		MustProvide(postgres.NewConnector).
		MustProvide(http.DefaultConfig).
		MustProvide(http.DefaultFiberConfig).
		MustProvide(http.Server).
		MustProvide(http.Router)
}
