package main

import (
	"context"
	"log/slog"

	"github.com/rakunlabs/into"
	"github.com/rakunlabs/logi"
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
		into.WithRunErrFn(func(err error) {
			slog.Error("command failed", slog.String("error", err.Error()))
		}),
	)
}

func run(ctx context.Context) error {
	appInfo := args.AppInfo{
		Version:     version,
		BuildCommit: commit,
		BuildDate:   date,
	}

	return args.Execute(ctx, appInfo)
}
