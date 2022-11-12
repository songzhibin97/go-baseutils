package bternaryexpr

import (
	"testing"
)

func TestTernaryExpr(t *testing.T) {
	type args struct {
		boolExpr    bool
		trueReturn  interface{}
		falseReturn interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "base",
			args: args{
				boolExpr:    true,
				trueReturn:  1,
				falseReturn: 2,
			},
			want: 1,
		},
		{
			name: "base",
			args: args{
				boolExpr:    false,
				trueReturn:  1,
				falseReturn: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TernaryExpr(tt.args.boolExpr, tt.args.trueReturn, tt.args.falseReturn); got != tt.want {
				t.Errorf("TernaryExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
