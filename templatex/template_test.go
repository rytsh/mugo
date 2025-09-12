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
