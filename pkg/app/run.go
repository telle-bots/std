package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"github.com/urfave/cli/v2"
)

func Run(app *cli.App) {
	if app.Usage == "" {
		app.Usage = "CLI application"
	}

	if app.Version == "" {
		updateVersionInfo(app)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		cancel()
	}()

	if err := app.RunContext(ctx, os.Args); err != nil {
		if errors.Is(err, context.Canceled) {
			return
		}
		_, _ = fmt.Fprintf(os.Stderr, "FATAL: %s\n", err)
		os.Exit(1)
	}
}

func updateVersionInfo(app *cli.App) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		app.Version = "unknown"
		return
	}

	var (
		vcsRevision string
		vcsTime     time.Time
		vcsModified bool
	)

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		case "vcs.modified":
			vcsModified, _ = strconv.ParseBool(setting.Value)
		}
	}

	app.Version = fmt.Sprintf("%s, build with %s", info.Main.Version, info.GoVersion)

	if vcsRevision != "" {
		app.Version += ", revision " + vcsRevision
	}
	if !vcsTime.IsZero() {
		app.Version += ", at " + vcsTime.Local().String()
	}
	if vcsModified {
		app.Version += ", modified"
	}
}
