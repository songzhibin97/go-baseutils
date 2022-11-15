package lscq

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLSCQ(t *testing.T) {
	lq := New[int]()

	rand.Seed(time.Now().Unix())
	q := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		v := rand.Int()
		q = append(q, v)
		lq.Enqueue(v)
	}
	for i := 0; i < 10; i++ {
		v, ok := lq.Dequeue()
		assert.Equal(t, true, ok)
		assert.Equal(t, q[i], v)
	}
	_, ok := lq.Dequeue()
	assert.Equal(t, false, ok)
}
