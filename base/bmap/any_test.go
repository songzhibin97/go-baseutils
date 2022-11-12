package bmap

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type tInterface interface {
	test() any
}

type tStruct struct {
	t any
}

func (t tStruct) test() any {
	return t
}

func TestAnyBMap(t *testing.T) {
	t1 := tStruct{t: 1}
	t2 := tStruct{t: 2}
	t3 := tInterface(nil)

	anyBMap := NewUnsafeAnyBMap[int, tInterface]()
	// map[int]tInterface{}

	assert.Equal(t, anyBMap.Keys(), []int{})
	assert.Equal(t, anyBMap.Values(), []tInterface{})
	assert.Equal(t, anyBMap.Size(), 0)
	assert.Equal(t, false, anyBMap.EqualFuncByMap(map[int]tInterface{1: t1, 2: t2, 3: t3}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, false, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{1: t1, 2: t2, 3: t3}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.IsEmpty())

	assert.Equal(t, false, anyBMap.IsExist(1))
	assert.Equal(t, false, anyBMap.ContainsKey(1))
	assert.Equal(t, false, anyBMap.ContainsValue(t1))

	// map[int]tInterface{}
	v1, ok := anyBMap.Get(1)
	assert.Equal(t, false, ok)
	if v1 != nil {
		t.Error("v1 should be nil")
	}
	v2 := anyBMap.GetOrDefault(1, t2)
	assert.Equal(t, t2, v2)

	// map[int]tInterface{1:t1}

	anyBMap.Put(1, t1)
	assert.Equal(t, anyBMap.Keys(), []int{1})
	assert.Equal(t, anyBMap.Values(), []tInterface{t1})
	assert.Equal(t, anyBMap.Size(), 1)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(map[int]tInterface{1: t1}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{1: t1}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, false, anyBMap.IsEmpty())
	assert.Equal(t, true, anyBMap.IsExist(1))
	assert.Equal(t, true, anyBMap.ContainsKey(1))
	assert.Equal(t, true, anyBMap.ContainsValue(t1))

	// map[int]tInterface{1:t1}
	v1, ok = anyBMap.Get(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, t1, v1)
	v2 = anyBMap.GetOrDefault(1, t2)
	assert.Equal(t, t1, v2)

	assert.Equal(t, false, anyBMap.PuTIfAbsent(1, t2))
	assert.Equal(t, anyBMap.Keys(), []int{1})
	assert.Equal(t, anyBMap.Values(), []tInterface{t1})
	assert.Equal(t, anyBMap.Size(), 1)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(map[int]tInterface{1: t1}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{1: t1}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, false, anyBMap.IsEmpty())
	assert.Equal(t, true, anyBMap.IsExist(1))
	assert.Equal(t, true, anyBMap.ContainsKey(1))
	assert.Equal(t, true, anyBMap.ContainsValue(t1))
	v1, ok = anyBMap.Get(1)
	assert.Equal(t, true, ok)
	assert.Equal(t, t1, v1)

	// map[int]tInterface{1:t1}

	v2 = anyBMap.GetOrDefault(1, t2)
	assert.Equal(t, t1, v2)

	assert.Equal(t, true, anyBMap.PuTIfAbsent(2, t2))

	// map[int]tInterface{1:t1, 2:t2}
	ks := anyBMap.Keys()
	vs := anyBMap.Values()
	sort.Ints(ks)
	sort.Slice(vs, func(i, j int) bool {
		if vs[i] != nil && vs[j] != nil {
			return vs[i].test().(tStruct).t.(int) < vs[j].test().(tStruct).t.(int)
		}
		return false
	})

	assert.Equal(t, ks, []int{1, 2})
	assert.Equal(t, vs, []tInterface{t1, t2})
	assert.Equal(t, anyBMap.Size(), 2)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(map[int]tInterface{1: t1, 2: t2}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{1: t1, 2: t2}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, false, anyBMap.IsEmpty())
	assert.Equal(t, true, anyBMap.IsExist(2))
	assert.Equal(t, true, anyBMap.ContainsKey(2))
	assert.Equal(t, true, anyBMap.ContainsValue(t2))
	// map[int]tInterface{1:t1, 2:t2}
	v1, ok = anyBMap.Get(2)
	assert.Equal(t, true, ok)
	assert.Equal(t, t2, v1)
	v2 = anyBMap.GetOrDefault(2, t1)
	assert.Equal(t, t2, v2)

	anyBMap.Delete(3)
	// map[int]tInterface{1:t1, 2:t2}

	ks = anyBMap.Keys()
	vs = anyBMap.Values()
	sort.Ints(ks)
	sort.Slice(vs, func(i, j int) bool {
		if vs[i] != nil && vs[j] != nil {
			return vs[i].test().(tStruct).t.(int) < vs[j].test().(tStruct).t.(int)
		}
		return false
	})

	assert.Equal(t, ks, []int{1, 2})
	assert.Equal(t, vs, []tInterface{t1, t2})
	assert.Equal(t, anyBMap.Size(), 2)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(map[int]tInterface{1: t1, 2: t2}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{1: t1, 2: t2}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, false, anyBMap.IsEmpty())

	d, ok := anyBMap.DeleteIfPresent(3)

	// map[int]tInterface{1:t1, 2:t2}

	assert.Equal(t, false, ok)
	if d != nil {
		t.Error("d should be nil")
	}

	ks = anyBMap.Keys()
	vs = anyBMap.Values()
	sort.Ints(ks)
	sort.Slice(vs, func(i, j int) bool {
		if vs[i] != nil && vs[j] != nil {
			return vs[i].test().(tStruct).t.(int) < vs[j].test().(tStruct).t.(int)
		}
		return false
	})
	assert.Equal(t, ks, []int{1, 2})
	assert.Equal(t, vs, []tInterface{t1, t2})
	assert.Equal(t, anyBMap.Size(), 2)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(map[int]tInterface{1: t1, 2: t2}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{1: t1, 2: t2}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))

	d2, ok := anyBMap.DeleteIfPresent(1)

	// map[int]tInterface{2:t2}
	assert.Equal(t, true, ok)
	assert.Equal(t, t1, d2)

	assert.Equal(t, anyBMap.Keys(), []int{2})
	assert.Equal(t, anyBMap.Values(), []tInterface{t2})
	assert.Equal(t, anyBMap.Size(), 1)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(map[int]tInterface{2: t2}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{2: t2}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))

	assert.Equal(t, false, anyBMap.ContainsKey(1))
	assert.Equal(t, false, anyBMap.ContainsValue(t1))

	// map[int]tInterface{2:t2}

	assert.Equal(t, false, anyBMap.Replace(2, t1, t3))
	assert.Equal(t, true, anyBMap.Replace(2, t2, t3))
	assert.Equal(t, true, anyBMap.ContainsValue(t3))

	anyBMap.MergeByMap(map[int]tInterface{2: t1}, func(k int, v tInterface) bool {
		return k != 2
	})

	assert.Equal(t, anyBMap.Keys(), []int{2})
	assert.Equal(t, anyBMap.Values(), []tInterface{t3})
	assert.Equal(t, true, anyBMap.Replace(2, t3, t1))
	assert.Equal(t, anyBMap.Size(), 1)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(map[int]tInterface{2: t1}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{2: t1}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))

	// map[int]tInterface{2:t2}
	anyBMap.MergeByMap(map[int]tInterface{1: t1, 2: t2, 3: t3}, func(k int, v tInterface) bool {
		return k != 2
	})

	assert.Equal(t, len(anyBMap.Keys()), 3)
	assert.Equal(t, len(anyBMap.Values()), 3)
	assert.Equal(t, anyBMap.Size(), 3)
	v1, ok = anyBMap.Get(2)
	assert.Equal(t, true, ok)
	assert.Equal(t, v1, t1)

	c1 := anyBMap.CloneToMap()
	assert.Equal(t, true, anyBMap.EqualFuncByMap(c1, func(v1, v2 tInterface) bool {
		if v1 == v2 {
			return true
		}
		return v1.test() == v2.test()
	}))
	c2 := anyBMap.CloneToBMap()
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(c2, func(v1, v2 tInterface) bool {
		if v1 == v2 {
			return true
		}
		return v1.test() == v2.test()
	}))

	mp := make(map[int]tInterface)
	anyBMap.CopyByMap(mp)
	assert.Equal(t, true, anyBMap.EqualFuncByMap(mp, func(v1, v2 tInterface) bool {
		if v1 == v2 {
			return true
		}
		return v1.test() == v2.test()
	}))
	bmp := NewUnsafeAnyBMap[int, tInterface]()
	anyBMap.CopyByBMap(bmp)
	assert.Equal(t, true, anyBMap.EqualFuncByBMap(bmp, func(v1, v2 tInterface) bool {
		if v1 == v2 {
			return true
		}
		return v1.test() == v2.test()
	}))

	anyBMap.Clear()
	assert.Equal(t, anyBMap.Keys(), []int{})
	assert.Equal(t, anyBMap.Values(), []tInterface{})
	assert.Equal(t, anyBMap.Size(), 0)
	assert.Equal(t, false, anyBMap.EqualFuncByMap(map[int]tInterface{1: t1, 2: t2, 3: t3}, func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, false, anyBMap.EqualFuncByBMap(NewUnsafeAnyBMapByMap(map[int]tInterface{1: t1, 2: t2, 3: t3}), func(v1, v2 tInterface) bool {
		return v1.test() == v2.test()
	}))
	assert.Equal(t, true, anyBMap.IsEmpty())
}

