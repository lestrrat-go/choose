package choose

import (
	"math/rand"
	"time"
)

// SliceChooser is used to choose random elements from a slice
type SliceChooser[T any] struct {
	rng *rand.Rand
	src []T
	n int
}

// Slice creates a new SliceChooser specialied for type T.
// Upon calling this function, the elements from `src` are
// copied into a local storage: thus modifying `src` does
// not affect the subsequent operations.
func Slice[T any](src []T) *SliceChooser[T] {
	n := len(src)
	elements := make([]T, n)
	copy(elements, src)
	return &SliceChooser[T]{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
		src: elements,
		n: n,
	}
}

// One chooses one element from the source slice `src`
func (c *SliceChooser[T]) One() T {
	return c.src[c.rng.Intn(c.n)]
}

// N chooses as many elements specified by `howmany` parameter
// from the source slice `src`
func (c *SliceChooser[T]) N(howmany int) []T {
	indices := c.rng.Perm(c.n)
	dst := make([]T, howmany)
	for i, idx := range indices {
		if i >= howmany {
			break
		}
		dst[i] = c.src[idx]
	}
	return dst
}

type MapChooser[K comparable, V any] struct {
	rng *rand.Rand
	keys []K
	src map[K]V
	n int
}
type MapElement[K comparable, V any] struct {
	Key K
	Value V
}

func Map[K comparable, V any](src map[K]V) *MapChooser[K,V] {
	l := len(src)
	keys := make([]K, 0, l)
	for k := range src {
		keys = append(keys, k)
	}
	return &MapChooser[K,V]{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
		keys: keys,
		n: l,
	}
}

func (c *MapChooser[K,V]) One(src map[K]V) MapElement[K,V] {
	key := c.keys[c.rng.Intn(c.n)]
	return MapElement[K,V]{
		Key: key,
		Value: src[key],
	}
}

func (c *MapChooser[K,V]) N(howmany int) []MapElement[K,V] {
	indices := c.rng.Perm(c.n)
	dst := make([]MapElement[K,V], howmany)
	for i, idx := range indices {
		if i >= howmany {
			break
		}
		key := c.keys[idx]
		dst[i] = MapElement[K,V]{
			Key: key,
			Value: c.src[key],
		}
	}
	return dst
}
