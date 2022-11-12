package bmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparableBMap(t *testing.T) {
	compBMap := NewUnsafeComparableBMapByMap[int, int](map[int]int{1: 1, 2: 2, 3: 3})
	cp := compBMap.CloneToMap()
	assert.Equal(t, true, compBMap.EqualByMap(cp))
	assert.Equal(t, false, compBMap.EqualByMap(map[int]int{1: 1, 3: 3}))
	cpm := compBMap.CloneToBMap()
	assert.Equal(t, true, compBMap.EqualByBMap(cpm))
	assert.Equal(t, false, compBMap.EqualByBMap(NewUnsafeAnyBMapByMap(map[int]int{1: 1, 3: 3})))

	assert.Equal(t, false, compBMap.Replace(1, 2, 3))
	assert.Equal(t, true, compBMap.Replace(1, 1, 3))
	v, ok := compBMap.Get(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, 3, v)
}

func TestUnsafeComparableBMap_EqualByMap(t *testing.T) {
	type fields struct {
		UnsafeAnyBMap ComparableBMap[int, int]
	}
	type args struct {
		m map[int]int
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
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBMapByMap(tt.fields.UnsafeAnyBMap.ToMetaMap())
			assert.Equalf(t, tt.want, x.EqualByMap(tt.args.m), "EqualByMap(%v)", tt.args.m)
		})
	}
}

func TestUnsafeComparableBMap_EqualByBMap(t *testing.T) {
	type fields struct {
		UnsafeAnyBMap ComparableBMap[int, int]
	}
	type args struct {
		m map[int]int
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
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				UnsafeAnyBMap: NewUnsafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeComparableBMapByMap(tt.fields.UnsafeAnyBMap.ToMetaMap())
			assert.Equalf(t, tt.want, x.EqualByBMap(NewUnsafeComparableBMapByMap(tt.args.m)), "EqualByBMap(%v)", tt.args.m)
		})
	}
}

func TestSafeComparableBMap_EqualByMap(t *testing.T) {
	type fields struct {
		SafeAnyBMap ComparableBMap[int, int]
	}
	type args struct {
		m map[int]int
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
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBMapByMap(tt.fields.SafeAnyBMap.ToMetaMap())
			assert.Equalf(t, tt.want, x.EqualByMap(tt.args.m), "EqualByMap(%v)", tt.args.m)
		})
	}
}

func TestSafeComparableBMap_EqualByBMap(t *testing.T) {
	type fields struct {
		SafeAnyBMap ComparableBMap[int, int]
	}
	type args struct {
		m map[int]int
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
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](nil),
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "nil",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				SafeAnyBMap: NewSafeComparableBMapByMap[int, int](map[int]int{1: 1}),
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeComparableBMapByMap(tt.fields.SafeAnyBMap.ToMetaMap())
			assert.Equalf(t, tt.want, x.EqualByBMap(NewSafeComparableBMapByMap(tt.args.m)), "EqualByBMap(%v)", tt.args.m)
		})
	}
}
