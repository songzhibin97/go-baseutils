package bslice

import (
	"github.com/songzhibin97/go-baseutils/base/bcomparator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnsafeAnyBSlice_EqualFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		es []int
		f  func(int, int) bool
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
				e: nil,
			},
			args: args{
				es: nil,
				f:  nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
				f:  nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{},
				f:  nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2, 3},
				f:  nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 1},
			},
			args: args{
				es: []int{2, 2, 2},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 1},
			},
			args: args{
				es: []int{2, 2, 2},
				f: func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			if got := x.EqualFunc(tt.args.es, tt.args.f); got != tt.want {
				t.Errorf("EqualFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnsafeAnyBSlice_CompareFunc(t *testing.T) {
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
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			if got := x.CompareFunc(tt.args.es, bcomparator.IntComparator()); got != tt.want {
				t.Errorf("CompareFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnsafeAnyBSlice_IndexFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
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
				f: func(int2 int) bool {
					return true
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(int2 int) bool {
					return false
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == -1
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 1
				},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 2
				},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 3
				},
			},
			want: 2,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 4
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			if got := x.IndexFunc(tt.args.f); got != tt.want {
				t.Errorf("IndexFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnsafeAnyBSlice_Insert(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: nil,
			},
			want: nil,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5},
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want: []int{1, 4, 5, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 2,
				e: []int{4, 5},
			},
			want: []int{1, 2, 4, 5, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 3,
				e: []int{4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 4,
				e: []int{4, 5},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.Insert(tt.args.i, tt.args.e...)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_InsertE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		e []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5},
			wanterr: false,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want:    nil,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5, 1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want:    []int{1, 4, 5, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 2,
				e: []int{4, 5},
			},
			want:    []int{1, 2, 4, 5, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 3,
				e: []int{4, 5},
			},
			want:    []int{1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 4,
				e: []int{4, 5},
			},
			want:    []int{1, 2, 3},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			err := x.InsertE(tt.args.i, tt.args.e...)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_Delete(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: []int{0, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want: []int{0, 1, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.Delete(tt.args.i, tt.args.j)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_DeleteE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want:    []int{0, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want:    []int{0, 1, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want:    []int{0, 1, 2},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			err := x.DeleteE(tt.args.i, tt.args.j)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_DeleteToSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: []int{0, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want: []int{0, 1, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.DeleteToSlice(tt.args.i, tt.args.j), "DeleteToSlice(%v, %v)", tt.args.i, tt.args.j)
		})
	}
}

func TestUnsafeAnyBSlice_DeleteToSliceE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want:    nil,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want:    []int{0, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want:    []int{0, 1, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want:    []int{0, 1, 2},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			v, err := x.DeleteToSliceE(tt.args.i, tt.args.j)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestUnsafeAnyBSlice_DeleteToBSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: []int{0, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want: []int{0, 1, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.DeleteToBSlice(tt.args.i, tt.args.j).ToMetaSlice(), "DeleteToSlice(%v, %v)", tt.args.i, tt.args.j)
		})
	}
}

func TestUnsafeAnyBSlice_DeleteToBSliceE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want:    []int{},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want:    []int{0, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want:    []int{0, 1, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want:    []int{0, 1, 2},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			v, err := x.DeleteToBSliceE(tt.args.i, tt.args.j)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, v.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_Replace(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5, 0, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 4,
				e: []int{4, 5},
			},
			want: []int{4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 5,
				e: []int{4, 5},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 1,
				e: []int{4, 5},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 4,
				e: []int{4, 5},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 1,
				j: 3,
				e: []int{4, 5},
			},
			want: []int{0, 4, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.Replace(tt.args.i, tt.args.j, tt.args.e...)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_ReplaceE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
		e []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want:    nil,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want:    []int{},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5, 0, 1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 4,
				e: []int{4, 5},
			},
			want:    []int{4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 5,
				e: []int{4, 5},
			},
			want:    []int{0, 1, 2, 3},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 1,
				e: []int{4, 5},
			},
			want:    []int{0, 1, 2, 3},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 4,
				e: []int{4, 5},
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 1,
				j: 3,
				e: []int{4, 5},
			},
			want:    []int{0, 4, 5, 3},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			err := x.ReplaceE(tt.args.i, tt.args.j, tt.args.e...)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_CloneToSlice(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CloneToSlice(), "CloneToSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_CloneToBSlice(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CloneToBSlice().ToMetaSlice(), "CloneToSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_CompactFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 2, 2, 3, 3, 3},
			},
			args: args{
				func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			x.CompactFunc(tt.args.f)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_Grow(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
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
				i: 0,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 10,
			},
			want: 10,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 10,
			},
			want: 10,
		},
		{
			name: "",
			fields: fields{
				e: make([]int, 10),
			},
			args: args{
				i: 10,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			x.Grow(tt.args.i)
			assert.Equal(t, tt.want, cap(x.ToMetaSlice()))
		})
	}
}

func TestUnsafeAnyBSlice_Clip(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: make([]int, 10, 20),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.Clip()
		})
	}
}

