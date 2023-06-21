package templatex

import (
	"testing"

	"github.com/rytsh/mugo/pkg/templatex/store"
)

func TestTemplate_Execute(t *testing.T) {
	type args struct {
		v       any
		content string
	}
	tests := []struct {
		name    string
		args    args
		opts    []store.Option
		want    string
		wantErr bool
	}{
		{
			name: "test template",
			args: args{
				v:       map[string]interface{}{"name": "test"},
				content: `{{ .name }}`,
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "custom func",
			args: args{
				v:       map[string]interface{}{"name": "x"},
				content: `{{ custom .name }}`,
			},
			opts: []store.Option{store.WithAddFunc("custom", func(x string) string {
				return x + "custom"
			})},
			want:    "xcustom",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			templateEngine := New(tt.opts...)

			got, err := templateEngine.ExecuteBuffer(WithData(tt.args.v), WithContent(tt.args.content))
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != tt.want {
				t.Errorf("Template.Execute() = %s, want %s", got, tt.want)
			}
		})
	}
}
