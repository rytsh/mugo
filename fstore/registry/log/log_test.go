package log

import "testing"

type recorder struct {
	level string
	msg   string
	args  []any
}

func (r *recorder) Error(msg string, args ...any) { r.record("error", msg, args) }
func (r *recorder) Info(msg string, args ...any)  { r.record("info", msg, args) }
func (r *recorder) Debug(msg string, args ...any) { r.record("debug", msg, args) }
func (r *recorder) Warn(msg string, args ...any)  { r.record("warn", msg, args) }

func (r *recorder) record(level, msg string, args []any) {
	r.level = level
	r.msg = msg
	r.args = args
}

func TestLogUsesAdapter(t *testing.T) {
	recorder := &recorder{}
	logger := New(recorder)

	got := logger.Info("rendered", "name", "mugo")
	if recorder.level != "info" || recorder.msg != "rendered" {
		t.Fatalf("recorded level/message = %q/%q", recorder.level, recorder.msg)
	}
	if len(recorder.args) != 2 || recorder.args[1] != "mugo" {
		t.Fatalf("recorded args = %#v", recorder.args)
	}
	if args, ok := got.([]any); !ok || len(args) != 2 || args[1] != "mugo" {
		t.Fatalf("Info() = %#v", got)
	}
}
