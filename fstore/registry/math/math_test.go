package math

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func TestMath_Sub(t *testing.T) {
	type args struct {
		a any
		v []any
	}
	tests := []struct {
		name    string
		m       Math
		args    args
		want    json.Number
		wantErr bool
	}{
		{
			name: "Test 1",
			m:    Math{},
			args: args{
				a: decimal.NewFromFloat(1.323),
				v: []any{1.4324, 0.00321, 0.1},
			},
			want: json.Number("-0.21261"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Math{}
			got, err := m.Sub(tt.args.a, tt.args.v...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Math.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Math.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
