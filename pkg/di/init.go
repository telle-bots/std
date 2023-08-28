package di

import (
	"context"

	"github.com/mymmrac/mdi"
	"github.com/urfave/cli/v2"

	"github.com/telle-bots/std/pkg/id"
)

func Init(app *cli.Context) *mdi.DI {
	ctx, cancel := context.WithCancel(app.Context)
	return mdi.New().
		MustProvide(app).
		MustProvide(func() context.Context { return ctx }).
		MustProvide(func() context.CancelFunc { return cancel }).
		MustProvide(id.NewGenerator)
}