func TestUnsafeAnyBSlice_SortFunc(t *testing.T) {
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
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			x.SortFunc(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_SortFuncToSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortFuncToSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestUnsafeAnyBSlice_SortFuncToBSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortFuncToBSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_SortStableFunc(t *testing.T) {
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
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			x.SortStableFunc(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_SortStableFuncToSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortStableFuncToSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestUnsafeAnyBSlice_SortStableFuncToBSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortStableFuncToBSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v.ToMetaSlice())
		})
	}
}

func TestUnsafeAnyBSlice_IsSortedFunc(t *testing.T) {
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
			x := NewUnsafeComparableBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.IsSortedFunc(func(i, j int) bool {
				return i < j
			}), "IsSortedFunc()")
		})
	}
}

func TestUnsafeAnyBSlice_BinarySearchFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		e int
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
				e: 0,
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
				e: 0,
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
				e: 0,
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
				e: 1,
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
				e: 2,
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
				e: 3,
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
				e: 4,
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
				e: 5,
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
				e: 6,
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
				e: 0,
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
				e: 1,
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
				e: 2,
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
				e: 3,
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
				e: 4,
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
				e: 5,
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
				e: 6,
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
				e: 7,
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
				e: 8,
			},
			want:  4,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			got, got1 := x.BinarySearchFunc(tt.args.e, bcomparator.IntComparator())
			assert.Equalf(t, tt.want, got, "BinarySearchFunc(%v, %v)", tt.args.e, bcomparator.IntComparator())
			assert.Equalf(t, tt.want1, got1, "BinarySearchFunc(%v, %v)", tt.args.e, bcomparator.IntComparator())
		})
	}
}

func TestUnsafeAnyBSlice_Filter(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{0, 2, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 != 0
				},
			},
			want: []int{1, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return false
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.Filter(tt.args.f)
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Filter(%v)", tt.args.f)
		})
	}
}

func TestUnsafeAnyBSlice_FilterToSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{0, 2, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 != 0
				},
			},
			want: []int{1, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return false
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.FilterToSlice(tt.args.f), "FilterToSlice(%v)", tt.args.f)
		})
	}
}

func TestUnsafeAnyBSlice_FilterToBSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{0, 2, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 != 0
				},
			},
			want: []int{1, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return false
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.FilterToBSlice(tt.args.f).ToMetaSlice(), "FilterToBSlice(%v)", tt.args.f)
		})
	}
}

func TestUnsafeAnyBSlice_Reverse(t *testing.T) {
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
				e: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.Reverse()
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Reverse()")
		})
	}
}

func TestUnsafeAnyBSlice_ReverseToSlice(t *testing.T) {
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
			want: []int{},
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
				e: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.ReverseToSlice(), "ReverseToSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_ReverseToBSlice(t *testing.T) {
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
			want: []int{},
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
				e: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.ReverseToBSlice().ToMetaSlice(), "ReverseToBSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_Marshal(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: []byte("[]"),
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: []byte("[]"),
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []byte("[1,2,3]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			got, err := x.Marshal()
			assert.Equal(t, false, err != nil)
			assert.Equalf(t, tt.want, got, "Marshal()")
		})
	}
}

func TestUnsafeAnyBSlice_Unmarshal(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				data: []byte("[]"),
			},
			want: []int{},
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				data: []byte("[1,2,3]"),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			err := x.Unmarshal(tt.args.data)
			assert.Equal(t, false, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())

		})
	}
}

func TestUnsafeAnyBSlice_Len(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 0, 10),
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10),
			},
			want: 10,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10, 20),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.Len(), "Len()")
		})
	}
}