func TestUnsafeAnyBMap_ToMetaMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.ToMetaMap(), "ToMetaMap()")
		})
	}
}

func TestUnsafeAnyBMap_Keys(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			v := x.Keys()
			sort.Ints(v)
			assert.Equalf(t, tt.want, v, "Keys()")
		})
	}
}

func TestUnsafeAnyBMap_Values(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			v := x.Values()
			sort.Ints(v)
			assert.Equalf(t, tt.want, v, "Values()")
		})
	}
}

func TestUnsafeAnyBMap_EqualFuncByMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	eq := func(v1, v2 int) bool {
		return v1 == v2
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
				mp: nil,
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1, 2: 2},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.EqualFuncByMap(tt.args.m, eq), "EqualFuncByMap(%v)", tt.args.m)
		})
	}
}

func TestUnsafeAnyBMap_EqualFuncByBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	eq := func(v1, v2 int) bool {
		return v1 == v2
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
				mp: nil,
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1, 2: 2},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.EqualFuncByBMap(NewUnsafeAnyBMapByMap(tt.args.m), eq), "EqualFuncByBMap(%v)", tt.args.m)
		})
	}
}

func TestUnsafeAnyBMap_Clear(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.Clear()
			assert.Equalf(t, map[int]int{}, x.ToMetaMap(), "Clear()")
		})
	}
}

