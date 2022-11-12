package bslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnsafeOrderedBSlice_Compare(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				es: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{1},
			},
			args: args{
				es: []int{},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: []int{1},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 2},
			},
			args: args{
				es: []int{2, 1, 1},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 1, 1},
			},
			args: args{
				es: []int{2, 2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeOrderedBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.Compare(tt.args.es), "Compare(%v)", tt.args.es)
		})
	}
}

func TestUnsafeOrderedBSlice_Sort(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 1, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeOrderedBSliceBySlice(tt.fields.e)
			x.Sort()
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Sort")
		})
	}
}

func TestUnsafeOrderedBSlice_IsSorted(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 1},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 1, 3},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{3, 2, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeOrderedBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.IsSorted(), "IsSorted()")
		})
	}
}

func TestUnsafeOrderedBSlice_BinarySearch(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		target int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 1,
			},
			want:  0,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 2,
			},
			want:  1,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 3,
			},
			want:  1,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 4,
			},
			want:  2,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 5,
			},
			want:  2,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 6,
			},
			want:  3,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 1,
			},
			want:  0,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 2,
			},
			want:  1,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 3,
			},
			want:  1,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 4,
			},
			want:  2,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 5,
			},
			want:  2,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 6,
			},
			want:  3,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 7,
			},
			want:  3,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 8,
			},
			want:  4,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeOrderedBSliceBySlice(tt.fields.e)
			got, got1 := x.BinarySearch(tt.args.target)
			assert.Equalf(t, tt.want, got, "BinarySearch(%v)", tt.args.target)
			assert.Equalf(t, tt.want1, got1, "BinarySearch(%v)", tt.args.target)
		})
	}
}

func TestSafeOrderedBSlice_Compare(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				es: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{1},
			},
			args: args{
				es: []int{},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: []int{1},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 2},
			},
			args: args{
				es: []int{2, 1, 1},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 1, 1},
			},
			args: args{
				es: []int{2, 2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeOrderedBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.Compare(tt.args.es), "Compare(%v)", tt.args.es)
		})
	}
}

func TestSafeOrderedBSlice_Sort(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 1, 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeOrderedBSliceBySlice(tt.fields.e)
			x.Sort()
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Sort")
		})
	}
}

func TestSafeOrderedBSlice_IsSorted(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 1},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{2, 1, 3},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{3, 2, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeOrderedBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.IsSorted(), "IsSorted()")
		})
	}
}

func TestSafeOrderedBSlice_BinarySearch(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		target int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 1,
			},
			want:  0,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 2,
			},
			want:  1,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 3,
			},
			want:  1,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 4,
			},
			want:  2,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 5,
			},
			want:  2,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5},
			},
			args: args{
				target: 6,
			},
			want:  3,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 1,
			},
			want:  0,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 2,
			},
			want:  1,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 3,
			},
			want:  1,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 4,
			},
			want:  2,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 5,
			},
			want:  2,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 6,
			},
			want:  3,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 7,
			},
			want:  3,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 3, 5, 7},
			},
			args: args{
				target: 8,
			},
			want:  4,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeOrderedBSliceBySlice(tt.fields.e)
			got, got1 := x.BinarySearch(tt.args.target)
			assert.Equalf(t, tt.want, got, "BinarySearch(%v)", tt.args.target)
			assert.Equalf(t, tt.want1, got1, "BinarySearch(%v)", tt.args.target)
		})
	}
}
