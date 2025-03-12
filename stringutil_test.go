package flagx

import (
	"slices"
	"strconv"
	"testing"

	"github.com/carlmjohnson/be"
)

func TestJoin(t *testing.T) {
	be.Equal(t, "", joinFunc(slices.Values[[]int](nil), ", ", strconv.Itoa))
	be.Equal(t, "1", joinFunc(slices.Values([]int{1}), ", ", strconv.Itoa))
	be.Equal(t, "1, 2", joinFunc(slices.Values([]int{1, 2}), ", ", strconv.Itoa))
}