func TestUnsafeAnyBMap_CloneToMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: map[int]int{1: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.CloneToMap(), "CloneToMap()")
		})
	}
}

func TestUnsafeAnyBMap_CloneToBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   AnyBMap[int, int]
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: NewUnsafeAnyBMapByMap[int, int](nil),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: NewUnsafeAnyBMapByMap[int, int](map[int]int{}),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: NewUnsafeAnyBMapByMap[int, int](map[int]int{1: 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.CloneToBMap(), "CloneToBMap()")
		})
	}
}

func TestUnsafeAnyBMap_CopyByMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		dst map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				dst: map[int]int{},
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				dst: map[int]int{},
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				dst: map[int]int{1: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.CopyByMap(tt.args.dst)
			assert.Equalf(t, tt.args.dst, x.ToMetaMap(), "CopyByMap()")
		})
	}
}

func TestUnsafeAnyBMap_CopyByBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		dst AnyBMap[int, int]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				dst: NewUnsafeAnyBMapByMap[int, int](nil),
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				dst: NewUnsafeAnyBMapByMap(map[int]int{}),
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				dst: NewUnsafeAnyBMapByMap(map[int]int{1: 1}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.CopyByBMap(tt.args.dst)
			assert.Equalf(t, tt.args.dst, x, "CopyByBMap()")
		})
	}
}

