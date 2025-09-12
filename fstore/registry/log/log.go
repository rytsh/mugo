package log

import (
	"log/slog"

	"github.com/rytsh/mugo/fstore"
)

func init() {
	fstore.AddStruct("log", Log{})
}

type Log struct{}

func (Log) Debug(msg string, args ...any) any {
	slog.Debug(msg, args...)

	return args
}

func (Log) Info(msg string, args ...any) any {
	slog.Info(msg, args...)

	return args
}

func (Log) Warn(msg string, args ...any) any {
	slog.Warn(msg, args...)

	return args
}

func (Log) Error(msg string, args ...any) any {
	slog.Error(msg, args...)

	return args
}
