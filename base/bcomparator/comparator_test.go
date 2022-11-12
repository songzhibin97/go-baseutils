package bcomparator

import (
	"math"
	"reflect"
	"testing"
)

func TestStringComparator(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want int
	}{
		{
			name: "eq",
			s1:   "",
			s2:   "",
			want: 0,
		},
		{
			name: "eq",
			s1:   "123",
			s2:   "123",
			want: 0,
		},
		{
			name: "eq",
			s1:   "abc",
			s2:   "abc",
			want: 0,
		},
		{
			name: "gt",
			s1:   "1",
			s2:   "",
			want: 1,
		},
		{
			name: "gt",
			s1:   "321",
			s2:   "123",
			want: 1,
		},
		{
			name: "gt",
			s1:   "a",
			s2:   "",
			want: 1,
		},
		{
			name: "gt",
			s1:   "cba",
			s2:   "abc",
			want: 1,
		},
		{
			name: "lt",
			s1:   "",
			s2:   "1",
			want: -1,
		},
		{
			name: "lt",
			s1:   "123",
			s2:   "321",
			want: -1,
		},
		{
			name: "lt",
			s1:   "",
			s2:   "a",
			want: -1,
		},
		{
			name: "lt",
			s1:   "abc",
			s2:   "cba",
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringComparator()(tt.s1, tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntComparator(t *testing.T) {
	tests := []struct {
		name string
		i1   int
		i2   int
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxInt,
			i2:   math.MaxInt,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MinInt,
			i2:   math.MinInt,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt,
			i2:   math.MinInt,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxInt,
			want: -1,
		},
		{
			name: "lt",
			i1:   math.MinInt,
			i2:   math.MaxInt,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntComparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   int8
		i2   int8
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxInt8,
			i2:   math.MaxInt8,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MinInt8,
			i2:   math.MinInt8,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt8,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt8,
			i2:   math.MinInt8,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxInt8,
			want: -1,
		},
		{
			name: "lt",
			i1:   math.MinInt8,
			i2:   math.MaxInt8,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   int16
		i2   int16
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxInt16,
			i2:   math.MaxInt16,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MinInt16,
			i2:   math.MinInt16,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt16,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt16,
			i2:   math.MinInt16,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxInt16,
			want: -1,
		},
		{
			name: "lt",
			i1:   math.MinInt16,
			i2:   math.MaxInt16,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   int32
		i2   int32
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxInt32,
			i2:   math.MaxInt32,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MinInt32,
			i2:   math.MinInt32,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt32,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt32,
			i2:   math.MinInt32,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxInt32,
			want: -1,
		},
		{
			name: "lt",
			i1:   math.MinInt32,
			i2:   math.MaxInt32,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   int64
		i2   int64
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxInt64,
			i2:   math.MaxInt64,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MinInt64,
			i2:   math.MinInt64,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt64,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxInt64,
			i2:   math.MinInt64,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxInt64,
			want: -1,
		},
		{
			name: "lt",
			i1:   math.MinInt64,
			i2:   math.MaxInt64,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintComparator(t *testing.T) {
	tests := []struct {
		name string
		i1   uint
		i2   uint
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxUint,
			i2:   math.MaxUint,
			want: 0,
		},
		{
			name: "eq",
			i1:   1,
			i2:   1,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint,
			i2:   1,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxUint,
			want: -1,
		},
		{
			name: "lt",
			i1:   1,
			i2:   math.MaxUint,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintComparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   uint8
		i2   uint8
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxUint8,
			i2:   math.MaxUint8,
			want: 0,
		},
		{
			name: "eq",
			i1:   1,
			i2:   1,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint8,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint8,
			i2:   1,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxUint8,
			want: -1,
		},
		{
			name: "lt",
			i1:   1,
			i2:   math.MaxUint8,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   uint16
		i2   uint16
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxUint16,
			i2:   math.MaxUint16,
			want: 0,
		},
		{
			name: "eq",
			i1:   1,
			i2:   1,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint16,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint16,
			i2:   1,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxUint16,
			want: -1,
		},
		{
			name: "lt",
			i1:   1,
			i2:   math.MaxUint16,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   uint32
		i2   uint32
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxUint32,
			i2:   math.MaxUint32,
			want: 0,
		},
		{
			name: "eq",
			i1:   1,
			i2:   1,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint32,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint32,
			i2:   1,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxUint32,
			want: -1,
		},
		{
			name: "lt",
			i1:   1,
			i2:   math.MaxUint32,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Comparator(t *testing.T) {
	tests := []struct {
		name string
		i1   uint64
		i2   uint64
		want int
	}{
		{
			name: "eq",
			i1:   0,
			i2:   0,
			want: 0,
		},
		{
			name: "eq",
			i1:   math.MaxUint64,
			i2:   math.MaxUint64,
			want: 0,
		},
		{
			name: "eq",
			i1:   1,
			i2:   1,
			want: 0,
		},
		{
			name: "gt",
			i1:   1,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint64,
			i2:   0,
			want: 1,
		},
		{
			name: "gt",
			i1:   math.MaxUint64,
			i2:   1,
			want: 1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   1,
			want: -1,
		},
		{
			name: "lt",
			i1:   0,
			i2:   math.MaxUint64,
			want: -1,
		},
		{
			name: "lt",
			i1:   1,
			i2:   math.MaxUint64,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Comparator()(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Comparator(t *testing.T) {
	tests := []struct {
		name string
		f1   float32
		f2   float32
		want int
	}{
		{
			name: "eq",
			f1:   0.1,
			f2:   0.1,
			want: 0,
		},
		{
			name: "eq",
			f1:   math.MaxFloat32,
			f2:   math.MaxFloat32,
			want: 0,
		},
		{
			name: "eq",
			f1:   0.0000000001,
			f2:   0.0000000002,
			want: 0,
		},
		{
			name: "gt",
			f1:   1.11,
			f2:   1.10,
			want: 1,
		},
		{
			name: "gt",
			f1:   math.MaxFloat32,
			f2:   1.11,
			want: 1,
		},
		{
			name: "gt",
			f1:   math.MaxFloat32,
			f2:   1.11,
			want: 1,
		},
		{
			name: "lt",
			f1:   1.10,
			f2:   1.11,
			want: -1,
		},
		{
			name: "lt",
			f1:   1.11,
			f2:   math.MaxFloat32,
			want: -1,
		},
		{
			name: "lt",
			f1:   1.11,
			f2:   math.MaxFloat32,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Comparator()(tt.f1, tt.f2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Comparator(t *testing.T) {
	tests := []struct {
		name string
		f1   float64
		f2   float64
		want int
	}{
		{
			name: "eq",
			f1:   0.1,
			f2:   0.1,
			want: 0,
		},
		{
			name: "eq",
			f1:   math.MaxFloat64,
			f2:   math.MaxFloat64,
			want: 0,
		},
		{
			name: "eq",
			f1:   0.0000000001,
			f2:   0.0000000002,
			want: 0,
		},
		{
			name: "gt",
			f1:   1.11,
			f2:   1.10,
			want: 1,
		},
		{
			name: "gt",
			f1:   math.MaxFloat64,
			f2:   1.11,
			want: 1,
		},
		{
			name: "gt",
			f1:   math.MaxFloat64,
			f2:   1.11,
			want: 1,
		},
		{
			name: "lt",
			f1:   1.10,
			f2:   1.11,
			want: -1,
		},
		{
			name: "lt",
			f1:   1.11,
			f2:   math.MaxFloat64,
			want: -1,
		},
		{
			name: "lt",
			f1:   1.11,
			f2:   math.MaxFloat64,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Comparator()(tt.f1, tt.f2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolComparator(t *testing.T) {
	tests := []struct {
		name string
		b1   bool
		b2   bool
		want int
	}{
		{
			name: "eq",
			b1:   true,
			b2:   true,
			want: 0,
		},
		{
			name: "eq",
			b1:   false,
			b2:   false,
			want: 0,
		},
		{
			name: "gt",
			b1:   true,
			b2:   false,
			want: 1,
		},

		{
			name: "lt",
			b1:   false,
			b2:   true,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolComparator()(tt.b1, tt.b2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseComparator(t *testing.T) {
	tests := []struct {
		name   string
		want   int
		i1, i2 int
	}{
		{
			name: "eq",
			want: 0,
			i1:   0,
			i2:   0,
		},
		{
			name: "lt",
			want: -1,
			i1:   1,
			i2:   0,
		},
		{
			name: "gt",
			want: 1,
			i1:   0,
			i2:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseComparator(IntComparator())(tt.i1, tt.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseComparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
