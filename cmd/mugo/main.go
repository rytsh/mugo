package main

import (
	"context"
	"errors"

	"github.com/rakunlabs/into"
	"github.com/rakunlabs/logi"
	"github.com/rs/zerolog/log"
	"github.com/rytsh/mugo/cmd/mugo/args"
)

var (
	version = "v0.0.0"
	commit  = "0"
	date    = "0"
)

func main() {
	into.Init(
		run,
		into.WithLogger(logi.InitializeLog(logi.WithCaller(false))),
		into.WithStartFn(nil),
		into.WithStopFn(nil),
	)

}

func run(ctx context.Context) error {
	appInfo := args.AppInfo{
		Version:     version,
		BuildCommit: commit,
		BuildDate:   date,
	}

	if err := args.Execute(ctx, appInfo); err != nil {
		if !errors.Is(err, args.ErrShutdown) {
			log.Error().Err(err).Msg("failed to execute command")
		}
	}

	return nil
}
