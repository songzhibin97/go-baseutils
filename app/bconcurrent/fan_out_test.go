package concurrent

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func production() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			out <- i
		}
	}()
	return out
}

func consumer(ch <-chan int) (ret []int) {
	for v := range ch {
		ret = append(ret, v)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return ret
}

func TestFanOut(t *testing.T) {
	var outList []chan int
	for i := 0; i < 10; i++ {
		outList = append(outList, make(chan int))
	}
	FanOut[int](production(), outList, true)
	wg := sync.WaitGroup{}
	for _, c := range outList {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			assert.Equal(t, consumer(c), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}(c)
	}
	wg.Wait()
}
