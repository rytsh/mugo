package main

import (
	"context"
	"errors"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/rytsh/mugo/cmd/mugo/args"
	"github.com/worldline-go/initializer"
	"github.com/worldline-go/logz"
)

var (
	version = "v0.0.0"
	commit  = "0"
	date    = "0"
)

func main() {
	initializer.Init(
		run,
		initializer.WithInitLog(false),
		initializer.WithOptionsLogz(logz.WithCaller(false)),
	)

}

func run(ctx context.Context, wg *sync.WaitGroup) error {
	appInfo := args.AppInfo{
		Version:     version,
		BuildCommit: commit,
		BuildDate:   date,
	}

	if err := args.Execute(context.Background(), appInfo); err != nil {
		if !errors.Is(err, args.ErrShutdown) {
			log.Error().Err(err).Msg("failed to execute command")
		}
	}

	return nil
}
