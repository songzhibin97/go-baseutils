package bobjectstorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
}

func (t mock) Test() string {
	return "test"
}

func TestSet(t *testing.T) {
	mc := mock{}
	Set("key1", "string")
	Set("key2", 1)
	Set("key3", mc)
	Set("key4", &mc)
	assert.Equal(t, Get[string]("key1"), "string")
	assert.Equal(t, Get[int]("key2"), 1)
	assert.Equal(t, Get[mock]("key3"), mock{})
	assert.Equal(t, Get[*mock]("key4"), &mc)
}
