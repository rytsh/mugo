package templatex

import (
	"bytes"
	"testing"
)

func TestTemplate_Execute(t *testing.T) {
	type args struct {
		v       any
		content string
	}
	tests := []struct {
		name    string
		args    args
		opts    []OptionTemplate
		want    string
		wantErr bool
	}{
		{
			name: "test template",
			args: args{
				v:       map[string]any{"name": "test"},
				content: `{{ .name }}`,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "custom func",
			args: args{
				v:       map[string]any{"name": "x"},
				content: `{{ custom .name }}`,
			},
			opts: []OptionTemplate{WithAddFunc("custom", func(x string) string {
				return x + "custom"
			})},
			want:    "xcustom",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			templateEngine := New(tt.opts...)

			var buf bytes.Buffer
			err := templateEngine.Execute(WithIO(&buf), WithData(tt.args.v), WithContent(tt.args.content))
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if buf.String() != tt.want {
				t.Errorf("Template.Execute() = %s, want %s", buf.String(), tt.want)
			}
		})
	}
}

func TestTemplate_ExecuteWithExecFunc(t *testing.T) {
	t.Run("single exec func", func(t *testing.T) {
		tpl := New()

		var buf bytes.Buffer
		err := tpl.Execute(
			WithIO(&buf),
			WithData(map[string]any{"a": 2, "b": 3}),
			WithContent(`{{ add .a .b }}`),
			WithExecFunc("add", func(a, b int) int { return a + b }),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := buf.String(); got != "5" {
			t.Errorf("got %q, want %q", got, "5")
		}
	})

	t.Run("exec func map", func(t *testing.T) {
		tpl := New()

		var buf bytes.Buffer
		err := tpl.Execute(
			WithIO(&buf),
			WithData(map[string]any{"a": 10, "b": 4}),
			WithContent(`{{ add .a .b }} {{ sub .a .b }}`),
			WithExecFuncMap(map[string]any{
				"add": func(a, b int) int { return a + b },
				"sub": func(a, b int) int { return a - b },
			}),
		)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got := buf.String(); got != "14 6" {
			t.Errorf("got %q, want %q", got, "14 6")
		}
	})

	t.Run("exec func does not leak to next execution", func(t *testing.T) {
		tpl := New()

		// First execution with the exec func
		var buf1 bytes.Buffer
		err := tpl.Execute(
			WithIO(&buf1),
			WithData(map[string]any{"x": "hello"}),
			WithContent(`{{ upper .x }}`),
			WithExecFunc("upper", func(s string) string {
				return "UPPER:" + s
			}),
		)
		if err != nil {
			t.Fatalf("first execute: unexpected error: %v", err)
		}
		if got := buf1.String(); got != "UPPER:hello" {
			t.Errorf("first execute: got %q, want %q", got, "UPPER:hello")
		}

		// Second execution without the exec func should fail
		var buf2 bytes.Buffer
		err = tpl.Execute(
			WithIO(&buf2),
			WithData(map[string]any{"x": "hello"}),
			WithContent(`{{ upper .x }}`),
		)
		if err == nil {
			t.Fatal("second execute: expected error for undefined function, got nil")
		}
	})

	t.Run("exec func overrides base func for that execution only", func(t *testing.T) {
		tpl := New(WithAddFunc("greet", func(s string) string {
			return "hello " + s
		}))

		// Override greet for this execution
		var buf1 bytes.Buffer
		err := tpl.Execute(
			WithIO(&buf1),
			WithData(map[string]any{"name": "world"}),
			WithContent(`{{ greet .name }}`),
			WithExecFunc("greet", func(s string) string {
				return "hi " + s
			}),
		)
		if err != nil {
			t.Fatalf("overridden execute: unexpected error: %v", err)
		}
		if got := buf1.String(); got != "hi world" {
			t.Errorf("overridden execute: got %q, want %q", got, "hi world")
		}

		// Next execution should use original greet
		var buf2 bytes.Buffer
		err = tpl.Execute(
			WithIO(&buf2),
			WithData(map[string]any{"name": "world"}),
			WithContent(`{{ greet .name }}`),
		)
		if err != nil {
			t.Fatalf("original execute: unexpected error: %v", err)
		}
		if got := buf2.String(); got != "hello world" {
			t.Errorf("original execute: got %q, want %q", got, "hello world")
		}
	})
}
