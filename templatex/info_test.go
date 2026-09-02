package templatex

import (
	"reflect"
	"sort"
	"testing"
)

func TestFuncInfos(t *testing.T) {
	type args struct {
		fns map[string]any
	}
	tests := []struct {
		name string
		args args
		want []Info
	}{
		{
			name: "TestFuncInfos",
			args: args{
				fns: map[string]any{
					"uint64": func(v any) uint64 { return 0 },
					"test":   func(v any) (uint64, error) { return 0, nil },
				},
			},
			want: []Info{
				{
					Name:        "uint64",
					Description: "uint64(any) uint64",
				},
				{
					Name:        "test",
					Description: "test(any) (uint64, error)",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(WithAddFuncMap(tt.args.fns)).FuncInfos()
			sort.Slice(got, func(i, j int) bool {
				return got[i].Name > got[j].Name
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FuncInfos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDisplayType(t *testing.T) {
	tests := []struct {
		name   string
		typeOf reflect.Type
		want   string
	}{
		{name: "any", typeOf: reflect.TypeOf((*any)(nil)).Elem(), want: "any"},
		{name: "slice", typeOf: reflect.TypeOf([]any{}), want: "[]any"},
		{name: "map", typeOf: reflect.TypeOf(map[string]any{}), want: "map[string]any"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := displayType(tt.typeOf); got != tt.want {
				t.Fatalf("displayType() = %q, want %q", got, tt.want)
			}
		})
	}
}