func TestUnsafeAnyBSlice_Cap(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 0, 10),
			},
			want: 10,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10),
			},
			want: 10,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10, 20),
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.Cap(), "Cap()")
		})
	}
}

func TestUnsafeAnyBSlice_ToInterfaceSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: []interface{}{},
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			want: []interface{}{},
		},
		{
			name: "nil",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.ToInterfaceSlice(), "ToInterfaceSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_Append(t *testing.T) {
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
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				es: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.Append(tt.args.es...)
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Append()")
		})
	}
}

func TestUnsafeAnyBSlice_AppendToSlice(t *testing.T) {
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
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.AppendToSlice(tt.args.es...), "AppendToSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_AppendToBSlice(t *testing.T) {
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
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.AppendToBSlice(tt.args.es...).ToMetaSlice(), "AppendToBSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_CopyToSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name:   "",
			fields: fields{},

			want: []int{},
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
				e: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CopyToSlice(), "CopyToSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_CopyToBSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name:   "",
			fields: fields{},

			want: []int{},
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
				e: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CopyToBSlice().ToMetaSlice(), "CopyToBSlice()")
		})
	}
}

func TestUnsafeAnyBSlice_GetByIndex(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
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
				index: 0,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 1,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: -1,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
			},
			want: 2,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
			},
			want: 3,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.GetByIndex(tt.args.index), "GetByIndex(%v)", tt.args.index)
		})
	}
}

func TestUnsafeAnyBSlice_GetByIndexE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
			},
			want:    0,
			wanterr: true,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 1,
			},
			want:    0,
			wanterr: true,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: -1,
			},
			want:    0,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
			},
			want:    1,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
			},
			want:    2,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
			},
			want:    3,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
			},
			want:    0,
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			got, err := x.GetByIndexE(tt.args.index)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equalf(t, tt.want, got, "GetByIndexE(%v)", tt.args.index)
		})
	}
}

func TestUnsafeAnyBSlice_GetByIndexOrDefault(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index    int
		defaultE int
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
				index:    0,
				defaultE: 1,
			},
			want: 1,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			args: args{
				index:    0,
				defaultE: 1,
			},
			want: 1,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index:    0,
				defaultE: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.GetByIndexOrDefault(tt.args.index, tt.args.defaultE), "GetByIndexOrDefault(%v, %v)", tt.args.index, tt.args.defaultE)
		})
	}
}

func TestUnsafeAnyBSlice_GetByRange(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				start: 0,
				end:   0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.GetByRange(tt.args.start, tt.args.end), "GetByRange(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}

func TestUnsafeAnyBSlice_GetByRangeE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				start: 0,
				end:   0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   3,
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   4,
			},
			want:    nil,
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			got, err := x.GetByRangeE(tt.args.start, tt.args.end)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equalf(t, tt.want, got, "GetByRangeE(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}

func TestUnsafeAnyBSlice_SetByIndex(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		e     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: -1,
				e:     5,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
				e:     5,
			},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			x.SetByIndex(tt.args.index, tt.args.e)
			if !tt.wanterr {
				assert.Equalf(t, tt.args.e, x.GetByIndex(tt.args.index), "SetByIndex(%v, %v)", tt.args.index, tt.args.e)
			}
		})
	}
}

func TestUnsafeAnyBSlice_SetByIndexE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		e     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: -1,
				e:     5,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
				e:     5,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: -1,
				e:     5,
			},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			err := x.SetByIndexE(tt.args.index, tt.args.e)
			assert.Equal(t, tt.wanterr, err != nil)
			if !tt.wanterr {
				assert.Equalf(t, tt.args.e, x.GetByIndex(tt.args.index), "SetByIndexE(%v, %v)", tt.args.index, tt.args.e)
			}
		})
	}
}

func TestUnsafeAnyBSlice_SetByRange(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		es    []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: 2,
				es:    []int{3, 4, 5},
			},
			want:    []int{1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: -1,
				es:    []int{3, 4, 5},
			},
			want:    []int{},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			err := x.SetByRangeE(tt.args.index, tt.args.es)
			assert.Equal(t, tt.wanterr, err != nil)
			if !tt.wanterr {
				assert.Equalf(t, tt.want, x.ToMetaSlice(), "SetByRange()")
			}
		})
	}
}