func TestUnsafeAnyBMap_DeleteFunc(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		del func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				del: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return i == 1
				},
			},
			want: map[int]int{2: 2, 3: 3},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return i2 == 1
				},
			},
			want: map[int]int{2: 2, 3: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.DeleteFunc(tt.args.del)
			assert.Equalf(t, tt.want, x.ToMetaMap(), "DeleteFunc()")
		})
	}
}

func TestUnsafeAnyBMap_Marshal(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: []byte("{}"),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: []byte("{}"),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			want: []byte("{\"1\":1,\"2\":2}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			got, err := x.Marshal()
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "Marshal()")
		})
	}
}

func TestUnsafeAnyBMap_Unmarshal(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				data: []byte("{}"),
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				data: []byte("{\"1\":1,\"2\":2}"),
			},
			want: map[int]int{1: 1, 2: 2},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				data: []byte("{\"1\":1,\"2\":2}"),
			},
			want: map[int]int{1: 1, 2: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			err := x.Unmarshal(tt.args.data)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, x.ToMetaMap(), "Unmarshal()")
		})
	}
}

func TestUnsafeAnyBMap_Size(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.Size(), "Size()")
		})
	}
}

func TestUnsafeAnyBMap_IsEmpty(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestUnsafeAnyBMap_IsExist(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.IsExist(tt.args.k), "IsExist(%v)", tt.args.k)
		})
	}
}

func TestUnsafeAnyBMap_ContainsKey(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.ContainsKey(tt.args.k), "ContainsKey(%v)", tt.args.k)
		})
	}
}

func TestUnsafeAnyBMap_ContainsValue(t *testing.T) {
	type fields struct {
		mp map[int]tInterface
	}
	type args struct {
		v tInterface
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
				mp: nil,
			},
			args: args{
				v: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{},
			},
			args: args{
				v: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil},
			},
			args: args{
				v: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil},
			},
			args: args{
				v: &tStruct{
					t: nil,
				},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil},
			},
			args: args{
				v: &tStruct{
					t: 1,
				},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil, 2: &tStruct{
					t: 1,
				}},
			},
			args: args{
				v: &tStruct{
					t: 1,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.ContainsValue(tt.args.v), "ContainsValue(%v)", tt.args.v)
		})
	}
}

func TestUnsafeAnyBMap_Get(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 1,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 1,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 1,
			},
			want:  1,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 3,
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			got, got1 := x.Get(tt.args.k)
			assert.Equalf(t, tt.want, got, "Get(%v)", tt.args.k)
			assert.Equalf(t, tt.want1, got1, "Get(%v)", tt.args.k)
		})
	}
}

func TestUnsafeAnyBMap_GetOrDefault(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k            int
		defaultValue int
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
				mp: nil,
			},
			args: args{
				k:            0,
				defaultValue: 1,
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k:            0,
				defaultValue: 1,
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k:            1,
				defaultValue: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.GetOrDefault(tt.args.k, tt.args.defaultValue), "GetOrDefault(%v, %v)", tt.args.k, tt.args.defaultValue)
		})
	}
}

func TestUnsafeAnyBMap_Put(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				k: 1,
				v: 1,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 2,
				v: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.Put(tt.args.k, tt.args.v)
			v, ok := x.Get(tt.args.k)
			assert.Equal(t, true, ok)
			assert.Equal(t, tt.args.v, v)
		})
	}
}

