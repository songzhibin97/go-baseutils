package bslice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnsafeComparableBSlice_Contains(t *testing.T) {
	type fields struct {
		es []int
	}
	type args struct {
		e int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				es: nil,
			},
			args: args{
				e: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			args: args{
				e: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 1,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 2,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 3,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Contains(tt.args.e), "Contains(%v)", tt.args.e)
		})
	}
}

func TestUnsafeComparableBSlice_Equal(t *testing.T) {
	type fields struct {
		es []int
	}
	type args struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				es: nil,
			},
			args: args{
				es: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			args: args{
				es: []int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			args: args{
				es: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: nil,
			},
			args: args{
				es: []int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Equal(tt.args.es), "Equal(%v)", tt.args.es)
		})
	}
}

func TestUnsafeComparableBSlice_Compact(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				es: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 1, 2, 2, 3, 3, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3, 3, 2, 1},
			},

			want: []int{1, 2, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBSliceBySlice(tt.fields.es)
			x.Compact()
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Compact")
		})
	}
}

func TestSafeComparableBSlice_Contains(t *testing.T) {
	type fields struct {
		es []int
	}
	type args struct {
		e int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				es: nil,
			},
			args: args{
				e: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			args: args{
				e: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 1,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 2,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 3,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				e: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Contains(tt.args.e), "Contains(%v)", tt.args.e)
		})
	}
}

func TestSafeComparableBSlice_Equal(t *testing.T) {
	type fields struct {
		es []int
	}
	type args struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				es: nil,
			},
			args: args{
				es: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			args: args{
				es: []int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			args: args{
				es: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: nil,
			},
			args: args{
				es: []int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBSliceBySlice(tt.fields.es)
			assert.Equalf(t, tt.want, x.Equal(tt.args.es), "Equal(%v)", tt.args.es)
		})
	}
}

func TestSafeComparableBSlice_Compact(t *testing.T) {
	type fields struct {
		es []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				es: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				es: []int{},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 1, 2, 2, 3, 3, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				es: []int{1, 2, 3, 3, 2, 1},
			},

			want: []int{1, 2, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBSliceBySlice(tt.fields.es)
			x.Compact()
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Compact")
		})
	}
}