func TestUnsafeAnyBSlice_SetByRangeE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		es    []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: 2,
				es:    []int{3, 4, 5},
			},
			want:    []int{1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: -1,
				es:    []int{3, 4, 5},
			},
			want:    []int{},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBSliceBySlice(tt.fields.e)
			err := x.SetByRangeE(tt.args.index, tt.args.es)
			assert.Equal(t, tt.wanterr, err != nil)
			if !tt.wanterr {
				assert.Equalf(t, tt.want, x.ToMetaSlice(), "SetByRangeE()")
			}
		})
	}
}

func TestSafeAnyBSlice_EqualFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		es []int
		f  func(int, int) bool
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
				e: nil,
			},
			args: args{
				es: nil,
				f:  nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
				f:  nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{},
				f:  nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{1, 2, 3},
				f:  nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 1},
			},
			args: args{
				es: []int{2, 2, 2},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 1},
			},
			args: args{
				es: []int{2, 2, 2},
				f: func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			if got := x.EqualFunc(tt.args.es, tt.args.f); got != tt.want {
				t.Errorf("EqualFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeAnyBSlice_CompareFunc(t *testing.T) {
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
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			if got := x.CompareFunc(tt.args.es, bcomparator.IntComparator()); got != tt.want {
				t.Errorf("CompareFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeAnyBSlice_IndexFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
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
				f: func(int2 int) bool {
					return true
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(int2 int) bool {
					return false
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == -1
				},
			},
			want: -1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 1
				},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 2
				},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 3
				},
			},
			want: 2,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				f: func(i int) bool {
					return i == 4
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			if got := x.IndexFunc(tt.args.f); got != tt.want {
				t.Errorf("IndexFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeAnyBSlice_Insert(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: nil,
			},
			want: nil,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5},
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want: []int{1, 4, 5, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 2,
				e: []int{4, 5},
			},
			want: []int{1, 2, 4, 5, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 3,
				e: []int{4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 4,
				e: []int{4, 5},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.Insert(tt.args.i, tt.args.e...)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_InsertE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		e []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5},
			wanterr: false,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want:    nil,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5, 1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 1,
				e: []int{4, 5},
			},
			want:    []int{1, 4, 5, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 2,
				e: []int{4, 5},
			},
			want:    []int{1, 2, 4, 5, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 3,
				e: []int{4, 5},
			},
			want:    []int{1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				i: 4,
				e: []int{4, 5},
			},
			want:    []int{1, 2, 3},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			err := x.InsertE(tt.args.i, tt.args.e...)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_Delete(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: []int{0, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want: []int{0, 1, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.Delete(tt.args.i, tt.args.j)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_DeleteE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want:    []int{0, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want:    []int{0, 1, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want:    []int{0, 1, 2},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			err := x.DeleteE(tt.args.i, tt.args.j)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_DeleteToSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: []int{0, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want: []int{0, 1, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.DeleteToSlice(tt.args.i, tt.args.j), "DeleteToSlice(%v, %v)", tt.args.i, tt.args.j)
		})
	}
}

func TestSafeAnyBSlice_DeleteToSliceE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want:    nil,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want:    []int{0, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want:    []int{0, 1, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want:    []int{0, 1, 2},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			v, err := x.DeleteToSliceE(tt.args.i, tt.args.j)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestSafeAnyBSlice_DeleteToBSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want: []int{0, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want: []int{0, 1, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.DeleteToBSlice(tt.args.i, tt.args.j).ToMetaSlice(), "DeleteToSlice(%v, %v)", tt.args.i, tt.args.j)
		})
	}
}

func TestSafeAnyBSlice_DeleteToBSliceE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: -1,
				j: -1,
			},
			want:    []int{},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 6,
				j: 6,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 0,
				j: 6,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 1,
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 1,
				j: 2,
			},
			want:    []int{0, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 2,
				j: 3,
			},
			want:    []int{0, 1, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4, 5},
			},
			args: args{
				i: 3,
				j: 6,
			},
			want:    []int{0, 1, 2},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			v, err := x.DeleteToBSliceE(tt.args.i, tt.args.j)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, v.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_Replace(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want: []int{4, 5, 0, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 4,
				e: []int{4, 5},
			},
			want: []int{4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 5,
				e: []int{4, 5},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 1,
				e: []int{4, 5},
			},
			want: []int{0, 1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 4,
				e: []int{4, 5},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 1,
				j: 3,
				e: []int{4, 5},
			},
			want: []int{0, 4, 5, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.Replace(tt.args.i, tt.args.j, tt.args.e...)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_ReplaceE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
		j int
		e []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want:    nil,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
				j: 1,
				e: []int{4, 5},
			},
			want:    []int{},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 0,
				e: []int{4, 5},
			},
			want:    []int{4, 5, 0, 1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 4,
				e: []int{4, 5},
			},
			want:    []int{4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 0,
				j: 5,
				e: []int{4, 5},
			},
			want:    []int{0, 1, 2, 3},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 1,
				e: []int{4, 5},
			},
			want:    []int{0, 1, 2, 3},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 4,
				j: 4,
				e: []int{4, 5},
			},
			want:    []int{0, 1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3},
			},
			args: args{
				i: 1,
				j: 3,
				e: []int{4, 5},
			},
			want:    []int{0, 4, 5, 3},
			wanterr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			err := x.ReplaceE(tt.args.i, tt.args.j, tt.args.e...)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_CloneToSlice(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CloneToSlice(), "CloneToSlice()")
		})
	}
}

