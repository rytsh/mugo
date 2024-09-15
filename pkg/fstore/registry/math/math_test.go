package math

import "testing"

func TestMath_RoundDecimal(t *testing.T) {
	type args struct {
		precision int
		value     float64
	}
	tests := []struct {
		name string
		m    Math
		args args
		want float64
	}{
		{
			name: "round 1.2345 to 1.2",
			m:    Math{},
			args: args{
				precision: 1,
				value:     1.2345,
			},
			want: 1.2,
		},
		{
			name: "round 1.2345 to 1.23",
			m:    Math{},
			args: args{
				precision: 2,
				value:     1.2345,
			},
			want: 1.23,
		},
		{
			name: "round 1.2345 to 1.235",
			m:    Math{},
			args: args{
				precision: 3,
				value:     1.2345,
			},
			want: 1.235,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Math{}
			if got := m.RoundDecimal(tt.args.precision, tt.args.value); got != tt.want {
				t.Errorf("Math.RoundDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
