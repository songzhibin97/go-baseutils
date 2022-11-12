package bpoint

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type tInterface interface {
	test()
}

type tStruct struct {
	t any
}

func (t tStruct) test() {
}

func TestToPoint(t *testing.T) {
	var v int = 1
	rv := ToPoint(v)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv).Kind())
	assert.Equal(t, v, reflect.ValueOf(rv).Elem().Interface())

	var v2 *int = nil
	rv2 := ToPoint(v2)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv2).Kind())
	if rv2 != nil {
		t.Error("rv2 should be nil")
	}

	var v3 *int = &v
	rv3 := ToPoint(v3)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv2).Kind())
	assert.Equal(t, v3, reflect.ValueOf(rv3).Elem().Interface())

	var v4 tStruct = tStruct{t: 1}
	rv4 := ToPoint(v4)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv4).Kind())
	assert.Equal(t, v4, reflect.ValueOf(rv4).Elem().Interface())

	var v5 *tStruct = nil
	rv5 := ToPoint(v5)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv5).Kind())
	if rv5 != nil {
		t.Error("rv5 should be nil")
	}

	var v6 *tStruct = &v4
	rv6 := ToPoint(v6)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv6).Kind())
	assert.Equal(t, v6, reflect.ValueOf(rv6).Elem().Interface())

	var v7 tInterface = tStruct{t: 1}
	rv7 := ToPoint(v7)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv7).Kind())
	assert.Equal(t, v7, reflect.ValueOf(rv7).Elem().Interface())

	var v8 *tInterface = nil
	rv8 := ToPoint(v8)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv8).Kind())
	if rv8 != nil {
		t.Error("rv8 should be nil")
	}

	var v9 *tInterface = &v7
	rv9 := ToPoint(v9)
	assert.Equal(t, reflect.Ptr, reflect.ValueOf(rv9).Kind())
	assert.Equal(t, v9, reflect.ValueOf(rv9).Elem().Interface())
}