func TestUnsafeAnyBMap_PuTIfAbsent(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "nil",
			fields: fields{},
			args: args{
				k: 1,
				v: 1,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 1,
				v: 1,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 2},
			},
			args: args{
				k: 1,
				v: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.PuTIfAbsent(tt.args.k, tt.args.v), "PuTIfAbsent(%v, %v)", tt.args.k, tt.args.v)
		})
	}
}

func TestUnsafeAnyBMap_Delete(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				k: 0,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.Delete(tt.args.k)
			_, ok := x.Get(tt.args.k)
			assert.Equal(t, false, ok)
		})
	}
}

func TestUnsafeAnyBMap_DeleteIfPresent(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 1,
			},
			want:  1,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			got, got1 := x.DeleteIfPresent(tt.args.k)
			assert.Equalf(t, tt.want, got, "DeleteIfPresent(%v)", tt.args.k)
			assert.Equalf(t, tt.want1, got1, "DeleteIfPresent(%v)", tt.args.k)
			_, ok := x.Get(tt.args.k)
			assert.Equal(t, false, ok)
		})
	}
}

func TestUnsafeAnyBMap_MergeByMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		m map[int]int
		f func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: nil,
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 2},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 4, 4: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.MergeByMap(tt.args.m, tt.args.f)
			assert.Equal(t, tt.want, x.ToMetaMap())
		})
	}
}

func TestUnsafeAnyBMap_MergeByBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		m map[int]int
		f func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: nil,
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 2},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 4, 4: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			x.MergeByBMap(NewUnsafeAnyBMapByMap(tt.args.m), tt.args.f)
			assert.Equal(t, tt.want, x.ToMetaMap())
		})
	}
}

func TestUnsafeAnyBMap_Replace(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k  int
		ov int
		nv int
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
				mp: nil,
			},
			args: args{
				k:  0,
				ov: 0,
				nv: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k:  0,
				ov: 0,
				nv: 0,
			},
			want: false,
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k:  1,
				ov: 0,
				nv: 2,
			},
			want: false,
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k:  1,
				ov: 1,
				nv: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewUnsafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.Replace(tt.args.k, tt.args.ov, tt.args.nv), "Replace(%v, %v, %v)", tt.args.k, tt.args.ov, tt.args.nv)
		})
	}
}

func TestSafeAnyBMap_ToMetaMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.ToMetaMap(), "ToMetaMap()")
		})
	}
}

func TestSafeAnyBMap_Keys(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			v := x.Keys()
			sort.Ints(v)
			assert.Equalf(t, tt.want, v, "Keys()")
		})
	}
}

func TestSafeAnyBMap_Values(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: []int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			v := x.Values()
			sort.Ints(v)
			assert.Equalf(t, tt.want, v, "Values()")
		})
	}
}

func TestSafeAnyBMap_EqualFuncByMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	eq := func(v1, v2 int) bool {
		return v1 == v2
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
				mp: nil,
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1, 2: 2},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.EqualFuncByMap(tt.args.m, eq), "EqualFuncByMap(%v)", tt.args.m)
		})
	}
}

func TestSafeAnyBMap_EqualFuncByBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	eq := func(v1, v2 int) bool {
		return v1 == v2
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
				mp: nil,
			},
			args: args{
				m: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 1, 2: 2},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				m: map[int]int{1: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.EqualFuncByBMap(NewSafeAnyBMapByMap(tt.args.m), eq), "EqualFuncByBMap(%v)", tt.args.m)
		})
	}
}

func TestSafeAnyBMap_Clear(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.Clear()
			assert.Equalf(t, map[int]int{}, x.ToMetaMap(), "Clear()")
		})
	}
}

func TestSafeAnyBMap_CloneToMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: map[int]int{1: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.CloneToMap(), "CloneToMap()")
		})
	}
}

func TestSafeAnyBMap_CloneToBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   AnyBMap[int, int]
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: NewSafeAnyBMapByMap[int, int](nil),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: NewSafeAnyBMapByMap[int, int](map[int]int{}),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: NewSafeAnyBMapByMap[int, int](map[int]int{1: 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want.ToMetaMap(), x.CloneToBMap().ToMetaMap(), "CloneToBMap()")
		})
	}
}

