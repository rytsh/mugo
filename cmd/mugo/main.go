package main

import (
	"context"
	"errors"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/rytsh/mugo/cmd/mugo/args"
	"github.com/worldline-go/logz"
)

var (
	version = "v0.0.0"
	commit  = "0"
	date    = "0"
)

func main() {
	logz.InitializeLog(logz.WithCaller(false))

	appInfo := args.AppInfo{
		Version:     version,
		BuildCommit: commit,
		BuildDate:   date,
	}

	var err error
	defer func() {
		// recover from panic if one occured to prevent os.Exit
		if r := recover(); r != nil {
			log.Panic().Msgf("%v", r)
		}

		if err != nil {
			os.Exit(1)
		}
	}()

	if err = args.Execute(context.Background(), appInfo); err != nil {
		if !errors.Is(err, args.ErrShutdown) {
			log.Error().Err(err).Msg("failed to execute command")
		}
	}
}
