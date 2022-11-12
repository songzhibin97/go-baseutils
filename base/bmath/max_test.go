package bmath

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestMax(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	type ts struct {
		name string
		args args
		want float64
	}
	var tests []ts

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		tests = append(tests, ts{
			name: strconv.Itoa(i),
			args: args{
				a: a,
				b: b,
			},
			want: math.Max(a, b),
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max[float64](tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