func TestSafeAnyBSlice_CloneToBSlice(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CloneToBSlice().ToMetaSlice(), "CloneToSlice()")
		})
	}
}

func TestSafeAnyBSlice_CompactFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 1, 2, 2, 3, 3, 3},
			},
			args: args{
				func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3, 3, 2, 1},
			},
			args: args{
				func(i int, i2 int) bool {
					return i == i2
				},
			},
			want: []int{1, 2, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			x.CompactFunc(tt.args.f)
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_Grow(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		i int
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
				i: 0,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 0,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				i: 10,
			},
			want: 10,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				i: 10,
			},
			want: 10,
		},
		{
			name: "",
			fields: fields{
				e: make([]int, 10),
			},
			args: args{
				i: 10,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			x.Grow(tt.args.i)
			assert.Equal(t, tt.want, cap(x.ToMetaSlice()))
		})
	}
}

func TestSafeAnyBSlice_Clip(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: make([]int, 10, 20),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.Clip()
		})
	}
}

func TestSafeAnyBSlice_SortFunc(t *testing.T) {
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
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			x.SortFunc(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_SortFuncToSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortFuncToSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestSafeAnyBSlice_SortFuncToBSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortFuncToBSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_SortStableFunc(t *testing.T) {
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
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			x.SortStableFunc(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, x.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_SortStableFuncToSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortStableFuncToSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestSafeAnyBSlice_SortStableFuncToBSlice(t *testing.T) {
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
			want: []int{},
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
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			v := x.SortStableFuncToBSlice(func(i int, j int) bool {
				return i < j
			})
			assert.Equal(t, tt.want, v.ToMetaSlice())
		})
	}
}

func TestSafeAnyBSlice_IsSortedFunc(t *testing.T) {
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
			x := NewSafeComparableBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.IsSortedFunc(func(i, j int) bool {
				return i < j
			}), "IsSortedFunc()")
		})
	}
}

func TestSafeAnyBSlice_BinarySearchFunc(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		e int
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
				e: 0,
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
				e: 0,
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
				e: 0,
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
				e: 1,
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
				e: 2,
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
				e: 3,
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
				e: 4,
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
				e: 5,
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
				e: 6,
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
				e: 0,
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
				e: 1,
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
				e: 2,
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
				e: 3,
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
				e: 4,
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
				e: 5,
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
				e: 6,
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
				e: 7,
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
				e: 8,
			},
			want:  4,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			got, got1 := x.BinarySearchFunc(tt.args.e, bcomparator.IntComparator())
			assert.Equalf(t, tt.want, got, "BinarySearchFunc(%v, %v)", tt.args.e, bcomparator.IntComparator())
			assert.Equalf(t, tt.want1, got1, "BinarySearchFunc(%v, %v)", tt.args.e, bcomparator.IntComparator())
		})
	}
}

func TestSafeAnyBSlice_Filter(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{0, 2, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 != 0
				},
			},
			want: []int{1, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return false
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.Filter(tt.args.f)
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Filter(%v)", tt.args.f)
		})
	}
}

