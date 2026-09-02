package log

import "github.com/rytsh/mugo/fstore"

func init() {
	fstore.AddStructWithOptions(func(o fstore.Option) (string, Log) {
		return "log", New(o.Log)
	})
}

type Log struct {
	log fstore.Adapter
}

func New(log fstore.Adapter) Log {
	if log == nil {
		log = fstore.Noop{}
	}

	return Log{log: log}
}

func (l Log) Debug(msg string, args ...any) any {
	l.log.Debug(msg, args...)

	return args
}

func (l Log) Info(msg string, args ...any) any {
	l.log.Info(msg, args...)

	return args
}

func (l Log) Warn(msg string, args ...any) any {
	l.log.Warn(msg, args...)

	return args
}

func (l Log) Error(msg string, args ...any) any {
	l.log.Error(msg, args...)

	return args
}