func TestSafeAnyBMap_CopyByMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		dst map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				dst: map[int]int{},
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				dst: map[int]int{},
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				dst: map[int]int{1: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.CopyByMap(tt.args.dst)
			assert.Equalf(t, tt.args.dst, x.ToMetaMap(), "CopyByMap()")
		})
	}
}

func TestSafeAnyBMap_CopyByBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		dst AnyBMap[int, int]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				dst: NewSafeAnyBMapByMap[int, int](nil),
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				dst: NewSafeAnyBMapByMap(map[int]int{}),
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				dst: NewSafeAnyBMapByMap(map[int]int{1: 1}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.CopyByBMap(tt.args.dst)
			assert.Equalf(t, tt.args.dst, x, "CopyByBMap()")
		})
	}
}

func TestSafeAnyBMap_DeleteFunc(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		del func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				del: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return i == 1
				},
			},
			want: map[int]int{2: 2, 3: 3},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				del: func(i int, i2 int) bool {
					return i2 == 1
				},
			},
			want: map[int]int{2: 2, 3: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.DeleteFunc(tt.args.del)
			assert.Equalf(t, tt.want, x.ToMetaMap(), "DeleteFunc()")
		})
	}
}

func TestSafeAnyBMap_Marshal(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: []byte("{}"),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: []byte("{}"),
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			want: []byte("{\"1\":1,\"2\":2}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			got, err := x.Marshal()
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "Marshal()")
		})
	}
}

func TestSafeAnyBMap_Unmarshal(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				data: []byte("{}"),
			},
			want: map[int]int{},
		},
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				data: []byte("{\"1\":1,\"2\":2}"),
			},
			want: map[int]int{1: 1, 2: 2},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				data: []byte("{\"1\":1,\"2\":2}"),
			},
			want: map[int]int{1: 1, 2: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			err := x.Unmarshal(tt.args.data)
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, x.ToMetaMap(), "Unmarshal()")
		})
	}
}

func TestSafeAnyBMap_Size(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: 0,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.Size(), "Size()")
		})
	}
}

func TestSafeAnyBMap_IsEmpty(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestSafeAnyBMap_IsExist(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.IsExist(tt.args.k), "IsExist(%v)", tt.args.k)
		})
	}
}

func TestSafeAnyBMap_ContainsKey(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.ContainsKey(tt.args.k), "ContainsKey(%v)", tt.args.k)
		})
	}
}

func TestSafeAnyBMap_ContainsValue(t *testing.T) {
	type fields struct {
		mp map[int]tInterface
	}
	type args struct {
		v tInterface
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
				mp: nil,
			},
			args: args{
				v: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{},
			},
			args: args{
				v: nil,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil},
			},
			args: args{
				v: nil,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil},
			},
			args: args{
				v: &tStruct{
					t: nil,
				},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil},
			},
			args: args{
				v: &tStruct{
					t: 1,
				},
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]tInterface{1: nil, 2: &tStruct{
					t: 1,
				}},
			},
			args: args{
				v: &tStruct{
					t: 1,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.ContainsValue(tt.args.v), "ContainsValue(%v)", tt.args.v)
		})
	}
}

func TestSafeAnyBMap_Get(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 1,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 1,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 1,
			},
			want:  1,
			want1: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 3,
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			got, got1 := x.Get(tt.args.k)
			assert.Equalf(t, tt.want, got, "Get(%v)", tt.args.k)
			assert.Equalf(t, tt.want1, got1, "Get(%v)", tt.args.k)
		})
	}
}

func TestSafeAnyBMap_GetOrDefault(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k            int
		defaultValue int
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
				mp: nil,
			},
			args: args{
				k:            0,
				defaultValue: 1,
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k:            0,
				defaultValue: 1,
			},
			want: 1,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k:            1,
				defaultValue: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.GetOrDefault(tt.args.k, tt.args.defaultValue), "GetOrDefault(%v, %v)", tt.args.k, tt.args.defaultValue)
		})
	}
}