func TestSafeAnyBSlice_FilterToSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{0, 2, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 != 0
				},
			},
			want: []int{1, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return false
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.FilterToSlice(tt.args.f), "FilterToSlice(%v)", tt.args.f)
		})
	}
}

func TestSafeAnyBSlice_FilterToBSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		f func(int2 int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(int2 int) bool {
					return true
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{0, 2, 4},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return i%2 != 0
				},
			},
			want: []int{1, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{0, 1, 2, 3, 4},
			},
			args: args{
				f: func(i int) bool {
					return false
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.FilterToBSlice(tt.args.f).ToMetaSlice(), "FilterToBSlice(%v)", tt.args.f)
		})
	}
}

func TestSafeAnyBSlice_Reverse(t *testing.T) {
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
				e: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.Reverse()
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Reverse()")
		})
	}
}

func TestSafeAnyBSlice_ReverseToSlice(t *testing.T) {
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
			want: []int{},
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
				e: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.ReverseToSlice(), "ReverseToSlice()")
		})
	}
}

func TestSafeAnyBSlice_ReverseToBSlice(t *testing.T) {
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
			want: []int{},
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
				e: []int{1},
			},
			want: []int{1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.ReverseToBSlice().ToMetaSlice(), "ReverseToBSlice()")
		})
	}
}

func TestSafeAnyBSlice_Marshal(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: []byte("[]"),
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			want: []byte("[]"),
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []byte("[1,2,3]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			got, err := x.Marshal()
			assert.Equal(t, false, err != nil)
			assert.Equalf(t, tt.want, got, "Marshal()")
		})
	}
}

func TestSafeAnyBSlice_Unmarshal(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				data: []byte("[]"),
			},
			want: []int{},
		},
		{
			name:   "",
			fields: fields{},
			args: args{
				data: []byte("[1,2,3]"),
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			err := x.Unmarshal(tt.args.data)
			assert.Equal(t, false, err != nil)
			assert.Equal(t, tt.want, x.ToMetaSlice())

		})
	}
}

func TestSafeAnyBSlice_Len(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 0, 10),
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10),
			},
			want: 10,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10, 20),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.Len(), "Len()")
		})
	}
}

func TestSafeAnyBSlice_Cap(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 0, 10),
			},
			want: 10,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10),
			},
			want: 10,
		},
		{
			name: "nil",
			fields: fields{
				e: make([]int, 10, 20),
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.Cap(), "Cap()")
		})
	}
}

func TestSafeAnyBSlice_ToInterfaceSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			want: []interface{}{},
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			want: []interface{}{},
		},
		{
			name: "nil",
			fields: fields{
				e: []int{1, 2, 3},
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.ToInterfaceSlice(), "ToInterfaceSlice()")
		})
	}
}

func TestSafeAnyBSlice_Append(t *testing.T) {
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
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				es: nil,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{},
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.Append(tt.args.es...)
			assert.Equalf(t, tt.want, x.ToMetaSlice(), "Append()")
		})
	}
}

func TestSafeAnyBSlice_AppendToSlice(t *testing.T) {
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
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.AppendToSlice(tt.args.es...), "AppendToSlice()")
		})
	}
}

func TestSafeAnyBSlice_AppendToBSlice(t *testing.T) {
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
		want   []int
	}{
		{
			name:   "",
			fields: fields{},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				es: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				es: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.AppendToBSlice(tt.args.es...).ToMetaSlice(), "AppendToBSlice()")
		})
	}
}

func TestSafeAnyBSlice_CopyToSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name:   "",
			fields: fields{},

			want: []int{},
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
				e: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CopyToSlice(), "CopyToSlice()")
		})
	}
}

func TestSafeAnyBSlice_CopyToBSlice(t *testing.T) {
	type fields struct {
		e []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name:   "",
			fields: fields{},

			want: []int{},
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
				e: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: nil,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.CopyToBSlice().ToMetaSlice(), "CopyToBSlice()")
		})
	}
}

func TestSafeAnyBSlice_GetByIndex(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
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
				index: 0,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 1,
			},
			want: 0,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: -1,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
			},
			want: 2,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
			},
			want: 3,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.GetByIndex(tt.args.index), "GetByIndex(%v)", tt.args.index)
		})
	}
}

