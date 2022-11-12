package bmath

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	type args struct {
		a float64
	}
	type ts struct {
		name string
		args args
		want float64
	}
	var tests []ts

	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		v := rand.Intn(100) - 50
		tests = append(tests, ts{
			name: strconv.Itoa(v),
			args: args{
				a: float64(v),
			},
			want: math.Abs(float64(v)),
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.a); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
