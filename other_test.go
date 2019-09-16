package cronexpr

import (
	"sort"
	"testing"
)

func TestSearchInt(t *testing.T) {
	ints := sort.SearchInts([]int{10, 20, 30, 40, 50}, 11)
	t.Log(ints)
}
