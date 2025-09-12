package maps

import (
	"reflect"
	"testing"
)

func Test_hold(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		args args
		hold map[string]any
		want map[string]any
	}{
		{
			name: "hold",
			args: args{
				key:   "key",
				value: "value",
			},
			want: map[string]any{
				"key": "value",
			},
		},
		{
			name: "nested hold",
			args: args{
				key:   "key/nested/x",
				value: "value",
			},
			want: map[string]any{
				"key": map[string]any{
					"nested": map[string]any{
						"x": "value",
					},
				},
			},
		},
		{
			name: "nested hold with value",
			args: args{
				key:   "key/nested/x",
				value: "value",
			},
			hold: map[string]any{
				"key": map[string]any{
					"nested": "value",
					"foo":    "bar",
				},
			},
			want: map[string]any{
				"key": map[string]any{
					"nested": map[string]any{
						"x": "value",
					},
					"foo": "bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hold == nil {
				tt.hold = make(map[string]any)
			}

			m := Map{value: tt.hold}

			if got := m.Set(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	type args struct {
		key  string
		data map[string]any
		hold map[string]any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "get data",
			args: args{
				key: "key",
				data: map[string]any{
					"key": "value",
				},
			},
			want: "value",
		},
		{
			name: "get data with nil",
			args: args{
				key: "key",
			},
			want: nil,
		},
		{
			name: "get nested data",
			args: args{
				key: "key/nested/x",
				data: map[string]any{
					"key": map[string]any{
						"nested": map[string]any{
							"x": "value",
						},
					},
				},
			},
			want: "value",
		},
		{
			name: "unknown key",
			args: args{
				key: "key/nested/y",
				data: map[string]any{
					"key": map[string]any{
						"nested": map[string]any{
							"x": "value",
						},
					},
				},
			},
			want: nil,
		},
		{
			name: "from hold",
			args: args{
				key: "key/nested/y",
				hold: map[string]any{
					"key": map[string]any{
						"nested": map[string]any{
							"y": "value",
						},
					},
				},
			},
			want: "value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Map{}
			m.value = tt.args.hold

			if got := m.Get(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
