package template

import (
	"sort"
	"testing"

	"github.com/go-test/deep"
)

func TestFuncInfos(t *testing.T) {
	type args struct {
		fns map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []Info
	}{
		{
			name: "TestFuncInfos",
			args: args{
				fns: map[string]interface{}{
					"uint64": func(v interface{}) uint64 { return 0 },
					"test":   func(v interface{}) (uint64, error) { return 0, nil },
				},
			},
			want: []Info{
				{
					Name:        "uint64",
					Description: "uint64(interface {}) uint64",
				},
				{
					Name:        "test",
					Description: "test(interface {}) (uint64, error)",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FuncInfos(tt.args.fns)
			sort.Slice(got, func(i, j int) bool {
				return got[i].Name > got[j].Name
			})

			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Errorf("FuncInfos() = %v", diff)
			}
		})
	}
}
