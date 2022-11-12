package bslice

import (
	"github.com/songzhibin97/go-baseutils/base/bternaryexpr"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestUnsafeCalculableBSlice_Sum(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100)
		list := make([]int, ln)
		want := 0
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			want += v
		}
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Sum(), "Sum()")
		})
	}
}

func TestUnsafeCalculableBSlice_Avg(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100)
		list := make([]int, 0, ln)
		want := 0
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			want += v
		}
		want /= bternaryexpr.TernaryExpr(ln == 0, 1, ln)
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Avg(), "Avg(%v)", tt.fields.es)
		})
	}
}

func TestUnsafeCalculableBSlice_Max(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100) + 1
		list := make([]int, 0, ln)
		want := math.MinInt
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			if want < v {
				want = v
			}
		}
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Max(), "Max(%v)", tt.fields.es)
		})
	}
}

func TestUnsafeCalculableBSlice_Min(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100) + 1
		list := make([]int, 0, ln)
		want := math.MaxInt
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			if want > v {
				want = v
			}
		}
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Min(), "Min(%v)", tt.fields.es)
		})
	}
}

func TestSafeCalculableBSlice_Sum(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100)
		list := make([]int, ln)
		want := 0
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			want += v
		}
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Sum(), "Sum()")
		})
	}
}

func TestSafeCalculableBSlice_Avg(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100)
		list := make([]int, 0, ln)
		want := 0
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			want += v
		}
		want /= bternaryexpr.TernaryExpr(ln == 0, 1, ln)
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Avg(), "Avg(%v)", tt.fields.es)
		})
	}
}

func TestSafeCalculableBSlice_Max(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100) + 1
		list := make([]int, 0, ln)
		want := math.MinInt
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			if want < v {
				want = v
			}
		}
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Max(), "Max(%v)", tt.fields.es)
		})
	}
}

func TestSafeCalculableBSlice_Min(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000; i++ {
		ln := rand.Intn(100) + 1
		list := make([]int, 0, ln)
		want := math.MaxInt
		for j := 0; j < ln; j++ {
			v := rand.Int()
			list = append(list, v)
			if want > v {
				want = v
			}
		}
		tests = append(tests, struct {
			name   string
			fields fields
			want   int
		}{name: "", fields: fields{es: list}, want: want})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeCalculableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Min(), "Min(%v)", tt.fields.es)
		})
	}
}
