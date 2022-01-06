package choose_test

import (
	"testing"

	"github.com/lestrrat-go/choose"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	ints := []int{5, 3, 1, 2, 4}
	c := choose.Slice[int](ints)
	c.One()
	if !assert.Len(t, c.N(3), 3) {
		return
	}
}

func TestMap(t *testing.T) {
	src := map[string]int{
		"hello": 1,
		"world": 3,
		"foo":5,
		"bar":7,
		"baz": 9,
	}
	c := choose.Map[string, int](src)
	if !assert.Len(t, c.N(3), 3) {
		return
	}
}