func TestSafeAnyBMap_Put(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "",
			fields: fields{
				mp: nil,
			},
			args: args{
				k: 1,
				v: 1,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 2,
				v: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.Put(tt.args.k, tt.args.v)
			v, ok := x.Get(tt.args.k)
			assert.Equal(t, true, ok)
			assert.Equal(t, tt.args.v, v)
		})
	}
}

func TestSafeAnyBMap_PuTIfAbsent(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "nil",
			fields: fields{},
			args: args{
				k: 1,
				v: 1,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 1,
				v: 1,
			},
			want: true,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 2},
			},
			args: args{
				k: 1,
				v: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.PuTIfAbsent(tt.args.k, tt.args.v), "PuTIfAbsent(%v, %v)", tt.args.k, tt.args.v)
		})
	}
}

func TestSafeAnyBMap_Delete(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				k: 0,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.Delete(tt.args.k)
			_, ok := x.Get(tt.args.k)
			assert.Equal(t, false, ok)
		})
	}
}

func TestSafeAnyBMap_DeleteIfPresent(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k int
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
				mp: nil,
			},
			args: args{
				k: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k: 0,
			},
			want:  0,
			want1: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2},
			},
			args: args{
				k: 1,
			},
			want:  1,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			got, got1 := x.DeleteIfPresent(tt.args.k)
			assert.Equalf(t, tt.want, got, "DeleteIfPresent(%v)", tt.args.k)
			assert.Equalf(t, tt.want1, got1, "DeleteIfPresent(%v)", tt.args.k)
			_, ok := x.Get(tt.args.k)
			assert.Equal(t, false, ok)
		})
	}
}

func TestSafeAnyBMap_MergeByMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		m map[int]int
		f func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: nil,
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 2},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 4, 4: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.MergeByMap(tt.args.m, tt.args.f)
			assert.Equal(t, tt.want, x.ToMetaMap())
		})
	}
}

func TestSafeAnyBMap_MergeByBMap(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		m map[int]int
		f func(int, int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[int]int
	}{
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: nil,
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{},
				f: nil,
			},
			want: map[int]int{},
		},
		{
			name: "nil",
			fields: fields{
				mp: nil,
			},
			args: args{
				m: map[int]int{1: 1},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: nil,
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 2},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				m: map[int]int{1: 2},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return false
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3, 4: 4},
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{1: 1, 2: 2, 3: 3},
			},
			args: args{
				m: map[int]int{3: 4, 4: 4},
				f: func(i int, i2 int) bool {
					return true
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 4, 4: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			x.MergeByBMap(NewSafeAnyBMapByMap(tt.args.m), tt.args.f)
			assert.Equal(t, tt.want, x.ToMetaMap())
		})
	}
}

func TestSafeAnyBMap_Replace(t *testing.T) {
	type fields struct {
		mp map[int]int
	}
	type args struct {
		k  int
		ov int
		nv int
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
				mp: nil,
			},
			args: args{
				k:  0,
				ov: 0,
				nv: 0,
			},
			want: false,
		},
		{
			name: "",
			fields: fields{
				mp: map[int]int{},
			},
			args: args{
				k:  0,
				ov: 0,
				nv: 0,
			},
			want: false,
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k:  1,
				ov: 0,
				nv: 2,
			},
			want: false,
		},
		{
			name: "nil",
			fields: fields{
				mp: map[int]int{1: 1},
			},
			args: args{
				k:  1,
				ov: 1,
				nv: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewSafeAnyBMapByMap(tt.fields.mp)
			assert.Equalf(t, tt.want, x.Replace(tt.args.k, tt.args.ov, tt.args.nv), "Replace(%v, %v, %v)", tt.args.k, tt.args.ov, tt.args.nv)
		})
	}
}