func TestSafeAnyBSlice_GetByIndexE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
			},
			want:    0,
			wanterr: true,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 1,
			},
			want:    0,
			wanterr: true,
		},
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				index: -1,
			},
			want:    0,
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
			},
			want:    1,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
			},
			want:    2,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
			},
			want:    3,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
			},
			want:    0,
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			got, err := x.GetByIndexE(tt.args.index)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equalf(t, tt.want, got, "GetByIndexE(%v)", tt.args.index)
		})
	}
}

func TestSafeAnyBSlice_GetByIndexOrDefault(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index    int
		defaultE int
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
				index:    0,
				defaultE: 1,
			},
			want: 1,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{},
			},
			args: args{
				index:    0,
				defaultE: 1,
			},
			want: 1,
		},
		{
			name: "nil",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index:    0,
				defaultE: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.GetByIndexOrDefault(tt.args.index, tt.args.defaultE), "GetByIndexOrDefault(%v, %v)", tt.args.index, tt.args.defaultE)
		})
	}
}

func TestSafeAnyBSlice_GetByRange(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				start: 0,
				end:   0,
			},
			want: nil,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			assert.Equalf(t, tt.want, x.GetByRange(tt.args.start, tt.args.end), "GetByRange(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}

func TestSafeAnyBSlice_GetByRangeE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "nil",
			fields: fields{
				e: nil,
			},
			args: args{
				start: 0,
				end:   0,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   0,
			},
			want:    []int{},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   3,
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				start: 0,
				end:   4,
			},
			want:    nil,
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			got, err := x.GetByRangeE(tt.args.start, tt.args.end)
			assert.Equal(t, tt.wanterr, err != nil)
			assert.Equalf(t, tt.want, got, "GetByRangeE(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}

func TestSafeAnyBSlice_SetByIndex(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		e     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: -1,
				e:     5,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
				e:     5,
			},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			x.SetByIndex(tt.args.index, tt.args.e)
			if !tt.wanterr {
				assert.Equalf(t, tt.args.e, x.GetByIndex(tt.args.index), "SetByIndex(%v, %v)", tt.args.index, tt.args.e)
			}
		})
	}
}

func TestSafeAnyBSlice_SetByIndexE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		e     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 0,
				e:     1,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 0,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: -1,
				e:     5,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 1,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 2,
				e:     5,
			},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: 3,
				e:     5,
			},
			wanterr: true,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2, 3},
			},
			args: args{
				index: -1,
				e:     5,
			},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			err := x.SetByIndexE(tt.args.index, tt.args.e)
			assert.Equal(t, tt.wanterr, err != nil)
			if !tt.wanterr {
				assert.Equalf(t, tt.args.e, x.GetByIndex(tt.args.index), "SetByIndexE(%v, %v)", tt.args.index, tt.args.e)
			}
		})
	}
}

func TestSafeAnyBSlice_SetByRange(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		es    []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: 2,
				es:    []int{3, 4, 5},
			},
			want:    []int{1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: -1,
				es:    []int{3, 4, 5},
			},
			want:    []int{},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			err := x.SetByRangeE(tt.args.index, tt.args.es)
			assert.Equal(t, tt.wanterr, err != nil)
			if !tt.wanterr {
				assert.Equalf(t, tt.want, x.ToMetaSlice(), "SetByRange()")
			}
		})
	}
}

func TestSafeAnyBSlice_SetByRangeE(t *testing.T) {
	type fields struct {
		e []int
	}
	type args struct {
		index int
		es    []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wanterr bool
	}{
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    nil,
			},
			want:    nil,
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 0,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: nil,
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{},
			},
			args: args{
				index: 2,
				es:    []int{1, 2, 3},
			},
			want:    []int{1, 2, 3},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: 2,
				es:    []int{3, 4, 5},
			},
			want:    []int{1, 2, 3, 4, 5},
			wanterr: false,
		},
		{
			name: "",
			fields: fields{
				e: []int{1, 2},
			},
			args: args{
				index: -1,
				es:    []int{3, 4, 5},
			},
			want:    []int{},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBSliceBySlice(tt.fields.e)
			err := x.SetByRangeE(tt.args.index, tt.args.es)
			assert.Equal(t, tt.wanterr, err != nil)
			if !tt.wanterr {
				assert.Equalf(t, tt.want, x.ToMetaSlice(), "SetByRangeE()")
			}
		})
	}
}
