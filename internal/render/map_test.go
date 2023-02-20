package render

import (
	"reflect"
	"testing"
)

func Test_hold(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		hold map[string]interface{}
		want map[string]interface{}
	}{
		{
			name: "hold",
			args: args{
				key:   "key",
				value: "value",
			},
			want: map[string]interface{}{
				"key": "value",
			},
		},
		{
			name: "nested hold",
			args: args{
				key:   "key/nested/x",
				value: "value",
			},
			want: map[string]interface{}{
				"key": map[string]interface{}{
					"nested": map[string]interface{}{
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
			hold: map[string]interface{}{
				"key": map[string]interface{}{
					"nested": "value",
					"foo":    "bar",
				},
			},
			want: map[string]interface{}{
				"key": map[string]interface{}{
					"nested": map[string]interface{}{
						"x": "value",
					},
					"foo": "bar",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hold != nil {
				hMap = tt.hold
			}

			if got := Hold(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetData(t *testing.T) {
	type args struct {
		key  string
		data map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "get data",
			args: args{
				key: "key",
				data: map[string]interface{}{
					"key": "value",
				},
			},
			want: "value",
		},
		{
			name: "get nested data",
			args: args{
				key: "key/nested/x",
				data: map[string]interface{}{
					"key": map[string]interface{}{
						"nested": map[string]interface{}{
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
				data: map[string]interface{}{
					"key": map[string]interface{}{
						"nested": map[string]interface{}{
							"x": "value",
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetData(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
